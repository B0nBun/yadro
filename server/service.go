package main

import (
	"bytes"
	"slices"
	"context"
	"fmt"
	"os/exec"

	"google.golang.org/grpc/grpclog"
	"github.com/golang/protobuf/ptypes/empty"
	"strings"
	"yadro/gen/go/proto"
)

type DnsServiceServer struct {
	proto.UnimplementedDnsServiceServer
	Log *grpclog.LoggerV2
	DnsServers []*proto.DnsServer
}

func (s *DnsServiceServer) GetHostname(context.Context, *empty.Empty) (*proto.Hostname, error) {
	stdout, _, err := runCmd("hostname")
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed to get the hostname: %w", err)
	}
	name := strings.TrimSuffix(stdout, "\n")
	return &proto.Hostname{Name: name}, nil
}

func (s *DnsServiceServer) SetHostname(_ context.Context, in *proto.Hostname) (*proto.Hostname, error) {
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

func (s *DnsServiceServer) ListDnsServers(context.Context, *empty.Empty) (*proto.ListDnsServerResponse, error) {
	return &proto.ListDnsServerResponse {
		Servers: s.DnsServers,
	}, nil
}

func (s *DnsServiceServer) AddDnsServer(_ context.Context, server *proto.DnsServer) (*proto.AddDnsServerResponse, error) {
	// TODO: Not add duplicate dns servers
	s.DnsServers = append(s.DnsServers, server)
	return &proto.AddDnsServerResponse {
		Added: true,
		Servers: s.DnsServers,
	}, nil
}

func (s *DnsServiceServer) RemoveDnsServer(_ context.Context, server *proto.DnsServer) (*proto.RemoveDnsServerResponse, error) {
	idx := slices.IndexFunc(s.DnsServers, func (s *proto.DnsServer) bool { return s.Ip == server.Ip })
	if idx == -1 {
		return &proto.RemoveDnsServerResponse{
			Removed: false,
			Servers: s.DnsServers,
		}, nil
	}
	s.DnsServers = append(s.DnsServers[:idx], s.DnsServers[idx+1:]...)
	return &proto.RemoveDnsServerResponse{
		Removed: true,
		Servers: s.DnsServers,
	}, nil
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
