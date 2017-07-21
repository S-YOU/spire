// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/control_plane_ca/proto/control_plane_ca.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	plugins/control_plane_ca/proto/control_plane_ca.proto

It has these top-level messages:
	SignCsrRequest
	SignCsrResponse
	GenerateCsrRequest
	GenerateCsrResponse
	FetchCertificateRequest
	FetchCertificateResponse
	LoadCertificateRequest
	LoadCertificateResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SignCsrRequest struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *SignCsrRequest) Reset()                    { *m = SignCsrRequest{} }
func (m *SignCsrRequest) String() string            { return proto1.CompactTextString(m) }
func (*SignCsrRequest) ProtoMessage()               {}
func (*SignCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignCsrRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type SignCsrResponse struct {
	SignedCertificate []byte `protobuf:"bytes,1,opt,name=signedCertificate,proto3" json:"signedCertificate,omitempty"`
}

func (m *SignCsrResponse) Reset()                    { *m = SignCsrResponse{} }
func (m *SignCsrResponse) String() string            { return proto1.CompactTextString(m) }
func (*SignCsrResponse) ProtoMessage()               {}
func (*SignCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignCsrResponse) GetSignedCertificate() []byte {
	if m != nil {
		return m.SignedCertificate
	}
	return nil
}

type GenerateCsrRequest struct {
}

func (m *GenerateCsrRequest) Reset()                    { *m = GenerateCsrRequest{} }
func (m *GenerateCsrRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenerateCsrRequest) ProtoMessage()               {}
func (*GenerateCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GenerateCsrResponse struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *GenerateCsrResponse) Reset()                    { *m = GenerateCsrResponse{} }
func (m *GenerateCsrResponse) String() string            { return proto1.CompactTextString(m) }
func (*GenerateCsrResponse) ProtoMessage()               {}
func (*GenerateCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenerateCsrResponse) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type FetchCertificateRequest struct {
}

func (m *FetchCertificateRequest) Reset()                    { *m = FetchCertificateRequest{} }
func (m *FetchCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchCertificateRequest) ProtoMessage()               {}
func (*FetchCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type FetchCertificateResponse struct {
	StoredIntermediateCert []byte `protobuf:"bytes,1,opt,name=storedIntermediateCert,proto3" json:"storedIntermediateCert,omitempty"`
}

func (m *FetchCertificateResponse) Reset()                    { *m = FetchCertificateResponse{} }
func (m *FetchCertificateResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchCertificateResponse) ProtoMessage()               {}
func (*FetchCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchCertificateResponse) GetStoredIntermediateCert() []byte {
	if m != nil {
		return m.StoredIntermediateCert
	}
	return nil
}

type LoadCertificateRequest struct {
	SignedIntermediateCert []byte `protobuf:"bytes,1,opt,name=signedIntermediateCert,proto3" json:"signedIntermediateCert,omitempty"`
}

func (m *LoadCertificateRequest) Reset()                    { *m = LoadCertificateRequest{} }
func (m *LoadCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*LoadCertificateRequest) ProtoMessage()               {}
func (*LoadCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoadCertificateRequest) GetSignedIntermediateCert() []byte {
	if m != nil {
		return m.SignedIntermediateCert
	}
	return nil
}

type LoadCertificateResponse struct {
}

func (m *LoadCertificateResponse) Reset()                    { *m = LoadCertificateResponse{} }
func (m *LoadCertificateResponse) String() string            { return proto1.CompactTextString(m) }
func (*LoadCertificateResponse) ProtoMessage()               {}
func (*LoadCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto1.RegisterType((*SignCsrRequest)(nil), "proto.SignCsrRequest")
	proto1.RegisterType((*SignCsrResponse)(nil), "proto.SignCsrResponse")
	proto1.RegisterType((*GenerateCsrRequest)(nil), "proto.GenerateCsrRequest")
	proto1.RegisterType((*GenerateCsrResponse)(nil), "proto.GenerateCsrResponse")
	proto1.RegisterType((*FetchCertificateRequest)(nil), "proto.FetchCertificateRequest")
	proto1.RegisterType((*FetchCertificateResponse)(nil), "proto.FetchCertificateResponse")
	proto1.RegisterType((*LoadCertificateRequest)(nil), "proto.LoadCertificateRequest")
	proto1.RegisterType((*LoadCertificateResponse)(nil), "proto.LoadCertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ControlPlaneCA service

type ControlPlaneCAClient interface {
	SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error)
	GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error)
	FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error)
	LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error)
}

type controlPlaneCAClient struct {
	cc *grpc.ClientConn
}

func NewControlPlaneCAClient(cc *grpc.ClientConn) ControlPlaneCAClient {
	return &controlPlaneCAClient{cc}
}

func (c *controlPlaneCAClient) SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error) {
	out := new(SignCsrResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/SignCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error) {
	out := new(GenerateCsrResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/GenerateCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error) {
	out := new(FetchCertificateResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/FetchCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error) {
	out := new(LoadCertificateResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/LoadCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ControlPlaneCA service

type ControlPlaneCAServer interface {
	SignCsr(context.Context, *SignCsrRequest) (*SignCsrResponse, error)
	GenerateCsr(context.Context, *GenerateCsrRequest) (*GenerateCsrResponse, error)
	FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error)
	LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error)
}

func RegisterControlPlaneCAServer(s *grpc.Server, srv ControlPlaneCAServer) {
	s.RegisterService(&_ControlPlaneCA_serviceDesc, srv)
}

func _ControlPlaneCA_SignCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/SignCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, req.(*SignCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_GenerateCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/GenerateCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, req.(*GenerateCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_FetchCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/FetchCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, req.(*FetchCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_LoadCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/LoadCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, req.(*LoadCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlPlaneCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ControlPlaneCA",
	HandlerType: (*ControlPlaneCAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignCsr",
			Handler:    _ControlPlaneCA_SignCsr_Handler,
		},
		{
			MethodName: "GenerateCsr",
			Handler:    _ControlPlaneCA_GenerateCsr_Handler,
		},
		{
			MethodName: "FetchCertificate",
			Handler:    _ControlPlaneCA_FetchCertificate_Handler,
		},
		{
			MethodName: "LoadCertificate",
			Handler:    _ControlPlaneCA_LoadCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins/control_plane_ca/proto/control_plane_ca.proto",
}

func init() {
	proto1.RegisterFile("plugins/control_plane_ca/proto/control_plane_ca.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0xa2, 0xc2, 0x28, 0x6d, 0x5d, 0x35, 0xb6, 0x01, 0xab, 0xec, 0x45, 0x0f, 0xd2,
	0x82, 0xa2, 0x78, 0x13, 0x89, 0x28, 0x82, 0x87, 0x90, 0x3e, 0x40, 0x89, 0xc9, 0x18, 0x17, 0xe2,
	0x6e, 0xdc, 0xdd, 0x3e, 0x9b, 0xaf, 0x27, 0x49, 0x46, 0x6d, 0xb3, 0xc9, 0x29, 0x61, 0xfe, 0x99,
	0x6f, 0x76, 0xfe, 0x1f, 0x6e, 0x8a, 0x7c, 0x99, 0x09, 0x69, 0x66, 0x89, 0x92, 0x56, 0xab, 0x7c,
	0x51, 0xe4, 0xb1, 0xc4, 0x45, 0x12, 0xcf, 0x0a, 0xad, 0xac, 0x72, 0xca, 0xd3, 0xaa, 0xcc, 0xb6,
	0xaa, 0x0f, 0xe7, 0xd0, 0x9f, 0x8b, 0x4c, 0x06, 0x46, 0x47, 0xf8, 0xb5, 0x44, 0x63, 0xd9, 0x10,
	0x36, 0x13, 0xa3, 0x47, 0xbd, 0xb3, 0xde, 0xc5, 0x5e, 0x54, 0xfe, 0xf2, 0x7b, 0x18, 0xfc, 0xf5,
	0x98, 0x42, 0x49, 0x83, 0xec, 0x12, 0xf6, 0x8d, 0xc8, 0x24, 0xa6, 0x01, 0x6a, 0x2b, 0xde, 0x45,
	0x12, 0x5b, 0xa4, 0x11, 0x57, 0xe0, 0x87, 0xc0, 0x9e, 0x51, 0xa2, 0x8e, 0x2d, 0xfe, 0x2f, 0xe2,
	0xe7, 0x70, 0xb0, 0x56, 0x25, 0xb4, 0xbb, 0x7f, 0x0c, 0xc7, 0x4f, 0x68, 0x93, 0x8f, 0x15, 0xe4,
	0x2f, 0x23, 0x82, 0x91, 0x2b, 0x11, 0xe8, 0x16, 0x3c, 0x63, 0x95, 0xc6, 0xf4, 0x45, 0x5a, 0xd4,
	0x9f, 0x98, 0x8a, 0x72, 0x13, 0x6a, 0x4b, 0xec, 0x0e, 0x95, 0x87, 0xe0, 0xbd, 0xaa, 0x38, 0x75,
	0xb7, 0x55, 0xc4, 0xea, 0xb8, 0x4e, 0x62, 0xab, 0x5a, 0x1e, 0xe0, 0x10, 0xeb, 0x47, 0x5e, 0x7d,
	0x6f, 0x40, 0x3f, 0xa8, 0x13, 0x0a, 0xcb, 0x80, 0x82, 0x07, 0x76, 0x07, 0x3b, 0x64, 0x37, 0x3b,
	0xaa, 0xc3, 0x9a, 0xae, 0x47, 0xe4, 0x7b, 0xcd, 0x32, 0x5d, 0xfc, 0x08, 0xbb, 0x2b, 0x8e, 0xb2,
	0x31, 0xb5, 0xb9, 0xde, 0xfb, 0x7e, 0x9b, 0x44, 0x94, 0x39, 0x0c, 0x9b, 0x9e, 0xb2, 0x09, 0xf5,
	0x77, 0xe4, 0xe0, 0x9f, 0x76, 0xea, 0x04, 0x0d, 0x61, 0xd0, 0xb0, 0x80, 0x9d, 0xd0, 0x4c, 0xbb,
	0xd9, 0xfe, 0xa4, 0x4b, 0xae, 0x89, 0x6f, 0xdb, 0x95, 0x7c, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x29, 0x65, 0xb3, 0x02, 0x00, 0x03, 0x00, 0x00,
}
