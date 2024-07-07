package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"yadro/gen/go/proto"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "start grpc and http servers for the dns-service",
	Run:   runServer,
}

func init() {
	rootCmd.Flags().UintP("rest-port", "r", 1234, "port used for the REST http-server")
	rootCmd.Flags().UintP("grpc-port", "g", 1235, "port used for the grpc server")
	rootCmd.Flags().BoolP("debug-dns", "d", false, "if set to true, server doesn't set 'dns=none' in NetworkManager.conf")
	rootCmd.Flags().BoolP("grpc-logs", "l", false, "if set to true, prints the grpc library logs to stdout")
}

func runServer(cmd *cobra.Command, _args []string) {
	grpcPort, _ := cmd.Flags().GetUint("grpc-port")
	restPort, _ := cmd.Flags().GetUint("rest-port")
	debugDns, _ := cmd.Flags().GetBool("debug-dns")
	grpcLogs, _ := cmd.Flags().GetBool("grpc-logs")

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	if grpcLogs {
		grpclog.SetLoggerV2(log)
	}

	if !debugDns {
		err := makeDnsUnmanaged()
		if err != nil {
			log.Fatalf("failed to configure NetworkManager: %v", err)
		}
		log.Info("set 'dns=none' in NetworkManager.conf")
	} else {
		log.Warning("ran with debug-dns. Dns changes will not be permanent")
	}

	addr := fmt.Sprintf("0.0.0.0:%d", grpcPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterDnsServiceServer(s, &DnsServiceServer{
		Log: log,
	})

	log.Infof("serving gRPC on http://%s", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	restAddr := fmt.Sprintf("0.0.0.0:%d", restPort)
	rpcAddr := "dns:///" + addr
	log.Infof("serving REST on http://%s", restAddr)
	err = GatewayRun(restAddr, rpcAddr)

	if err != nil {
		log.Fatal(err)
	}
}

// Sets the DNS (resolv.conf) processing mode to none
// NetworkManager will not modify resolv.conf
// This is needed because otherwise we can't control the dns servers
func makeDnsUnmanaged() error {
	confPath := "/etc/NetworkManager/NetworkManager.conf"
	config, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(config), "\n")

	dnsLine := -1
	mainLine := -1
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		if line == "[main]" {
			mainLine = i
		}
		if strings.HasPrefix(line, "dns=") {
			dnsLine = i
		}
	}

	if mainLine == -1 && dnsLine != -1 {
		return fmt.Errorf("failed to parse: found dns key-value pair, but not [main] section")
	}

	if mainLine == -1 {
		lines = append([]string{"[main]", "dns=none"}, lines...)
	} else if dnsLine == -1 {
		lines = slices.Insert(lines, mainLine+1, "dns=none")
	} else {
		lines[dnsLine] = "dns=none"
	}

	info, err := os.Stat(confPath)
	if err != nil {
		return err
	}
	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(confPath, []byte(output), info.Mode())
}
