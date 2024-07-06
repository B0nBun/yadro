package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"google.golang.org/grpc/grpclog"
	"strings"
	"yadro-dns/gen/go/proto"
)

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
