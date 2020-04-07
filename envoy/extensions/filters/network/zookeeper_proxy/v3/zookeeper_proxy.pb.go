// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/zookeeper_proxy/v3/zookeeper_proxy.proto

package envoy_extensions_filters_network_zookeeper_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type ZooKeeperProxy struct {
	StatPrefix           string                `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	AccessLog            string                `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	MaxPacketBytes       *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_packet_bytes,json=maxPacketBytes,proto3" json:"max_packet_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ZooKeeperProxy) Reset()         { *m = ZooKeeperProxy{} }
func (m *ZooKeeperProxy) String() string { return proto.CompactTextString(m) }
func (*ZooKeeperProxy) ProtoMessage()    {}
func (*ZooKeeperProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c9f1bab8219a3e6, []int{0}
}

func (m *ZooKeeperProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZooKeeperProxy.Unmarshal(m, b)
}
func (m *ZooKeeperProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZooKeeperProxy.Marshal(b, m, deterministic)
}
func (m *ZooKeeperProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZooKeeperProxy.Merge(m, src)
}
func (m *ZooKeeperProxy) XXX_Size() int {
	return xxx_messageInfo_ZooKeeperProxy.Size(m)
}
func (m *ZooKeeperProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ZooKeeperProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ZooKeeperProxy proto.InternalMessageInfo

func (m *ZooKeeperProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ZooKeeperProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *ZooKeeperProxy) GetMaxPacketBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPacketBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*ZooKeeperProxy)(nil), "envoy.extensions.filters.network.zookeeper_proxy.v3.ZooKeeperProxy")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/zookeeper_proxy/v3/zookeeper_proxy.proto", fileDescriptor_8c9f1bab8219a3e6)
}

var fileDescriptor_8c9f1bab8219a3e6 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x49, 0x2b, 0x95, 0x4e, 0xa1, 0x94, 0xb8, 0xb0, 0x14, 0x5b, 0xaa, 0xab, 0xae, 0x66,
	0x68, 0xb3, 0x73, 0x67, 0x44, 0xa1, 0xea, 0x22, 0x14, 0x74, 0xd1, 0x4d, 0x98, 0xa4, 0x27, 0x31,
	0x34, 0x9d, 0x33, 0x4c, 0x26, 0x69, 0xea, 0xca, 0xa5, 0xcf, 0xe0, 0xa3, 0xb8, 0x17, 0xdc, 0xfa,
	0x20, 0xbe, 0x80, 0xab, 0xcb, 0x4c, 0x12, 0x2e, 0xb7, 0x97, 0xbb, 0xb9, 0xbb, 0xf0, 0xff, 0xe7,
	0xff, 0x72, 0xfe, 0x33, 0x64, 0x0b, 0xa2, 0xc2, 0x0b, 0x83, 0x5a, 0x83, 0x28, 0x32, 0x14, 0x05,
	0x4b, 0xb2, 0x5c, 0x83, 0x2a, 0x98, 0x00, 0x7d, 0x46, 0x75, 0x64, 0xdf, 0x10, 0x8f, 0x00, 0x12,
	0x54, 0x28, 0x15, 0xd6, 0x17, 0x56, 0x79, 0xd7, 0x12, 0x95, 0x0a, 0x35, 0xba, 0x9e, 0x45, 0xd1,
	0x5b, 0x14, 0x6d, 0x51, 0xb4, 0x45, 0xd1, 0xeb, 0x5c, 0xe5, 0xcd, 0x16, 0x29, 0x62, 0x9a, 0x03,
	0xb3, 0x88, 0xa8, 0x4c, 0xd8, 0x59, 0x71, 0x29, 0x4d, 0xc8, 0x2a, 0xb3, 0x79, 0x79, 0x90, 0x9c,
	0x71, 0x21, 0x50, 0x73, 0x6d, 0xf7, 0x2b, 0x34, 0xd7, 0x65, 0x67, 0xbf, 0xbc, 0x67, 0x57, 0xa0,
	0xcc, 0xcf, 0x33, 0x91, 0xb6, 0x23, 0xcf, 0x2b, 0x9e, 0x67, 0x07, 0xae, 0x81, 0x75, 0x1f, 0x8d,
	0xf1, 0xea, 0x9f, 0x43, 0xc6, 0x7b, 0xc4, 0x8f, 0x76, 0xa3, 0xc0, 0x2c, 0xe4, 0xae, 0xc8, 0xc8,
	0xe0, 0x43, 0xa9, 0x20, 0xc9, 0xea, 0xa9, 0xb3, 0x74, 0x56, 0x43, 0xff, 0xe9, 0x7f, 0xff, 0x89,
	0xea, 0x2d, 0x9d, 0x1d, 0x31, 0x5e, 0x60, 0x2d, 0x77, 0x4e, 0x08, 0x8f, 0x63, 0x28, 0x8a, 0x30,
	0xc7, 0x74, 0xda, 0x33, 0x83, 0xbb, 0x61, 0xa3, 0x7c, 0xc2, 0xd4, 0x7d, 0x4f, 0x26, 0x27, 0x5e,
	0x87, 0x92, 0xc7, 0x47, 0xd0, 0x61, 0x74, 0xd1, 0x50, 0x4c, 0xfb, 0x4b, 0x67, 0x35, 0xda, 0xbc,
	0xa0, 0x4d, 0x63, 0xda, 0x35, 0xa6, 0x9f, 0xb7, 0x42, 0x7b, 0x9b, 0x2f, 0x3c, 0x2f, 0x61, 0x37,
	0x3e, 0xf1, 0x3a, 0xb0, 0x21, 0xdf, 0x64, 0x5e, 0x7f, 0xf8, 0xf9, 0xfb, 0xc7, 0xe2, 0x1d, 0x79,
	0xdb, 0x9c, 0x36, 0x46, 0x91, 0x64, 0x69, 0x7b, 0xd6, 0x87, 0xaf, 0xba, 0xe6, 0xb9, 0xfc, 0xca,
	0xd7, 0xf4, 0x6e, 0x39, 0x3f, 0xfa, 0xf5, 0xfd, 0xcf, 0xdf, 0x41, 0x6f, 0xd2, 0x27, 0x6f, 0x32,
	0xa4, 0x96, 0xd8, 0x04, 0x1e, 0xf1, 0x6e, 0xfe, 0xb3, 0x7d, 0xa7, 0x59, 0x78, 0x60, 0xca, 0x04,
	0x4e, 0x34, 0xb0, 0xad, 0xbc, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x1d, 0x4a, 0xc4, 0x57,
	0x02, 0x00, 0x00,
}
