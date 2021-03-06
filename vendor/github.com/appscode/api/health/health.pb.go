// Code generated by protoc-gen-go.
// source: health.proto
// DO NOT EDIT!

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	StatusResponse
	URLBase
	NetConfig
	Metadata
*/
package health

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"
import appscode_version "github.com/appscode/api/version"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusResponse struct {
	Status   *appscode_dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Version  *appscode_version.Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Metadata *Metadata                 `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatusResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StatusResponse) GetVersion() *appscode_version.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *StatusResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type URLBase struct {
	Scheme   string `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	BaseAddr string `protobuf:"bytes,2,opt,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
}

func (m *URLBase) Reset()                    { *m = URLBase{} }
func (m *URLBase) String() string            { return proto.CompactTextString(m) }
func (*URLBase) ProtoMessage()               {}
func (*URLBase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *URLBase) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *URLBase) GetBaseAddr() string {
	if m != nil {
		return m.BaseAddr
	}
	return ""
}

type NetConfig struct {
	TeamAddr         string   `protobuf:"bytes,1,opt,name=team_addr,json=teamAddr" json:"team_addr,omitempty"`
	PublicUrls       *URLBase `protobuf:"bytes,2,opt,name=public_urls,json=publicUrls" json:"public_urls,omitempty"`
	TeamUrls         *URLBase `protobuf:"bytes,3,opt,name=team_urls,json=teamUrls" json:"team_urls,omitempty"`
	ClusterUrls      *URLBase `protobuf:"bytes,4,opt,name=cluster_urls,json=clusterUrls" json:"cluster_urls,omitempty"`
	InClusterUrls    *URLBase `protobuf:"bytes,5,opt,name=in_cluster_urls,json=inClusterUrls" json:"in_cluster_urls,omitempty"`
	URLShortenerUrls *URLBase `protobuf:"bytes,6,opt,name=URL_shortener_urls,json=URLShortenerUrls" json:"URL_shortener_urls,omitempty"`
	FileUrls         *URLBase `protobuf:"bytes,7,opt,name=file_urls,json=fileUrls" json:"file_urls,omitempty"`
}

func (m *NetConfig) Reset()                    { *m = NetConfig{} }
func (m *NetConfig) String() string            { return proto.CompactTextString(m) }
func (*NetConfig) ProtoMessage()               {}
func (*NetConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NetConfig) GetTeamAddr() string {
	if m != nil {
		return m.TeamAddr
	}
	return ""
}

func (m *NetConfig) GetPublicUrls() *URLBase {
	if m != nil {
		return m.PublicUrls
	}
	return nil
}

func (m *NetConfig) GetTeamUrls() *URLBase {
	if m != nil {
		return m.TeamUrls
	}
	return nil
}

func (m *NetConfig) GetClusterUrls() *URLBase {
	if m != nil {
		return m.ClusterUrls
	}
	return nil
}

func (m *NetConfig) GetInClusterUrls() *URLBase {
	if m != nil {
		return m.InClusterUrls
	}
	return nil
}

func (m *NetConfig) GetURLShortenerUrls() *URLBase {
	if m != nil {
		return m.URLShortenerUrls
	}
	return nil
}

func (m *NetConfig) GetFileUrls() *URLBase {
	if m != nil {
		return m.FileUrls
	}
	return nil
}

type Metadata struct {
	Env       string     `protobuf:"bytes,1,opt,name=env" json:"env,omitempty"`
	TeamId    string     `protobuf:"bytes,2,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	NetConfig *NetConfig `protobuf:"bytes,3,opt,name=net_config,json=netConfig" json:"net_config,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Metadata) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *Metadata) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Metadata) GetNetConfig() *NetConfig {
	if m != nil {
		return m.NetConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusResponse)(nil), "appscode.health.StatusResponse")
	proto.RegisterType((*URLBase)(nil), "appscode.health.URLBase")
	proto.RegisterType((*NetConfig)(nil), "appscode.health.NetConfig")
	proto.RegisterType((*Metadata)(nil), "appscode.health.Metadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/appscode.health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *appscode_dtypes.VoidRequest) (*StatusResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.health.Health/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Status(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Health_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}

func init() { proto.RegisterFile("health.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x95, 0x04, 0x36, 0xc9, 0xa4, 0xd0, 0xca, 0x12, 0x24, 0x84, 0x8a, 0xa2, 0xe5, 0xc2,
	0x29, 0x8b, 0x5a, 0xf5, 0x50, 0x21, 0x21, 0x48, 0x25, 0x04, 0x52, 0x40, 0x95, 0xab, 0xf4, 0xc0,
	0x65, 0xe5, 0xac, 0xa7, 0x89, 0xa5, 0x8d, 0xbd, 0xac, 0xbd, 0x91, 0xb8, 0xf6, 0x15, 0xb8, 0xf0,
	0x12, 0x1c, 0x79, 0x12, 0x5e, 0x81, 0x07, 0x41, 0x6b, 0x7b, 0x53, 0x96, 0xa8, 0xda, 0xcb, 0xfe,
	0x19, 0x7f, 0xbf, 0x6f, 0xc6, 0x1e, 0x0f, 0xec, 0xad, 0x90, 0xa5, 0x66, 0x35, 0xc9, 0x72, 0x65,
	0x14, 0xd9, 0x67, 0x59, 0xa6, 0x13, 0xc5, 0x71, 0xe2, 0xc2, 0xe3, 0xc3, 0xa5, 0x52, 0xcb, 0x14,
	0x23, 0x96, 0x89, 0x88, 0x49, 0xa9, 0x0c, 0x33, 0x42, 0x49, 0xed, 0xe4, 0xe3, 0x67, 0x95, 0xfc,
	0x8e, 0xf5, 0xa3, 0xda, 0x3a, 0x37, 0xdf, 0x32, 0xd4, 0x91, 0x7d, 0x7a, 0x41, 0x58, 0x13, 0x6c,
	0x30, 0xd7, 0x42, 0xc9, 0xea, 0xed, 0x34, 0xe1, 0xcf, 0x16, 0x3c, 0xbc, 0x34, 0xcc, 0x14, 0x9a,
	0xa2, 0xce, 0x94, 0xd4, 0x48, 0x22, 0x08, 0xb4, 0x8d, 0x8c, 0x5a, 0xcf, 0x5b, 0x2f, 0x07, 0xc7,
	0xc3, 0xc9, 0xb6, 0x6e, 0x97, 0x64, 0xe2, 0x01, 0x2f, 0x23, 0x27, 0xd0, 0xf5, 0xa6, 0xa3, 0xb6,
	0x25, 0x9e, 0xdc, 0x12, 0x55, 0xb6, 0x2b, 0xf7, 0xa6, 0x95, 0x92, 0x9c, 0x42, 0x6f, 0x8d, 0x86,
	0x71, 0x66, 0xd8, 0xa8, 0xf3, 0x3f, 0xe5, 0x8f, 0xed, 0x93, 0x17, 0xd0, 0xad, 0x34, 0x7c, 0x03,
	0xdd, 0x39, 0x9d, 0x4d, 0x99, 0x46, 0xf2, 0x18, 0x02, 0x9d, 0xac, 0x70, 0x8d, 0xb6, 0xce, 0x3e,
	0xf5, 0x7f, 0xe4, 0x29, 0xf4, 0x17, 0x4c, 0x63, 0xcc, 0x38, 0xcf, 0x6d, 0x41, 0x7d, 0xda, 0x2b,
	0x03, 0xef, 0x38, 0xcf, 0xc3, 0x1f, 0x1d, 0xe8, 0x7f, 0x46, 0x73, 0xae, 0xe4, 0xb5, 0x58, 0x96,
	0x52, 0x83, 0x6c, 0xed, 0xa4, 0xce, 0xa5, 0x57, 0x06, 0x4a, 0x29, 0x39, 0x83, 0x41, 0x56, 0x2c,
	0x52, 0x91, 0xc4, 0x45, 0x9e, 0x6a, 0xbf, 0xb5, 0xd1, 0x4e, 0x91, 0xbe, 0x1c, 0x0a, 0x4e, 0x3c,
	0xcf, 0x53, 0x4d, 0x4e, 0xbd, 0xaf, 0x05, 0x3b, 0x0d, 0xa0, 0xcd, 0x68, 0xb1, 0xd7, 0xb0, 0x97,
	0xa4, 0x85, 0x36, 0x98, 0x3b, 0xf2, 0x5e, 0x03, 0x39, 0xf0, 0x6a, 0x0b, 0xbf, 0x85, 0x7d, 0x21,
	0xe3, 0x1a, 0x7f, 0xbf, 0x81, 0x7f, 0x20, 0xe4, 0xf9, 0x3f, 0x0e, 0xef, 0x81, 0xcc, 0xe9, 0x2c,
	0xd6, 0x2b, 0x95, 0x1b, 0x94, 0x95, 0x49, 0xd0, 0x60, 0x72, 0x30, 0xa7, 0xb3, 0xcb, 0x0a, 0xa9,
	0x76, 0x7f, 0x2d, 0x52, 0x74, 0x78, 0xb7, 0x69, 0xf7, 0xa5, 0xb4, 0xc4, 0xc2, 0x0c, 0x7a, 0x55,
	0xc3, 0xc9, 0x01, 0x74, 0x50, 0x6e, 0x7c, 0x4b, 0xca, 0x4f, 0x32, 0x84, 0xae, 0x3d, 0x52, 0xc1,
	0x7d, 0x4f, 0x83, 0xf2, 0xf7, 0x23, 0x27, 0x67, 0x00, 0x12, 0x4d, 0x9c, 0xd8, 0x8e, 0xfa, 0xc3,
	0x1e, 0xef, 0xa4, 0xdb, 0xf6, 0x9c, 0xf6, 0x65, 0xf5, 0x79, 0xbc, 0x81, 0xe0, 0x83, 0x5d, 0x26,
	0x29, 0x04, 0xee, 0x52, 0x93, 0xc3, 0x9d, 0xdb, 0x7e, 0xa5, 0x04, 0xa7, 0xf8, 0xb5, 0x40, 0x6d,
	0xc6, 0x47, 0x3b, 0xc6, 0xf5, 0xe1, 0x09, 0x5f, 0xdc, 0xfc, 0x1a, 0xb5, 0x7b, 0xad, 0x9b, 0xdf,
	0x7f, 0xbe, 0xb7, 0x87, 0xe4, 0x51, 0x14, 0xd7, 0x66, 0xd0, 0x31, 0xd3, 0x57, 0x30, 0x4c, 0xd4,
	0xfa, 0xd6, 0x8a, 0x65, 0xc2, 0xdb, 0x4d, 0x07, 0xae, 0xa0, 0x8b, 0x72, 0x38, 0x2f, 0x5a, 0x5f,
	0x02, 0x17, 0x5e, 0x04, 0x76, 0x5a, 0x4f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x1a, 0xa4,
	0xa2, 0x51, 0x04, 0x00, 0x00,
}
