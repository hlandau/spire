// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workload.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	workload.proto

It has these top-level messages:
	WLSVIDEntry
	FederateEntry
	FetchSVIDBundleRequest
	FetchSVIDBundleResponse
	FetchSVIDBundlesResponse
	FetchFederatedBundleRequest
	FetchFederatedBundleResponse
	FetchFederatedBundlesResponse
	Empty
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// * A WLSVIDEntry represents a Workload's SVID and its associated information
type WLSVIDEntry struct {
	SpiffeId   string `protobuf:"bytes,1,opt,name=spiffeId" json:"spiffeId,omitempty"`
	Svid       []byte `protobuf:"bytes,2,opt,name=svid,proto3" json:"svid,omitempty"`
	PrivateKey []byte `protobuf:"bytes,3,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	Ttl        int32  `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *WLSVIDEntry) Reset()                    { *m = WLSVIDEntry{} }
func (m *WLSVIDEntry) String() string            { return proto.CompactTextString(m) }
func (*WLSVIDEntry) ProtoMessage()               {}
func (*WLSVIDEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WLSVIDEntry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *WLSVIDEntry) GetSvid() []byte {
	if m != nil {
		return m.Svid
	}
	return nil
}

func (m *WLSVIDEntry) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *WLSVIDEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

// * represents cert bundle of a remote control plane for the purposes of trusting remote workloads
type FederateEntry struct {
	SpiffeId      string `protobuf:"bytes,1,opt,name=spiffeId" json:"spiffeId,omitempty"`
	CaTrustBundle []byte `protobuf:"bytes,2,opt,name=caTrustBundle,proto3" json:"caTrustBundle,omitempty"`
	Ttl           int32  `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *FederateEntry) Reset()                    { *m = FederateEntry{} }
func (m *FederateEntry) String() string            { return proto.CompactTextString(m) }
func (*FederateEntry) ProtoMessage()               {}
func (*FederateEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FederateEntry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *FederateEntry) GetCaTrustBundle() []byte {
	if m != nil {
		return m.CaTrustBundle
	}
	return nil
}

func (m *FederateEntry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

// * represents a workload request for a SVID and the control plane's cert bundle of a specific SPIFFEID
type FetchSVIDBundleRequest struct {
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffeId" json:"spiffeId,omitempty"`
}

func (m *FetchSVIDBundleRequest) Reset()                    { *m = FetchSVIDBundleRequest{} }
func (m *FetchSVIDBundleRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchSVIDBundleRequest) ProtoMessage()               {}
func (*FetchSVIDBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchSVIDBundleRequest) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

// * represents a response specific to the requesting workload SPIFFEId,
// Includes the workload's SVID Entry(SVID and its corresponding information )
// and the Control Plane's trusted cert bundle
type FetchSVIDBundleResponse struct {
	WLSVIDEntry *WLSVIDEntry `protobuf:"bytes,1,opt,name=WLSVIDEntry" json:"WLSVIDEntry,omitempty"`
	CpBundle    []byte       `protobuf:"bytes,2,opt,name=cpBundle,proto3" json:"cpBundle,omitempty"`
}

func (m *FetchSVIDBundleResponse) Reset()                    { *m = FetchSVIDBundleResponse{} }
func (m *FetchSVIDBundleResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchSVIDBundleResponse) ProtoMessage()               {}
func (*FetchSVIDBundleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FetchSVIDBundleResponse) GetWLSVIDEntry() *WLSVIDEntry {
	if m != nil {
		return m.WLSVIDEntry
	}
	return nil
}

func (m *FetchSVIDBundleResponse) GetCpBundle() []byte {
	if m != nil {
		return m.CpBundle
	}
	return nil
}

// * represents response the includes all the SVIDs the and Control Plane's trusted cert bundle workload
type FetchSVIDBundlesResponse struct {
	WLSVIDEntry []*WLSVIDEntry `protobuf:"bytes,1,rep,name=WLSVIDEntry" json:"WLSVIDEntry,omitempty"`
	CpBundle    []byte         `protobuf:"bytes,2,opt,name=cpBundle,proto3" json:"cpBundle,omitempty"`
}

func (m *FetchSVIDBundlesResponse) Reset()                    { *m = FetchSVIDBundlesResponse{} }
func (m *FetchSVIDBundlesResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchSVIDBundlesResponse) ProtoMessage()               {}
func (*FetchSVIDBundlesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FetchSVIDBundlesResponse) GetWLSVIDEntry() []*WLSVIDEntry {
	if m != nil {
		return m.WLSVIDEntry
	}
	return nil
}

func (m *FetchSVIDBundlesResponse) GetCpBundle() []byte {
	if m != nil {
		return m.CpBundle
	}
	return nil
}

// * represents a Federated cert Bundle request corresponding to a specific SPIFFEId
type FetchFederatedBundleRequest struct {
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffeId" json:"spiffeId,omitempty"`
}

func (m *FetchFederatedBundleRequest) Reset()                    { *m = FetchFederatedBundleRequest{} }
func (m *FetchFederatedBundleRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchFederatedBundleRequest) ProtoMessage()               {}
func (*FetchFederatedBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchFederatedBundleRequest) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

// * represents cert Bundles that a specific workload's SPIFFEId is registered to trust
type FetchFederatedBundleResponse struct {
	FederateEntry []*FederateEntry `protobuf:"bytes,1,rep,name=FederateEntry" json:"FederateEntry,omitempty"`
}

func (m *FetchFederatedBundleResponse) Reset()                    { *m = FetchFederatedBundleResponse{} }
func (m *FetchFederatedBundleResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchFederatedBundleResponse) ProtoMessage()               {}
func (*FetchFederatedBundleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FetchFederatedBundleResponse) GetFederateEntry() []*FederateEntry {
	if m != nil {
		return m.FederateEntry
	}
	return nil
}

// * represents all the cert Bundles that a workload is registered to trust
type FetchFederatedBundlesResponse struct {
	FederateEntry []*FederateEntry `protobuf:"bytes,1,rep,name=FederateEntry" json:"FederateEntry,omitempty"`
}

func (m *FetchFederatedBundlesResponse) Reset()                    { *m = FetchFederatedBundlesResponse{} }
func (m *FetchFederatedBundlesResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchFederatedBundlesResponse) ProtoMessage()               {}
func (*FetchFederatedBundlesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FetchFederatedBundlesResponse) GetFederateEntry() []*FederateEntry {
	if m != nil {
		return m.FederateEntry
	}
	return nil
}

// *
// @exclude Represents a message with no fields
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*WLSVIDEntry)(nil), "pb.WLSVIDEntry")
	proto.RegisterType((*FederateEntry)(nil), "pb.FederateEntry")
	proto.RegisterType((*FetchSVIDBundleRequest)(nil), "pb.FetchSVIDBundleRequest")
	proto.RegisterType((*FetchSVIDBundleResponse)(nil), "pb.FetchSVIDBundleResponse")
	proto.RegisterType((*FetchSVIDBundlesResponse)(nil), "pb.FetchSVIDBundlesResponse")
	proto.RegisterType((*FetchFederatedBundleRequest)(nil), "pb.FetchFederatedBundleRequest")
	proto.RegisterType((*FetchFederatedBundleResponse)(nil), "pb.FetchFederatedBundleResponse")
	proto.RegisterType((*FetchFederatedBundlesResponse)(nil), "pb.FetchFederatedBundlesResponse")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Workload service

type WorkloadClient interface {
	// /Requests SVID and cert bundle of the control plane corresponding to a specific SPIFFEId(useful for rotation)
	FetchSVIDBundle(ctx context.Context, in *FetchSVIDBundleRequest, opts ...grpc.CallOption) (*FetchSVIDBundleResponse, error)
	// /Requests all SVIDs and cert bundle of the control plane associated with the workload
	FetchSVIDBundles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FetchSVIDBundlesResponse, error)
	// /Requests trusted external CA cert bundles corresponding to a specific SPIFFEId (useful for rotation)
	FetchFederatedBundle(ctx context.Context, in *FetchFederatedBundleRequest, opts ...grpc.CallOption) (*FetchFederatedBundleResponse, error)
	// /Requests all trusted external CA cert bundles associated with the workload
	FetchFederatedBundles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FetchFederatedBundlesResponse, error)
}

type workloadClient struct {
	cc *grpc.ClientConn
}

func NewWorkloadClient(cc *grpc.ClientConn) WorkloadClient {
	return &workloadClient{cc}
}

func (c *workloadClient) FetchSVIDBundle(ctx context.Context, in *FetchSVIDBundleRequest, opts ...grpc.CallOption) (*FetchSVIDBundleResponse, error) {
	out := new(FetchSVIDBundleResponse)
	err := grpc.Invoke(ctx, "/pb.Workload/FetchSVIDBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadClient) FetchSVIDBundles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FetchSVIDBundlesResponse, error) {
	out := new(FetchSVIDBundlesResponse)
	err := grpc.Invoke(ctx, "/pb.Workload/FetchSVIDBundles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadClient) FetchFederatedBundle(ctx context.Context, in *FetchFederatedBundleRequest, opts ...grpc.CallOption) (*FetchFederatedBundleResponse, error) {
	out := new(FetchFederatedBundleResponse)
	err := grpc.Invoke(ctx, "/pb.Workload/FetchFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadClient) FetchFederatedBundles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FetchFederatedBundlesResponse, error) {
	out := new(FetchFederatedBundlesResponse)
	err := grpc.Invoke(ctx, "/pb.Workload/FetchFederatedBundles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Workload service

type WorkloadServer interface {
	// /Requests SVID and cert bundle of the control plane corresponding to a specific SPIFFEId(useful for rotation)
	FetchSVIDBundle(context.Context, *FetchSVIDBundleRequest) (*FetchSVIDBundleResponse, error)
	// /Requests all SVIDs and cert bundle of the control plane associated with the workload
	FetchSVIDBundles(context.Context, *Empty) (*FetchSVIDBundlesResponse, error)
	// /Requests trusted external CA cert bundles corresponding to a specific SPIFFEId (useful for rotation)
	FetchFederatedBundle(context.Context, *FetchFederatedBundleRequest) (*FetchFederatedBundleResponse, error)
	// /Requests all trusted external CA cert bundles associated with the workload
	FetchFederatedBundles(context.Context, *Empty) (*FetchFederatedBundlesResponse, error)
}

func RegisterWorkloadServer(s *grpc.Server, srv WorkloadServer) {
	s.RegisterService(&_Workload_serviceDesc, srv)
}

func _Workload_FetchSVIDBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchSVIDBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadServer).FetchSVIDBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Workload/FetchSVIDBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadServer).FetchSVIDBundle(ctx, req.(*FetchSVIDBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workload_FetchSVIDBundles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadServer).FetchSVIDBundles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Workload/FetchSVIDBundles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadServer).FetchSVIDBundles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workload_FetchFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchFederatedBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadServer).FetchFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Workload/FetchFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadServer).FetchFederatedBundle(ctx, req.(*FetchFederatedBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workload_FetchFederatedBundles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadServer).FetchFederatedBundles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Workload/FetchFederatedBundles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadServer).FetchFederatedBundles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Workload",
	HandlerType: (*WorkloadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSVIDBundle",
			Handler:    _Workload_FetchSVIDBundle_Handler,
		},
		{
			MethodName: "FetchSVIDBundles",
			Handler:    _Workload_FetchSVIDBundles_Handler,
		},
		{
			MethodName: "FetchFederatedBundle",
			Handler:    _Workload_FetchFederatedBundle_Handler,
		},
		{
			MethodName: "FetchFederatedBundles",
			Handler:    _Workload_FetchFederatedBundles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workload.proto",
}

func init() { proto.RegisterFile("workload.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdf, 0x4f, 0xc2, 0x30,
	0x10, 0xce, 0x18, 0x28, 0x1c, 0x22, 0xd8, 0xf8, 0x63, 0x19, 0xa8, 0x73, 0xf1, 0x81, 0x27, 0x12,
	0xd1, 0xc4, 0x18, 0xdf, 0x8c, 0x90, 0xa0, 0x3e, 0x4d, 0x23, 0xfa, 0x38, 0xb6, 0x12, 0x16, 0x91,
	0xd5, 0xb6, 0x60, 0xf8, 0xc3, 0xfc, 0xff, 0xcc, 0x4a, 0x19, 0x6c, 0x16, 0x35, 0xe1, 0xad, 0xbd,
	0xbb, 0xef, 0xbe, 0xef, 0xbe, 0x6b, 0x0a, 0xdb, 0x9f, 0x21, 0x7d, 0x1b, 0x86, 0xae, 0xdf, 0x20,
	0x34, 0xe4, 0x21, 0xca, 0x90, 0x9e, 0x1d, 0x42, 0xb1, 0xfb, 0xf0, 0xf8, 0xdc, 0xb9, 0x6d, 0x8d,
	0x38, 0x9d, 0x22, 0x13, 0xf2, 0x8c, 0x04, 0xfd, 0x3e, 0xee, 0xf8, 0x86, 0x66, 0x69, 0xf5, 0x82,
	0x13, 0xdf, 0x11, 0x82, 0x2c, 0x9b, 0x04, 0xbe, 0x91, 0xb1, 0xb4, 0xfa, 0x96, 0x23, 0xce, 0xe8,
	0x08, 0x80, 0xd0, 0x60, 0xe2, 0x72, 0x7c, 0x8f, 0xa7, 0x86, 0x2e, 0x32, 0x4b, 0x11, 0x54, 0x01,
	0x9d, 0xf3, 0xa1, 0x91, 0xb5, 0xb4, 0x7a, 0xce, 0x89, 0x8e, 0xb6, 0x07, 0xa5, 0x36, 0xf6, 0x31,
	0x75, 0x39, 0xfe, 0x9b, 0xf2, 0x14, 0x4a, 0x9e, 0xfb, 0x44, 0xc7, 0x8c, 0xdf, 0x8c, 0x47, 0xfe,
	0x10, 0x4b, 0xee, 0x64, 0x70, 0x4e, 0xa2, 0x2f, 0x48, 0x2e, 0x60, 0xbf, 0x8d, 0xb9, 0x37, 0x88,
	0x06, 0x9b, 0x15, 0x39, 0xf8, 0x63, 0x8c, 0x19, 0xff, 0x8d, 0xcd, 0x1e, 0xc0, 0xc1, 0x0f, 0x14,
	0x23, 0xe1, 0x88, 0x61, 0x74, 0x96, 0xb0, 0x49, 0x20, 0x8b, 0xcd, 0x72, 0x83, 0xf4, 0x1a, 0x4b,
	0x61, 0x27, 0x6d, 0xa5, 0x47, 0x12, 0xb2, 0xe3, 0xbb, 0x1d, 0x80, 0x91, 0x62, 0x62, 0xab, 0xa9,
	0xf4, 0xb5, 0xa8, 0xae, 0xa0, 0x2a, 0xa8, 0xe6, 0xa6, 0xfb, 0xff, 0xf7, 0xa3, 0x0b, 0x35, 0x35,
	0x54, 0x2a, 0xbd, 0x4c, 0xad, 0x52, 0x6a, 0xdd, 0x89, 0xb4, 0x26, 0x12, 0x4e, 0xb2, 0xce, 0x7e,
	0x81, 0x43, 0x55, 0x63, 0xb6, 0x7e, 0xe7, 0x4d, 0xc8, 0xb5, 0xde, 0x09, 0x9f, 0x36, 0xbf, 0x32,
	0x90, 0xef, 0xca, 0xe7, 0x8e, 0xee, 0xa0, 0x9c, 0xb2, 0x1b, 0x99, 0xb3, 0x56, 0xaa, 0x37, 0x62,
	0x56, 0x95, 0x39, 0x29, 0xed, 0x1a, 0x2a, 0xe9, 0xd5, 0xa1, 0x42, 0x04, 0x10, 0xbc, 0x66, 0x4d,
	0x81, 0x5d, 0xcc, 0xf5, 0x0a, 0xbb, 0xaa, 0xc1, 0xd1, 0x71, 0x8c, 0x52, 0xaf, 0xc9, 0xb4, 0x56,
	0x17, 0xc8, 0xd6, 0x2d, 0xd8, 0x53, 0x7a, 0xba, 0x2c, 0xee, 0x64, 0x55, 0x97, 0x58, 0x61, 0x6f,
	0x43, 0x7c, 0x0d, 0xe7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x09, 0x0e, 0xe9, 0x2c, 0x04,
	0x00, 0x00,
}
