// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/access_loggers/grpc/v3/als.proto

package envoy_extensions_access_loggers_grpc_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HttpGrpcAccessLogConfig struct {
	CommonConfig                    *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	AdditionalRequestHeadersToLog   []string                   `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                   `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                   `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                   `json:"-"`
	XXX_unrecognized                []byte                     `json:"-"`
	XXX_sizecache                   int32                      `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3685c6e6be71fba, []int{0}
}

func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Size(m)
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3685c6e6be71fba, []int{1}
}

func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Size(m)
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

type CommonGrpcAccessLogConfig struct {
	LogName                 string                `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	GrpcService             *v3.GrpcService       `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	BufferFlushInterval     *duration.Duration    `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	BufferSizeBytes         *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	FilterStateObjectsToLog []string              `protobuf:"bytes,5,rep,name=filter_state_objects_to_log,json=filterStateObjectsToLog,proto3" json:"filter_state_objects_to_log,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3685c6e6be71fba, []int{2}
}

func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Size(m)
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *v3.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *duration.Duration {
	if m != nil {
		return m.BufferFlushInterval
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.BufferSizeBytes
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetFilterStateObjectsToLog() []string {
	if m != nil {
		return m.FilterStateObjectsToLog
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.extensions.access_loggers.grpc.v3.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.extensions.access_loggers.grpc.v3.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/access_loggers/grpc/v3/als.proto", fileDescriptor_d3685c6e6be71fba)
}

var fileDescriptor_d3685c6e6be71fba = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xc7, 0x7f, 0xdb, 0xf2, 0x77, 0xe0, 0x17, 0x71, 0x8d, 0x52, 0x50, 0x10, 0xea, 0x41, 0xe2,
	0x61, 0x56, 0xdb, 0x10, 0x0d, 0xf1, 0x42, 0x31, 0x5a, 0x94, 0x20, 0x29, 0xa8, 0xc7, 0xcd, 0x74,
	0xfb, 0x74, 0x18, 0x33, 0x9d, 0x59, 0x66, 0x66, 0x57, 0xe0, 0xe4, 0xd1, 0xf8, 0x12, 0x7c, 0x09,
	0xc4, 0x57, 0xe0, 0x59, 0x13, 0xaf, 0xbe, 0x1d, 0x4e, 0x66, 0x66, 0x16, 0x69, 0x85, 0x26, 0x1c,
	0xbd, 0xb5, 0xfb, 0x7c, 0xbf, 0x9f, 0x67, 0xbe, 0x33, 0xcf, 0x83, 0x1e, 0x81, 0xc8, 0xe5, 0x51,
	0x04, 0x87, 0x06, 0x84, 0x66, 0x52, 0xe8, 0x88, 0x24, 0x09, 0x68, 0x1d, 0x73, 0x49, 0x29, 0x28,
	0x1d, 0x51, 0x95, 0x26, 0x51, 0x5e, 0x8f, 0x08, 0xd7, 0x38, 0x55, 0xd2, 0xc8, 0xf0, 0xbe, 0xb3,
	0xe0, 0x73, 0x0b, 0x1e, 0xb4, 0x60, 0x6b, 0xc1, 0x79, 0x7d, 0xde, 0x0b, 0xa3, 0x44, 0x8a, 0x2e,
	0xa3, 0x51, 0x22, 0x15, 0x58, 0x90, 0xad, 0xc6, 0x1a, 0x54, 0xce, 0x12, 0xf0, 0xc4, 0xf9, 0x45,
	0x2a, 0x25, 0xe5, 0x10, 0xb9, 0x7f, 0xed, 0xac, 0x1b, 0x75, 0x32, 0x45, 0x0c, 0x93, 0x62, 0x58,
	0xfd, 0x83, 0x22, 0x69, 0x6a, 0x5b, 0xf9, 0xfa, 0x42, 0xd6, 0x49, 0x49, 0x44, 0x84, 0x90, 0xc6,
	0xd9, 0x74, 0xa4, 0x0d, 0x31, 0xd9, 0x59, 0x79, 0xf9, 0x42, 0x39, 0x07, 0x65, 0x4f, 0xce, 0x04,
	0x2d, 0x24, 0xb3, 0x39, 0xe1, 0xac, 0x43, 0x0c, 0x44, 0x67, 0x3f, 0x7c, 0xa1, 0x7a, 0x52, 0x46,
	0xb3, 0x4d, 0x63, 0xd2, 0x17, 0x2a, 0x4d, 0xd6, 0x5d, 0xcc, 0x2d, 0x49, 0x37, 0x5c, 0xa2, 0xf0,
	0x00, 0xfd, 0x9f, 0xc8, 0x5e, 0x4f, 0x8a, 0xd8, 0x47, 0xac, 0x04, 0x4b, 0xc1, 0xca, 0x54, 0xad,
	0x81, 0xaf, 0x78, 0x41, 0x78, 0xc3, 0xb9, 0x2f, 0x41, 0x37, 0x26, 0x4e, 0x1b, 0xa3, 0x9f, 0x83,
	0xd2, 0x4c, 0xd0, 0x9a, 0xf6, 0x2d, 0x8a, 0x96, 0x4d, 0xb4, 0x4c, 0x3a, 0x1d, 0x66, 0x53, 0x10,
	0x1e, 0x2b, 0x38, 0xc8, 0x40, 0x9b, 0x78, 0x1f, 0x48, 0x07, 0x94, 0x8e, 0x8d, 0xb4, 0x2d, 0x2a,
	0xa5, 0xa5, 0xf2, 0xca, 0x64, 0x6b, 0xe1, 0x5c, 0xd8, 0xf2, 0xba, 0xa6, 0x97, 0xed, 0xc9, 0x2d,
	0x49, 0xc3, 0x97, 0xa8, 0x3a, 0x40, 0xd2, 0xa9, 0x14, 0x1a, 0xfe, 0x46, 0x95, 0x1d, 0x6a, 0xb1,
	0x1f, 0xe5, 0x85, 0x03, 0xac, 0x2d, 0x74, 0xef, 0x32, 0x96, 0x51, 0x84, 0xf1, 0x3e, 0xd8, 0x88,
	0x83, 0xdd, 0xbd, 0x08, 0xdb, 0x2b, 0x84, 0x8e, 0xb6, 0xf6, 0xe4, 0xcb, 0x8f, 0x4f, 0x8b, 0xf5,
	0x62, 0x32, 0xb1, 0xbf, 0xda, 0xe2, 0x06, 0xb9, 0xa4, 0x38, 0xaf, 0xe1, 0x21, 0x0f, 0x52, 0xfd,
	0x1e, 0xa0, 0x5b, 0x7b, 0xc9, 0x3f, 0xf2, 0x56, 0x6b, 0x8f, 0x6d, 0x8e, 0x1a, 0x7a, 0x38, 0x3c,
	0xc7, 0xe5, 0x67, 0xad, 0x7e, 0x2d, 0xa3, 0xb9, 0xa1, 0xed, 0xc2, 0x2a, 0x9a, 0xe0, 0x92, 0xc6,
	0x82, 0xf4, 0xc0, 0x85, 0x98, 0x6c, 0x8c, 0x9f, 0x36, 0x46, 0x54, 0x69, 0x29, 0x68, 0x8d, 0x73,
	0x49, 0xb7, 0x49, 0x0f, 0xc2, 0x6d, 0x34, 0xdd, 0xbf, 0x66, 0x95, 0x92, 0x0b, 0xbb, 0x8c, 0x07,
	0x8e, 0x62, 0x17, 0xd2, 0x26, 0xb3, 0x4d, 0x76, 0xbd, 0xb0, 0x2f, 0xcb, 0x14, 0x3d, 0xff, 0x1c,
	0xbe, 0x43, 0x37, 0xdb, 0x59, 0xb7, 0x0b, 0x2a, 0xee, 0xf2, 0x4c, 0xef, 0xc7, 0x4c, 0x18, 0x50,
	0x39, 0xe1, 0x95, 0xb2, 0x03, 0xcf, 0x61, 0xbf, 0xa0, 0xf8, 0x6c, 0x41, 0xf1, 0xb3, 0x62, 0x81,
	0x1d, 0xf0, 0x24, 0x28, 0x3d, 0xf8, 0xaf, 0x75, 0xc3, 0x13, 0x9e, 0x5b, 0xc0, 0x66, 0xe1, 0x0f,
	0x9b, 0xe8, 0x7a, 0x01, 0xd6, 0xec, 0x18, 0xe2, 0xf6, 0x91, 0x01, 0x5d, 0x19, 0x71, 0xd0, 0x3b,
	0x17, 0xa0, 0x6f, 0x36, 0x85, 0xa9, 0xd7, 0xde, 0x12, 0x9e, 0x41, 0xeb, 0x9a, 0xb7, 0xed, 0xb2,
	0x63, 0x68, 0x58, 0x53, 0xf8, 0x14, 0xdd, 0xee, 0x32, 0x6e, 0x2c, 0xc9, 0x10, 0x03, 0xb1, 0x6c,
	0xbf, 0x87, 0xc4, 0xfc, 0x99, 0xbd, 0x51, 0x37, 0x7b, 0xb3, 0x5e, 0xb2, 0x6b, 0x15, 0xaf, 0xbd,
	0xc0, 0xcf, 0xdc, 0x9a, 0x7d, 0xab, 0x55, 0x54, 0x1f, 0xfe, 0x56, 0xc3, 0xdf, 0xff, 0xd5, 0xb7,
	0x8f, 0x3f, 0x7f, 0x8d, 0x95, 0x66, 0xca, 0x68, 0x95, 0x49, 0x7f, 0xc5, 0xa9, 0x92, 0x87, 0x47,
	0x57, 0x1d, 0xad, 0xc6, 0xc4, 0x3a, 0xd7, 0x3b, 0x36, 0xe4, 0x4e, 0xd0, 0x1e, 0x73, 0x69, 0xeb,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x17, 0x32, 0x67, 0x98, 0x05, 0x00, 0x00,
}
