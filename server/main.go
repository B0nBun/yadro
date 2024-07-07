package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

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
	rootCmd.Flags().UintP("rest-port", "r", 0, "port used for the REST http-gateway (Gateway doesn't run if the port is not specified)")
	rootCmd.Flags().UintP("grpc-port", "g", 1234, "port used for the grpc server")
	rootCmd.Flags().BoolP("grpc-logs", "l", false, "if set to true, prints the grpc library logs to stdout")
}

func runServer(cmd *cobra.Command, _args []string) {
	grpcPort, _ := cmd.Flags().GetUint("grpc-port")
	restPort, _ := cmd.Flags().GetUint("rest-port")
	grpcLogs, _ := cmd.Flags().GetBool("grpc-logs")

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	if grpcLogs {
		grpclog.SetLoggerV2(log)
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

	go func() {
		log.Infof("serving gRPC on http://%s", addr)
		log.Fatal(s.Serve(lis))
	}()

	if restPort != 0 {
		restAddr := fmt.Sprintf("0.0.0.0:%d", restPort)
		rpcAddr := "dns:///" + addr
		go func() {
			log.Infof("serving REST on http://%s", restAddr)
			log.Fatal(GatewayRun(restAddr, rpcAddr))
		}()
	}

	// Sleep forever, until the gateway or gRPC server quits with log.Fatal
	<-make(chan struct{})
}
