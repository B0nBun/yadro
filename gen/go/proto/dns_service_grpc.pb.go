// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: dns_service.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DnsService_SetHostname_FullMethodName     = "/dns_service.DnsService/SetHostname"
	DnsService_GetHostname_FullMethodName     = "/dns_service.DnsService/GetHostname"
	DnsService_ListDnsServers_FullMethodName  = "/dns_service.DnsService/ListDnsServers"
	DnsService_AddDnsServer_FullMethodName    = "/dns_service.DnsService/AddDnsServer"
	DnsService_RemoveDnsServer_FullMethodName = "/dns_service.DnsService/RemoveDnsServer"
)

// DnsServiceClient is the client API for DnsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsServiceClient interface {
	SetHostname(ctx context.Context, in *Hostname, opts ...grpc.CallOption) (*Hostname, error)
	GetHostname(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Hostname, error)
	ListDnsServers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DnsServers, error)
	AddDnsServer(ctx context.Context, in *DnsServers, opts ...grpc.CallOption) (*DnsServers, error)
	// "DELETE" method isn't good for requests with body and specifying a list of dns-servers in query-params
	// is quite awful, so I chose to use POST with a special endpoint
	RemoveDnsServer(ctx context.Context, in *DnsServers, opts ...grpc.CallOption) (*DnsServers, error)
}

type dnsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsServiceClient(cc grpc.ClientConnInterface) DnsServiceClient {
	return &dnsServiceClient{cc}
}

func (c *dnsServiceClient) SetHostname(ctx context.Context, in *Hostname, opts ...grpc.CallOption) (*Hostname, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Hostname)
	err := c.cc.Invoke(ctx, DnsService_SetHostname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsServiceClient) GetHostname(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Hostname, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Hostname)
	err := c.cc.Invoke(ctx, DnsService_GetHostname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsServiceClient) ListDnsServers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DnsServers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DnsServers)
	err := c.cc.Invoke(ctx, DnsService_ListDnsServers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsServiceClient) AddDnsServer(ctx context.Context, in *DnsServers, opts ...grpc.CallOption) (*DnsServers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DnsServers)
	err := c.cc.Invoke(ctx, DnsService_AddDnsServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsServiceClient) RemoveDnsServer(ctx context.Context, in *DnsServers, opts ...grpc.CallOption) (*DnsServers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DnsServers)
	err := c.cc.Invoke(ctx, DnsService_RemoveDnsServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsServiceServer is the server API for DnsService service.
// All implementations must embed UnimplementedDnsServiceServer
// for forward compatibility
type DnsServiceServer interface {
	SetHostname(context.Context, *Hostname) (*Hostname, error)
	GetHostname(context.Context, *empty.Empty) (*Hostname, error)
	ListDnsServers(context.Context, *empty.Empty) (*DnsServers, error)
	AddDnsServer(context.Context, *DnsServers) (*DnsServers, error)
	// "DELETE" method isn't good for requests with body and specifying a list of dns-servers in query-params
	// is quite awful, so I chose to use POST with a special endpoint
	RemoveDnsServer(context.Context, *DnsServers) (*DnsServers, error)
	mustEmbedUnimplementedDnsServiceServer()
}

// UnimplementedDnsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDnsServiceServer struct {
}

func (UnimplementedDnsServiceServer) SetHostname(context.Context, *Hostname) (*Hostname, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHostname not implemented")
}
func (UnimplementedDnsServiceServer) GetHostname(context.Context, *empty.Empty) (*Hostname, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostname not implemented")
}
func (UnimplementedDnsServiceServer) ListDnsServers(context.Context, *empty.Empty) (*DnsServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDnsServers not implemented")
}
func (UnimplementedDnsServiceServer) AddDnsServer(context.Context, *DnsServers) (*DnsServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDnsServer not implemented")
}
func (UnimplementedDnsServiceServer) RemoveDnsServer(context.Context, *DnsServers) (*DnsServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDnsServer not implemented")
}
func (UnimplementedDnsServiceServer) mustEmbedUnimplementedDnsServiceServer() {}

// UnsafeDnsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsServiceServer will
// result in compilation errors.
type UnsafeDnsServiceServer interface {
	mustEmbedUnimplementedDnsServiceServer()
}

func RegisterDnsServiceServer(s grpc.ServiceRegistrar, srv DnsServiceServer) {
	s.RegisterService(&DnsService_ServiceDesc, srv)
}

func _DnsService_SetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hostname)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).SetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_SetHostname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).SetHostname(ctx, req.(*Hostname))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsService_GetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).GetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_GetHostname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).GetHostname(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsService_ListDnsServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).ListDnsServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_ListDnsServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).ListDnsServers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsService_AddDnsServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsServers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).AddDnsServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_AddDnsServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).AddDnsServer(ctx, req.(*DnsServers))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsService_RemoveDnsServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsServers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).RemoveDnsServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_RemoveDnsServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).RemoveDnsServer(ctx, req.(*DnsServers))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsService_ServiceDesc is the grpc.ServiceDesc for DnsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dns_service.DnsService",
	HandlerType: (*DnsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHostname",
			Handler:    _DnsService_SetHostname_Handler,
		},
		{
			MethodName: "GetHostname",
			Handler:    _DnsService_GetHostname_Handler,
		},
		{
			MethodName: "ListDnsServers",
			Handler:    _DnsService_ListDnsServers_Handler,
		},
		{
			MethodName: "AddDnsServer",
			Handler:    _DnsService_AddDnsServer_Handler,
		},
		{
			MethodName: "RemoveDnsServer",
			Handler:    _DnsService_RemoveDnsServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dns_service.proto",
}
