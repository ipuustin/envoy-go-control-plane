// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/load_stats/v3/lrs.proto

package envoy_service_load_stats_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LoadStatsRequest struct {
	Node                 *v3.Node            `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClusterStats         []*v31.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a87f04ee5f8fca3, []int{0}
}

func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsRequest.Unmarshal(m, b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
}
func (m *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(m, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return xxx_messageInfo_LoadStatsRequest.Size(m)
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *v3.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*v31.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

type LoadStatsResponse struct {
	Clusters                  []string           `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	LoadReportingInterval     *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	ReportEndpointGranularity bool               `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a87f04ee5f8fca3, []int{1}
}

func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsResponse.Unmarshal(m, b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
}
func (m *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(m, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return xxx_messageInfo_LoadStatsResponse.Size(m)
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func (m *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if m != nil {
		return m.ReportEndpointGranularity
	}
	return false
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v3.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v3.LoadStatsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v3/lrs.proto", fileDescriptor_4a87f04ee5f8fca3)
}

var fileDescriptor_4a87f04ee5f8fca3 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xf1, 0x1d, 0x54, 0xc1, 0x05, 0x51, 0x22, 0x50, 0xd3, 0x54, 0x94, 0xe3, 0x04, 0x28,
	0x20, 0x6a, 0xa3, 0x0b, 0x53, 0x07, 0x86, 0x00, 0x42, 0x88, 0x13, 0x2a, 0xb9, 0x0f, 0x70, 0xf2,
	0x25, 0x6e, 0x64, 0x29, 0xf8, 0x05, 0xdb, 0x89, 0xb8, 0x0d, 0x31, 0x75, 0x62, 0x60, 0xe4, 0xa3,
	0x30, 0x83, 0xc4, 0xca, 0xd7, 0x61, 0x42, 0x89, 0x9d, 0xf6, 0x4a, 0xd5, 0xaa, 0x5b, 0xec, 0xf7,
	0xfb, 0xff, 0xfd, 0xde, 0xff, 0x05, 0x3f, 0xe0, 0xb2, 0x81, 0x25, 0xd5, 0x5c, 0x35, 0x22, 0xe3,
	0xb4, 0x04, 0x96, 0xcf, 0xb5, 0x61, 0x46, 0xd3, 0x26, 0xa6, 0xa5, 0xd2, 0xa4, 0x52, 0x60, 0xc0,
	0xdf, 0xee, 0x30, 0xe2, 0x30, 0x72, 0x8c, 0x91, 0x26, 0x0e, 0xef, 0x5a, 0x8f, 0x0c, 0xe4, 0x81,
	0x28, 0x68, 0x06, 0x8a, 0xb7, 0xe2, 0x05, 0xd3, 0xdc, 0xaa, 0xc3, 0xc7, 0x27, 0x00, 0x2e, 0xf3,
	0x0a, 0x84, 0x34, 0xdd, 0x0b, 0xad, 0x91, 0xe2, 0x15, 0x28, 0xe3, 0xd8, 0x9d, 0x02, 0xa0, 0x28,
	0x39, 0xed, 0x4e, 0x8b, 0xfa, 0x80, 0xe6, 0xb5, 0x62, 0x46, 0x80, 0x74, 0xf5, 0x3b, 0x75, 0x5e,
	0x31, 0xca, 0xa4, 0x04, 0xd3, 0x5d, 0x6b, 0xda, 0xf6, 0x51, 0xbb, 0x46, 0xc3, 0x7b, 0xa7, 0xca,
	0x0d, 0x57, 0x5a, 0x80, 0x14, 0xb2, 0x70, 0xc8, 0x66, 0xc3, 0x4a, 0x91, 0x33, 0xc3, 0x69, 0xff,
	0x61, 0x0b, 0xe3, 0x9f, 0x08, 0x6f, 0x4c, 0x81, 0xe5, 0xb3, 0x76, 0xb0, 0x94, 0x7f, 0xac, 0xb9,
	0x36, 0x3e, 0xc1, 0x97, 0x25, 0xe4, 0x3c, 0x40, 0x23, 0x14, 0xad, 0x4f, 0x42, 0x62, 0x83, 0xb0,
	0xa3, 0x90, 0x76, 0x56, 0xd2, 0xc4, 0xe4, 0x1d, 0xe4, 0x3c, 0xed, 0x38, 0xff, 0x2d, 0xbe, 0x9e,
	0x95, 0xb5, 0x36, 0x5c, 0xd9, 0x80, 0x82, 0xc1, 0x68, 0x18, 0xad, 0x4f, 0x1e, 0x9e, 0x14, 0xf6,
	0x19, 0xb4, 0xe2, 0x17, 0x16, 0xb7, 0xaf, 0x5e, 0xcb, 0x56, 0x4e, 0x7b, 0xf1, 0xf7, 0x5f, 0x87,
	0x3b, 0x04, 0x3f, 0x39, 0x3b, 0xfd, 0x09, 0xf9, 0xbf, 0xe3, 0xf1, 0x97, 0x01, 0xbe, 0xb9, 0x72,
	0xa9, 0x2b, 0x90, 0x9a, 0xfb, 0xf7, 0xb1, 0xe7, 0xac, 0x75, 0x80, 0x46, 0xc3, 0xe8, 0x6a, 0xe2,
	0xfd, 0x4d, 0xae, 0x7c, 0x43, 0x03, 0x0f, 0xa5, 0x47, 0x15, 0xff, 0x3d, 0xde, 0x5c, 0x59, 0x89,
	0x90, 0xc5, 0x5c, 0x48, 0xc3, 0x55, 0xc3, 0xca, 0x60, 0xd0, 0x05, 0xb0, 0x45, 0xec, 0x7e, 0x48,
	0xbf, 0x1f, 0xf2, 0xd2, 0xed, 0x27, 0xbd, 0xdd, 0x2a, 0xd3, 0x5e, 0xf8, 0xc6, 0xe9, 0xfc, 0xe7,
	0x78, 0xdb, 0xba, 0xcd, 0xfb, 0xa1, 0xe7, 0x85, 0x62, 0xb2, 0x2e, 0x99, 0x12, 0x66, 0x19, 0x0c,
	0x47, 0x28, 0xf2, 0xd2, 0x2d, 0x8b, 0xbc, 0x72, 0xc4, 0xeb, 0x63, 0x60, 0xef, 0x59, 0x9b, 0x01,
	0xc5, 0xbb, 0x17, 0xcc, 0xc0, 0x8e, 0x3b, 0xf9, 0x8a, 0xf0, 0xad, 0xe9, 0x6a, 0x3f, 0x33, 0x2b,
	0xf4, 0x1b, 0x7c, 0x63, 0x66, 0x14, 0x67, 0x1f, 0x8e, 0x34, 0xfe, 0x2e, 0x39, 0xe7, 0xef, 0x3e,
	0x95, 0x6f, 0x48, 0x2e, 0x8a, 0xdb, 0x56, 0xc6, 0x97, 0x22, 0xf4, 0x14, 0x25, 0xc9, 0x8f, 0xcf,
	0xbf, 0xff, 0xac, 0x0d, 0x36, 0x86, 0xf8, 0x91, 0x00, 0xeb, 0x50, 0x29, 0xf8, 0xb4, 0x3c, 0xcf,
	0x2c, 0xf1, 0xa6, 0x4a, 0xef, 0xb7, 0x41, 0xef, 0xa3, 0x43, 0x84, 0x16, 0x6b, 0x5d, 0xe8, 0xf1,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x30, 0x98, 0xb2, 0xb5, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v3.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

// UnimplementedLoadReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadReportingServiceServer struct {
}

func (*UnimplementedLoadReportingServiceServer) StreamLoadStats(srv LoadReportingService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v3.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v3/lrs.proto",
}
