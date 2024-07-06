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

type DnsServiceServer struct {
	proto.UnimplementedDnsServiceServer
	Log *grpclog.LoggerV2
}

func (s *DnsServiceServer) GetHostname(ctx context.Context, in *proto.GetHostnameParams) (*proto.Hostname, error) {
	stdout, _, err := runCmd("hostname")
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed to get the hostname: %w", err)
	}
	name := strings.TrimSuffix(stdout, "\n")
	return &proto.Hostname{Name: name}, nil
}

func (s *DnsServiceServer) SetHostname(ctx context.Context, in *proto.Hostname) (*proto.Hostname, error) {
	_, _, err := runCmd("hostnamectl", "hostname", in.GetName())
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed to change the hostname: %w", err)
	}
	// Invalid names may be changed without any error status or indication from 'hostnamectl', so we have to get it back manually
	stdout, _, err := runCmd("hostname")
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed ot get the changed hostname: %w", err)
	}
	name := strings.TrimSuffix(stdout, "\n")
	return &proto.Hostname{Name: name}, nil
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
