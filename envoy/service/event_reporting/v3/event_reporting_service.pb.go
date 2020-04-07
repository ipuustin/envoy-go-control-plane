// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/event_reporting/v3/event_reporting_service.proto

package envoy_service_event_reporting_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type StreamEventsRequest struct {
	Identifier           *StreamEventsRequest_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Events               []*any.Any                      `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *StreamEventsRequest) Reset()         { *m = StreamEventsRequest{} }
func (m *StreamEventsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamEventsRequest) ProtoMessage()    {}
func (*StreamEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_006b9de436985c31, []int{0}
}

func (m *StreamEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventsRequest.Unmarshal(m, b)
}
func (m *StreamEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventsRequest.Marshal(b, m, deterministic)
}
func (m *StreamEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventsRequest.Merge(m, src)
}
func (m *StreamEventsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamEventsRequest.Size(m)
}
func (m *StreamEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventsRequest proto.InternalMessageInfo

func (m *StreamEventsRequest) GetIdentifier() *StreamEventsRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamEventsRequest) GetEvents() []*any.Any {
	if m != nil {
		return m.Events
	}
	return nil
}

type StreamEventsRequest_Identifier struct {
	Node                 *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEventsRequest_Identifier) Reset()         { *m = StreamEventsRequest_Identifier{} }
func (m *StreamEventsRequest_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamEventsRequest_Identifier) ProtoMessage()    {}
func (*StreamEventsRequest_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_006b9de436985c31, []int{0, 0}
}

func (m *StreamEventsRequest_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventsRequest_Identifier.Unmarshal(m, b)
}
func (m *StreamEventsRequest_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventsRequest_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamEventsRequest_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventsRequest_Identifier.Merge(m, src)
}
func (m *StreamEventsRequest_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamEventsRequest_Identifier.Size(m)
}
func (m *StreamEventsRequest_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventsRequest_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventsRequest_Identifier proto.InternalMessageInfo

func (m *StreamEventsRequest_Identifier) GetNode() *v3.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type StreamEventsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEventsResponse) Reset()         { *m = StreamEventsResponse{} }
func (m *StreamEventsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamEventsResponse) ProtoMessage()    {}
func (*StreamEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_006b9de436985c31, []int{1}
}

func (m *StreamEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventsResponse.Unmarshal(m, b)
}
func (m *StreamEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventsResponse.Marshal(b, m, deterministic)
}
func (m *StreamEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventsResponse.Merge(m, src)
}
func (m *StreamEventsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamEventsResponse.Size(m)
}
func (m *StreamEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StreamEventsRequest)(nil), "envoy.service.event_reporting.v3.StreamEventsRequest")
	proto.RegisterType((*StreamEventsRequest_Identifier)(nil), "envoy.service.event_reporting.v3.StreamEventsRequest.Identifier")
	proto.RegisterType((*StreamEventsResponse)(nil), "envoy.service.event_reporting.v3.StreamEventsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/event_reporting/v3/event_reporting_service.proto", fileDescriptor_006b9de436985c31)
}

var fileDescriptor_006b9de436985c31 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6b, 0xd4, 0x40,
	0x14, 0xc6, 0x9d, 0x5d, 0x5d, 0xca, 0xd4, 0x43, 0x89, 0x15, 0x6b, 0x40, 0x5d, 0xf7, 0xb4, 0xa7,
	0x19, 0xd9, 0xc5, 0xa2, 0x51, 0xa4, 0x0d, 0x2a, 0x88, 0x20, 0x25, 0x3d, 0xf5, 0x54, 0x67, 0x37,
	0x6f, 0xe3, 0xc0, 0x3a, 0x2f, 0xce, 0x4c, 0x82, 0xb9, 0x09, 0x1e, 0x2c, 0x1e, 0xf5, 0x26, 0xf8,
	0x8f, 0x78, 0x17, 0xbc, 0xfa, 0xef, 0x78, 0x92, 0xc9, 0x24, 0x5a, 0xeb, 0x42, 0xe8, 0xde, 0x92,
	0xf7, 0xbe, 0xf7, 0x9b, 0xef, 0x7d, 0x33, 0xf4, 0x11, 0xa8, 0x12, 0x2b, 0x6e, 0x40, 0x97, 0x72,
	0x0e, 0x1c, 0x4a, 0x50, 0xf6, 0x58, 0x43, 0x8e, 0xda, 0x4a, 0x95, 0xf1, 0x72, 0x7a, 0xb6, 0x74,
	0xdc, 0x48, 0x59, 0xae, 0xd1, 0x62, 0x30, 0xac, 0xe7, 0x59, 0x5b, 0x3c, 0x23, 0x66, 0xe5, 0x34,
	0xbc, 0xe5, 0x4f, 0x98, 0xa3, 0x5a, 0xc8, 0x8c, 0xcf, 0x51, 0x83, 0xa3, 0xce, 0x84, 0x69, 0x10,
	0xe1, 0xf5, 0x0c, 0x31, 0x5b, 0x02, 0xaf, 0xff, 0x66, 0xc5, 0x82, 0x0b, 0x55, 0x35, 0xad, 0x1b,
	0x45, 0x9a, 0x0b, 0x2e, 0x94, 0x42, 0x2b, 0xac, 0x44, 0x65, 0xb8, 0xb1, 0xc2, 0x16, 0xa6, 0x69,
	0xdf, 0xfe, 0xaf, 0x5d, 0x82, 0x36, 0x12, 0x95, 0x3b, 0xda, 0x4b, 0xae, 0x95, 0x62, 0x29, 0x53,
	0x61, 0x81, 0xb7, 0x1f, 0xbe, 0x31, 0xfa, 0xd0, 0xa7, 0x57, 0x0e, 0xad, 0x06, 0xf1, 0xfa, 0x89,
	0xf3, 0x6c, 0x12, 0x78, 0x53, 0x80, 0xb1, 0xc1, 0x4b, 0x4a, 0x65, 0x0a, 0xca, 0xca, 0x85, 0x04,
	0xbd, 0x43, 0x86, 0x64, 0xbc, 0x39, 0xd9, 0x63, 0x5d, 0x5b, 0xb2, 0x15, 0x28, 0xf6, 0xec, 0x0f,
	0x27, 0x39, 0xc5, 0x0c, 0x76, 0xe9, 0xa0, 0x06, 0x98, 0x9d, 0xde, 0xb0, 0x3f, 0xde, 0x9c, 0x6c,
	0x33, 0x1f, 0x00, 0x6b, 0x03, 0x60, 0xfb, 0xaa, 0x8a, 0x37, 0x7e, 0xc5, 0x97, 0x3e, 0x91, 0xde,
	0x06, 0x49, 0x1a, 0x75, 0xf8, 0x99, 0x50, 0xfa, 0x17, 0x19, 0xdc, 0xa3, 0x17, 0x15, 0xa6, 0xd0,
	0x58, 0x0c, 0x1b, 0x8b, 0x3e, 0x66, 0xe6, 0x62, 0x76, 0xb6, 0x5e, 0x60, 0x0a, 0x35, 0xea, 0x23,
	0xe9, 0x6d, 0x91, 0xa4, 0x9e, 0x88, 0x9e, 0x7f, 0xf9, 0x7e, 0x72, 0xf3, 0x29, 0x7d, 0xdc, 0xb1,
	0xd4, 0x44, 0x2c, 0xf3, 0x57, 0xa2, 0x63, 0xb3, 0x68, 0xcf, 0xc1, 0x1e, 0xd0, 0xfb, 0x6b, 0xc3,
	0x46, 0x47, 0x74, 0xfb, 0xdf, 0xb2, 0xc9, 0x51, 0x19, 0x88, 0xf6, 0x1d, 0xf9, 0x21, 0x8d, 0xd6,
	0x21, 0x7b, 0xc4, 0xe4, 0x2b, 0xa1, 0x57, 0xeb, 0x52, 0xd2, 0xea, 0x0f, 0x3d, 0x26, 0x78, 0x4f,
	0xe8, 0xe5, 0xd3, 0x23, 0xc1, 0xdd, 0xb5, 0xee, 0x38, 0xdc, 0x3d, 0xef, 0x98, 0x77, 0x36, 0xba,
	0x30, 0x26, 0x77, 0x48, 0x7c, 0xf4, 0xed, 0xdd, 0x8f, 0x9f, 0x83, 0xde, 0x56, 0x9f, 0x32, 0x89,
	0x9e, 0x94, 0x6b, 0x7c, 0x5b, 0x75, 0x42, 0xe3, 0x70, 0xe5, 0x5a, 0x07, 0xee, 0x05, 0x1d, 0x90,
	0x13, 0x42, 0x66, 0x83, 0xfa, 0x35, 0x4d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xe1, 0x78,
	0x3f, 0xe1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventReportingServiceClient is the client API for EventReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventReportingServiceClient interface {
	StreamEvents(ctx context.Context, opts ...grpc.CallOption) (EventReportingService_StreamEventsClient, error)
}

type eventReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventReportingServiceClient(cc *grpc.ClientConn) EventReportingServiceClient {
	return &eventReportingServiceClient{cc}
}

func (c *eventReportingServiceClient) StreamEvents(ctx context.Context, opts ...grpc.CallOption) (EventReportingService_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventReportingService_serviceDesc.Streams[0], "/envoy.service.event_reporting.v3.EventReportingService/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventReportingServiceStreamEventsClient{stream}
	return x, nil
}

type EventReportingService_StreamEventsClient interface {
	Send(*StreamEventsRequest) error
	Recv() (*StreamEventsResponse, error)
	grpc.ClientStream
}

type eventReportingServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *eventReportingServiceStreamEventsClient) Send(m *StreamEventsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventReportingServiceStreamEventsClient) Recv() (*StreamEventsResponse, error) {
	m := new(StreamEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventReportingServiceServer is the server API for EventReportingService service.
type EventReportingServiceServer interface {
	StreamEvents(EventReportingService_StreamEventsServer) error
}

// UnimplementedEventReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventReportingServiceServer struct {
}

func (*UnimplementedEventReportingServiceServer) StreamEvents(srv EventReportingService_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}

func RegisterEventReportingServiceServer(s *grpc.Server, srv EventReportingServiceServer) {
	s.RegisterService(&_EventReportingService_serviceDesc, srv)
}

func _EventReportingService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventReportingServiceServer).StreamEvents(&eventReportingServiceStreamEventsServer{stream})
}

type EventReportingService_StreamEventsServer interface {
	Send(*StreamEventsResponse) error
	Recv() (*StreamEventsRequest, error)
	grpc.ServerStream
}

type eventReportingServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *eventReportingServiceStreamEventsServer) Send(m *StreamEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventReportingServiceStreamEventsServer) Recv() (*StreamEventsRequest, error) {
	m := new(StreamEventsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.event_reporting.v3.EventReportingService",
	HandlerType: (*EventReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _EventReportingService_StreamEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/event_reporting/v3/event_reporting_service.proto",
}
