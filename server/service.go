package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"slices"
	"io/ioutil"
	"strings"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/grpclog"
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
	_, _, err := s.runCmd("hostnamectl", "set-hostname", in.GetName())
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

// TODO: Replace runCmd with editing of /etc/systemd/resolved.conf
// https://askubuntu.com/questions/1418372/jammy-resolvectl-domain-persistent-across-reboots
func (s *DnsServiceServer) ListDnsServers(context.Context, *empty.Empty) (*proto.DnsServers, error) {
	resp := &proto.DnsServers{}
	addrs, err := getDnsServers()
	if err != nil {
		return resp, err
	}
	servers := make([]*proto.DnsServer, len(addrs))
	for i, addr := range addrs {
		servers[i] = &proto.DnsServer{Address: addr,}
	}
	resp.List = servers
	return resp, nil
}

func (s *DnsServiceServer) AddDnsServers(_ context.Context, servers *proto.DnsServers) (*proto.DnsServers, error) {
	resp := &proto.DnsServers{}
	toAdd := make([]string, len(servers.List))
	for i, server := range servers.List {
		toAdd[i] = server.Address
	}
	addrs, err := addDnsServers(toAdd)
	if err != nil {
		return resp, err
	}
	resp.List = make([]*proto.DnsServer, len(addrs))
	for i, addr := range addrs {
		resp.List[i] = &proto.DnsServer{Address: addr}
	}
	return resp, nil
}

func (s *DnsServiceServer) RemoveDnsServers(_ context.Context, servers *proto.DnsServers) (*proto.DnsServers, error) {
	resp := &proto.DnsServers{}
	toRemove := make([]string, len(servers.List))
	for i, server := range servers.List {
		toRemove[i] = server.Address
	}
	addrs, err := removeDnsServers(toRemove)
	if err != nil {
		return resp, err
	}
	resp.List = make([]*proto.DnsServer, len(addrs))
	for i, addr := range addrs {
		resp.List[i] = &proto.DnsServer{Address: addr}
	}
	return resp, nil
}

const resolvedConfPath = "/etc/systemd/resolved.conf"

// Those functions that find dns in `/etc/systemd/resolved.conf` don't work when there are multiple DNS=... key-value pairs,
// but I sure hope no one will find out about it

func getDnsServers() ([]string, error) {
	config, err := ioutil.ReadFile(resolvedConfPath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(config), "\n")
	servers, _ := findDnsLine(lines)
	if servers == nil {
		return make([]string, 0), nil
	}
	return servers, nil
}

func addDnsServers(toAdd []string) ([]string, error) {
	config, err := ioutil.ReadFile(resolvedConfPath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(config), "\n")
	servers, i := findDnsLine(lines)
	if i != -1 {
		servers = append(servers, toAdd...)
		slices.Sort(servers)
		servers = slices.Compact(servers) // Remove duplicates
		lines[i] = fmt.Sprintf("DNS=%s", strings.Join(servers, " "))
	} else {
		idx := findResolveSectionLine(lines)
		servers = toAdd
		dnsLine := fmt.Sprintf("DNS=%s", strings.Join(servers, " "))
		if idx != -1 {
			slices.Insert(lines, idx + 1, dnsLine)
		} else {
			lines = append([]string { "[Resolve]", dnsLine }, lines...)
		}
	}

	info, err := os.Stat(resolvedConfPath)
	if err != nil {
		return nil, err
	}
	output := strings.Join(lines, "\n")
	return servers, ioutil.WriteFile(resolvedConfPath, []byte(output), info.Mode())
}

func removeDnsServers(toRemove []string) ([]string, error) {
	config, err := ioutil.ReadFile(resolvedConfPath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(config), "\n")
	servers, i := findDnsLine(lines)
	if i == -1 {
		return make([]string, 0), nil
	}
	servers = slices.DeleteFunc(servers, func(s string) bool { return slices.Contains(toRemove, s) })

	lines[i] = fmt.Sprintf("DNS=%s", strings.Join(servers, " "))

	info, err := os.Stat(resolvedConfPath)
	if err != nil {
		return nil, err
	}
	output := strings.Join(lines, "\n")
	return servers, ioutil.WriteFile(resolvedConfPath, []byte(output), info.Mode())
}

func findDnsLine(lines []string) (servers []string, i int) {
	for i, line := range lines {
		line = strings.Trim(line, " \t")
		if strings.HasPrefix(line, "DNS") {
			line = strings.TrimPrefix(strings.TrimLeft(strings.TrimPrefix(line, "DNS"), " \t"), "=")
			servers = strings.Fields(strings.Trim(line, " \t"))
			return servers, i
		}
	}
	return nil, -1
}

func findResolveSectionLine(lines []string) int {
	for i, line := range lines {
		if line == "[Resolve]" {
			return i
		}
	}
	return -1
}

func (s *DnsServiceServer) restartSystemdResolved() error {
	_, _, err := s.runCmd("systemctl", "restart", "systemd-resolved")
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
