package cmd

import (
	"fmt"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/spf13/cobra"
	"yadro-dns/gen/go/proto"
	"yadro-dns/server/gateway"
	"yadro-dns/server/service"
)

var rootCmd = &cobra.Command{
	Use: "server",
	Short: "start grpc and http servers for the dns-service",
	Run: runServer,
}

func init() {
	rootCmd.Flags().UintP("http-port", "H", 1234, "port used for the REST http-server")
	rootCmd.Flags().UintP("grpc-port", "G", 1235, "port used for the grpc server")
}

func runServer(cmd *cobra.Command, args []string) {
	grpcPort, _ := cmd.Flags().GetUint("grpc-port")
	httpPort, _ := cmd.Flags().GetUint("http-port")

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := fmt.Sprintf("0.0.0.0:%d", grpcPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // TODO: Replace with a certificate if necessary!

	proto.RegisterDnsServiceServer(s, &service.Server{
		Log: &log,
	})

	log.Infof("serving gRPC on http://%s", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	restAddr := fmt.Sprintf("0.0.0.0:%d", httpPort)
	rpcAddr := "dns:///" + addr
	log.Infof("serving REST on http://%s", restAddr)
	err = gateway.Run(restAddr, rpcAddr)

	if err != nil {
		log.Fatal(err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}