package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	proto "yadro-dns/gen/go/proto"
	"yadro-dns/server/gateway"
)

// TODO: Actual task with DNS and stuff
// TODO: Check out if HTTPS is needed

type Server struct {
	proto.UnimplementedDnsServiceServer
	log grpclog.LoggerV2
}

func (s *Server) Echo(ctx context.Context, in *proto.StringMessage) (*proto.StringMessage, error) {
	s.log.Infof("Recieved: %v", in)
	return in, nil
}

// TODO: All configuration should be done with flags (Cobra)
const GRPC_PORT = 1234
const REST_PORT = 1235

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := fmt.Sprintf("0.0.0.0:%d", GRPC_PORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // TODO: Replace with a certificate if necessary!

	proto.RegisterDnsServiceServer(s, &Server{
		log: log,
	})

	log.Infof("serving gRPC on http://%s", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	restAddr := fmt.Sprintf("0.0.0.0:%d", REST_PORT)
	rpcAddr := "dns:///" + addr
	log.Infof("serving REST on http://%s", restAddr)
	err = gateway.Run(restAddr, rpcAddr)

	if err != nil {
		log.Fatal(err)
	}
}
