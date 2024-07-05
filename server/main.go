package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"strings"
	"yadro-dns/gen/go/proto"
	"yadro-dns/server/gateway"
)

// TODO: Actual task with DNS and stuff
// TODO: Check out if HTTPS is needed

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
		log: &log,
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


type Server struct {
	proto.UnimplementedDnsServiceServer
	log *grpclog.LoggerV2
}

func (s *Server) GetHostname(ctx context.Context, in *proto.GetHostnameParams) (*proto.Hostname, error) {
	stdout, _, err := runCmd("hostnamectl", "hostname")
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed to get the hostname: %w", err)
	}
	name := strings.TrimSuffix(stdout, "\n")
	return &proto.Hostname{Name: name}, nil
}

func (s *Server) SetHostname(ctx context.Context, in *proto.Hostname) (*proto.Hostname, error) {
	// TODO: Check what happens when an invalid name is provided
	_, _, err := runCmd("hostnamectl", "hostname", in.GetName())
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed to change the hostname: %w", err)
	}
	return in, nil
}

func runCmd(cmd string, args ...string) (stdout string, stderr string, err error) {
	c := exec.Command(cmd, args...)
	stdoutbuf, stderrbuf := new(bytes.Buffer), new(bytes.Buffer)
	c.Stdout = stdoutbuf
	c.Stderr = stderrbuf
	err = c.Run()
	stdout = stdoutbuf.String()
	stderr = stderrbuf.String()
	if err != nil {
		return stdout, stderr, fmt.Errorf("%w: %s", err, stderr)
	}
	return
}
