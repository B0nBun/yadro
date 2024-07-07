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
	Log grpclog.LoggerV2
}

func (s *DnsServiceServer) GetHostname(context.Context, *empty.Empty) (*proto.Hostname, error) {
	stdout, _, err := s.runCmd("hostname")
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed to get the hostname: %w", err)
	}
	name := strings.TrimSuffix(stdout, "\n")
	return &proto.Hostname{Name: name}, nil
}

func (s *DnsServiceServer) SetHostname(_ context.Context, in *proto.Hostname) (*proto.Hostname, error) {
	_, _, err := s.runCmd("hostnamectl", "hostname", in.GetName())
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed to change the hostname: %w", err)
	}
	// Invalid names may be changed without any error status or indication from 'hostnamectl', so we have to get it back manually
	stdout, _, err := s.runCmd("hostname")
	if err != nil {
		return &proto.Hostname{}, fmt.Errorf("failed ot get the changed hostname: %w", err)
	}
	name := strings.TrimSuffix(stdout, "\n")
	return &proto.Hostname{Name: name}, nil
}


func (s *DnsServiceServer) ListDnsServers(context.Context, *empty.Empty) (*proto.DnsServers, error) {
	resp := &proto.DnsServers{} 
	iface, err :=  s.getDefaultIface()
	if err != nil {
		return resp, err
	}
	addrs, err :=  s.getDnsServers(iface)
	if err != nil {
		return resp, err
	}
	servers := make([]*proto.DnsServer, len(addrs))
	for i, addr := range addrs {
		servers[i] = &proto.DnsServer{
			Address: addr,
		}
	}
	resp.List = servers
	return resp, nil
}

func (s *DnsServiceServer) AddDnsServers(_ context.Context, servers *proto.DnsServers) (*proto.DnsServers, error) {
	resp := &proto.DnsServers{}
	iface, err :=  s.getDefaultIface()
	if err != nil {
		return resp, err
	}
	addrs, err :=  s.getDnsServers(iface)
	if err != nil {
		return resp, err
	}
	for _, server := range servers.List {
		addrs = append(addrs, server.Address)
	}
	err =  s.setDnsServers(iface, normalizeAddrs(addrs))
	if err != nil {
		return resp, err
	}
	addrs, err = s.getDnsServers(iface)
	resp.List = make([]*proto.DnsServer, len(addrs))
	for i, addr := range addrs {
		resp.List[i] = &proto.DnsServer{Address: addr}
	}
	return resp, nil
}

func (s *DnsServiceServer) RemoveDnsServers(_ context.Context, toDelete *proto.DnsServers) (*proto.DnsServers, error) {
	resp := &proto.DnsServers{}
	iface, err :=  s.getDefaultIface()
	if err != nil {
		return resp, err
	}
	addrs, err :=  s.getDnsServers(iface)
	if err != nil {
		return resp, err
	}
	addrs = slices.DeleteFunc(addrs, func (addr string) bool {
		return slices.IndexFunc(toDelete.List, func (s *proto.DnsServer) bool { return s.Address == addr }) != -1
	})
	err =  s.setDnsServers(iface, normalizeAddrs(addrs))
	if err != nil {
		return resp, err
	}
	addrs, err = s.getDnsServers(iface)
	resp.List = make([]*proto.DnsServer, len(addrs))
	for i, addr := range addrs {
		resp.List[i] = &proto.DnsServer{ Address: addr }
	}
	return resp, nil
}

// `resolvectl` accepts full formats like this: "111.222.333.444:9953%ifname#example.com" for IPv4 and "[1111:2222::3333]:9953%ifname#example.com" for IPv6
// However, inconsistent/dynamic ifname may cause errors in some scenarios, so we just strip it away
func normalizeAddrs(addrs []string) (res []string) {
	res = make([]string, len(addrs))
	for i, addr := range addrs {
		res[i] = strings.Split(addr, "%")[0]
	}
	return
}

func (s *DnsServiceServer) getDefaultIface() (string, error) {
	stdout, _, err := s.runCmd("ip", "route", "show", "default")
	if err != nil {
		return "", err
	}
	split := strings.Split(stdout, " ")
	if len(split) < 5 {
		return "", fmt.Errorf("expected interface as 5th word, got '%s'", stdout)
	}
	return split[4], nil
}

func (s *DnsServiceServer) getDnsServers(iface string) ([]string, error) {
	stdout, _, err := s.runCmd("resolvectl", "dns", iface)
	if err != nil {
		return nil, err
	}
	split := strings.SplitN(strings.TrimSuffix(stdout, "\n"), ": ", 2)
	if len(split) < 2 {
		return nil, fmt.Errorf("expected colon seperator ': ', got '%s'", stdout)
	}
	addrs := split[1]
	return strings.Fields(addrs), nil
}

func (s *DnsServiceServer) setDnsServers(iface string, addrs []string) error {
	args := append([]string { "dns", iface }, addrs...)
	_, _, err := s.runCmd("resolvectl", args...)
	return err
}

func (s *DnsServiceServer) runCmd(cmd string, args ...string) (stdout string, stderr string, err error) {
	c := exec.Command(cmd, args...)
	s.Log.Infof("executing command: %v %v", cmd, args)
	stdoutbuf, stderrbuf := new(bytes.Buffer), new(bytes.Buffer)
	c.Stdout = stdoutbuf
	c.Stderr = stderrbuf
	err = c.Run()
	stdout = stdoutbuf.String()
	stderr = stderrbuf.String()
	if err != nil {
		err := fmt.Errorf("%w: %s", err, stderr)
		s.Log.Infof("command result err: %v", err)
		return stdout, stderr, err
	}
	return
}
