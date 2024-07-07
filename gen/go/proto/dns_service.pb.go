// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: dns_service.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Hostname struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Hostname) Reset() {
	*x = Hostname{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hostname) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hostname) ProtoMessage() {}

func (x *Hostname) ProtoReflect() protoreflect.Message {
	mi := &file_dns_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hostname.ProtoReflect.Descriptor instead.
func (*Hostname) Descriptor() ([]byte, []int) {
	return file_dns_service_proto_rawDescGZIP(), []int{0}
}

func (x *Hostname) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DnsServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DnsServer) Reset() {
	*x = DnsServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsServer) ProtoMessage() {}

func (x *DnsServer) ProtoReflect() protoreflect.Message {
	mi := &file_dns_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsServer.ProtoReflect.Descriptor instead.
func (*DnsServer) Descriptor() ([]byte, []int) {
	return file_dns_service_proto_rawDescGZIP(), []int{1}
}

func (x *DnsServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DnsServers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*DnsServer `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *DnsServers) Reset() {
	*x = DnsServers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsServers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsServers) ProtoMessage() {}

func (x *DnsServers) ProtoReflect() protoreflect.Message {
	mi := &file_dns_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsServers.ProtoReflect.Descriptor instead.
func (*DnsServers) Descriptor() ([]byte, []int) {
	return file_dns_service_proto_rawDescGZIP(), []int{2}
}

func (x *DnsServers) GetList() []*DnsServer {
	if x != nil {
		return x.List
	}
	return nil
}

var File_dns_service_proto protoreflect.FileDescriptor

var file_dns_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x44,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x12, 0x2a, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6e, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xde, 0x08, 0x0a,
	0x0a, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xde, 0x01, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x64, 0x6e,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x15, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa0, 0x01, 0x92, 0x41, 0x84, 0x01,
	0x12, 0x1e, 0x53, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x1a, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x27, 0x73, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x20,
	0x76, 0x69, 0x61, 0x20, 0x61, 0x20, 0x60, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x63,
	0x74, 0x6c, 0x60, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x20, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x69, 0x66, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xb0, 0x01, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x92, 0x41, 0x5a,
	0x12, 0x20, 0x47, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x1a, 0x36, 0x47, 0x65, 0x74, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x27, 0x73, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x76,
	0x69, 0x61, 0x20, 0x61, 0x20, 0x60, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x63, 0x74,
	0x6c, 0x60, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0xe5, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x64, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x22, 0xa1, 0x01, 0x92, 0x41, 0x85, 0x01, 0x12, 0x2c, 0x47, 0x65, 0x74, 0x20,
	0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x6e, 0x73, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x65, 0x1a, 0x55, 0x47, 0x65, 0x74, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x27, 0x73, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x64, 0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x60,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x63, 0x74, 0x6c, 0x20, 0x64, 0x6e, 0x73, 0x60, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6e, 0x73, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0xe0, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x64, 0x6e, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x1a, 0x17, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x9c, 0x01, 0x92, 0x41,
	0x7e, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x20, 0x64, 0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x75,
	0x73, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x65, 0x1a, 0x4c, 0x41, 0x64, 0x64, 0x73, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x64, 0x6e, 0x73, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x27, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x63, 0x61, 0x6c, 0x6c, 0x65, 0x65, 0x27, 0x73, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x20, 0x76, 0x69, 0x61, 0x20, 0x60,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x63, 0x74, 0x6c, 0x20, 0x64, 0x6e, 0x73, 0x60, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0xf0, 0x01, 0x0a, 0x10, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x17, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x1a, 0x17, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x22, 0xa9, 0x01, 0x92, 0x41, 0x83, 0x01, 0x12, 0x33, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x64,
	0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20,
	0x63, 0x61, 0x6c, 0x6c, 0x65, 0x65, 0x27, 0x73, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x4c, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x20, 0x64, 0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x65, 0x27, 0x73, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x20, 0x76, 0x69, 0x61, 0x20, 0x60, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x63, 0x74, 0x6c, 0x20, 0x64, 0x6e, 0x73, 0x60, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x3a, 0x01, 0x2a, 0x1a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6e, 0x73, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x14, 0x5a,
	0x12, 0x79, 0x61, 0x64, 0x72, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dns_service_proto_rawDescOnce sync.Once
	file_dns_service_proto_rawDescData = file_dns_service_proto_rawDesc
)

func file_dns_service_proto_rawDescGZIP() []byte {
	file_dns_service_proto_rawDescOnce.Do(func() {
		file_dns_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_dns_service_proto_rawDescData)
	})
	return file_dns_service_proto_rawDescData
}

var file_dns_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dns_service_proto_goTypes = []any{
	(*Hostname)(nil),    // 0: dns_service.Hostname
	(*DnsServer)(nil),   // 1: dns_service.DnsServer
	(*DnsServers)(nil),  // 2: dns_service.DnsServers
	(*empty.Empty)(nil), // 3: google.protobuf.Empty
}
var file_dns_service_proto_depIdxs = []int32{
	1, // 0: dns_service.DnsServers.list:type_name -> dns_service.DnsServer
	0, // 1: dns_service.DnsService.SetHostname:input_type -> dns_service.Hostname
	3, // 2: dns_service.DnsService.GetHostname:input_type -> google.protobuf.Empty
	3, // 3: dns_service.DnsService.ListDnsServers:input_type -> google.protobuf.Empty
	2, // 4: dns_service.DnsService.AddDnsServers:input_type -> dns_service.DnsServers
	2, // 5: dns_service.DnsService.RemoveDnsServers:input_type -> dns_service.DnsServers
	0, // 6: dns_service.DnsService.SetHostname:output_type -> dns_service.Hostname
	0, // 7: dns_service.DnsService.GetHostname:output_type -> dns_service.Hostname
	2, // 8: dns_service.DnsService.ListDnsServers:output_type -> dns_service.DnsServers
	2, // 9: dns_service.DnsService.AddDnsServers:output_type -> dns_service.DnsServers
	2, // 10: dns_service.DnsService.RemoveDnsServers:output_type -> dns_service.DnsServers
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dns_service_proto_init() }
func file_dns_service_proto_init() {
	if File_dns_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dns_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Hostname); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dns_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DnsServer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dns_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DnsServers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dns_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dns_service_proto_goTypes,
		DependencyIndexes: file_dns_service_proto_depIdxs,
		MessageInfos:      file_dns_service_proto_msgTypes,
	}.Build()
	File_dns_service_proto = out.File
	file_dns_service_proto_rawDesc = nil
	file_dns_service_proto_goTypes = nil
	file_dns_service_proto_depIdxs = nil
}
