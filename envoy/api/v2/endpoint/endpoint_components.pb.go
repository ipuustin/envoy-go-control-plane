// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint/endpoint_components.proto

package envoy_api_v2_endpoint

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
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

type Endpoint struct {
	Address              *core.Address               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	HealthCheckConfig    *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	Hostname             string                      `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d96d13bf4e60dd1, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

func (m *Endpoint) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type Endpoint_HealthCheckConfig struct {
	PortValue            uint32   `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d96d13bf4e60dd1, []int{0, 0}
}

func (m *Endpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Endpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (m *Endpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_HealthCheckConfig.Merge(m, src)
}
func (m *Endpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Size(m)
}
func (m *Endpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

func (m *Endpoint_HealthCheckConfig) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type LbEndpoint struct {
	// Types that are valid to be assigned to HostIdentifier:
	//	*LbEndpoint_Endpoint
	//	*LbEndpoint_EndpointName
	HostIdentifier       isLbEndpoint_HostIdentifier `protobuf_oneof:"host_identifier"`
	HealthStatus         core.HealthStatus           `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	Metadata             *core.Metadata              `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value       `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d96d13bf4e60dd1, []int{1}
}

func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (m *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(m, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

type isLbEndpoint_HostIdentifier interface {
	isLbEndpoint_HostIdentifier()
}

type LbEndpoint_Endpoint struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3,oneof"`
}

type LbEndpoint_EndpointName struct {
	EndpointName string `protobuf:"bytes,5,opt,name=endpoint_name,json=endpointName,proto3,oneof"`
}

func (*LbEndpoint_Endpoint) isLbEndpoint_HostIdentifier() {}

func (*LbEndpoint_EndpointName) isLbEndpoint_HostIdentifier() {}

func (m *LbEndpoint) GetHostIdentifier() isLbEndpoint_HostIdentifier {
	if m != nil {
		return m.HostIdentifier
	}
	return nil
}

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_Endpoint); ok {
		return x.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetEndpointName() string {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_EndpointName); ok {
		return x.EndpointName
	}
	return ""
}

func (m *LbEndpoint) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LbEndpoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LbEndpoint_Endpoint)(nil),
		(*LbEndpoint_EndpointName)(nil),
	}
}

type LocalityLbEndpoints struct {
	Locality             *core.Locality        `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	LbEndpoints          []*LbEndpoint         `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	Priority             uint32                `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Proximity            *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=proximity,proto3" json:"proximity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d96d13bf4e60dd1, []int{2}
}

func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(m, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LocalityLbEndpoints) GetProximity() *wrappers.UInt32Value {
	if m != nil {
		return m.Proximity
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v2.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v2.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v2.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v2.endpoint.LocalityLbEndpoints")
}

func init() {
	proto.RegisterFile("envoy/api/v2/endpoint/endpoint_components.proto", fileDescriptor_2d96d13bf4e60dd1)
}

var fileDescriptor_2d96d13bf4e60dd1 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x4e, 0x13, 0x41,
	0x18, 0xa6, 0x5b, 0x5a, 0xda, 0x29, 0xa8, 0x0c, 0x21, 0x6c, 0x2a, 0x02, 0x22, 0x26, 0x0d, 0x87,
	0xdd, 0x58, 0x4c, 0x4c, 0x4c, 0x3c, 0xb8, 0x60, 0x82, 0x09, 0x1a, 0x32, 0x46, 0x0d, 0xa7, 0xcd,
	0xdf, 0xdd, 0xa1, 0x3b, 0x71, 0x3b, 0xb3, 0xd9, 0x4e, 0x8b, 0xdc, 0x78, 0x03, 0xaf, 0xc6, 0x47,
	0xf1, 0x09, 0xbc, 0xf2, 0x04, 0xbe, 0x83, 0x47, 0x0e, 0x62, 0x66, 0x76, 0x67, 0xdb, 0xd2, 0x12,
	0x0e, 0xde, 0x76, 0xe6, 0xff, 0xfe, 0xef, 0xfb, 0xbf, 0x6f, 0xfe, 0x16, 0xb9, 0x94, 0x0f, 0xc5,
	0xb9, 0x0b, 0x09, 0x73, 0x87, 0x6d, 0x97, 0xf2, 0x30, 0x11, 0x8c, 0xcb, 0xe2, 0xc3, 0x0f, 0x44,
	0x2f, 0x11, 0x9c, 0x72, 0xd9, 0x77, 0x92, 0x54, 0x48, 0x81, 0x57, 0x75, 0x83, 0x03, 0x09, 0x73,
	0x86, 0x6d, 0xc7, 0xe0, 0x9a, 0x9b, 0x13, 0x3c, 0x81, 0x48, 0xa9, 0x0b, 0x61, 0x98, 0xd2, 0x7e,
	0xde, 0xd7, 0x5c, 0x9f, 0x06, 0x74, 0xa0, 0x4f, 0xf3, 0xea, 0xce, 0x74, 0x35, 0xa2, 0x10, 0xcb,
	0xc8, 0x0f, 0x22, 0x1a, 0x7c, 0xc9, 0x51, 0x1b, 0x5d, 0x21, 0xba, 0x31, 0x75, 0xf5, 0xa9, 0x33,
	0x38, 0x75, 0xcf, 0x52, 0x48, 0x12, 0x9a, 0x1a, 0x8d, 0x8d, 0x41, 0x98, 0x80, 0x0b, 0x9c, 0x0b,
	0x09, 0x92, 0x09, 0xde, 0x77, 0x7b, 0xac, 0x9b, 0x82, 0x34, 0x2a, 0x8f, 0xa6, 0xea, 0x7d, 0x09,
	0x72, 0x60, 0xda, 0xd7, 0x86, 0x10, 0xb3, 0x10, 0x24, 0x75, 0xcd, 0x47, 0x56, 0xd8, 0xfe, 0x61,
	0xa1, 0xda, 0x9b, 0xdc, 0x29, 0x7e, 0x8e, 0x16, 0x72, 0x67, 0x76, 0x69, 0xab, 0xd4, 0x6a, 0xb4,
	0x9b, 0xce, 0x44, 0x24, 0x6a, 0x78, 0xe7, 0x75, 0x86, 0x20, 0x06, 0x8a, 0x01, 0xad, 0x8c, 0x1b,
	0xf2, 0x03, 0xc1, 0x4f, 0x59, 0xd7, 0xb6, 0x34, 0xc3, 0x33, 0x67, 0x66, 0xa8, 0x8e, 0xd1, 0x74,
	0x0e, 0x75, 0xeb, 0xbe, 0xea, 0xdc, 0xd7, 0x8d, 0x64, 0x39, 0xba, 0x79, 0x85, 0x9b, 0xa8, 0x16,
	0x89, 0xbe, 0xe4, 0xd0, 0xa3, 0x76, 0x79, 0xab, 0xd4, 0xaa, 0x93, 0xe2, 0xdc, 0x3c, 0x41, 0xcb,
	0x53, 0x1c, 0xb8, 0x85, 0x50, 0x22, 0x52, 0xe9, 0x0f, 0x21, 0x1e, 0x50, 0x6d, 0x66, 0xc9, 0xab,
	0x5f, 0x79, 0xd5, 0xdd, 0x79, 0xfb, 0xfa, 0xba, 0x4c, 0xea, 0xaa, 0xf8, 0x49, 0xd5, 0x26, 0xa8,
	0xad, 0x49, 0xea, 0xed, 0xdf, 0x16, 0x42, 0x47, 0x9d, 0x22, 0x9e, 0x57, 0xa8, 0x66, 0xe6, 0xcf,
	0xf3, 0xd9, 0xbc, 0xc3, 0xdd, 0xe1, 0x1c, 0x29, 0x5a, 0xf0, 0x53, 0xb4, 0x54, 0xec, 0x9e, 0x96,
	0xab, 0x28, 0xb9, 0xc3, 0x39, 0xb2, 0x68, 0xae, 0xdf, 0x43, 0x8f, 0xe2, 0x03, 0xb4, 0x94, 0xc7,
	0x99, 0xbd, 0xa0, 0x9e, 0xea, 0xde, 0x4d, 0x29, 0xfd, 0x14, 0x99, 0xef, 0x0f, 0x1a, 0x46, 0x16,
	0xa3, 0xb1, 0x13, 0x7e, 0x81, 0x6a, 0x3d, 0x2a, 0x21, 0x04, 0x09, 0x3a, 0xb1, 0x46, 0xfb, 0xe1,
	0x0c, 0x82, 0x77, 0x39, 0x84, 0x14, 0x60, 0x7c, 0x82, 0x56, 0x63, 0x01, 0xa1, 0xdf, 0x81, 0x18,
	0x78, 0xc0, 0x78, 0xd7, 0x3f, 0xa3, 0xac, 0x1b, 0x49, 0x7b, 0x5e, 0xb3, 0xac, 0x3b, 0xd9, 0xa2,
	0x3a, 0x66, 0x51, 0x9d, 0x8f, 0x6f, 0xb9, 0xdc, 0x6b, 0xeb, 0x30, 0xbd, 0x85, 0x2b, 0x6f, 0x7e,
	0xd7, 0x6a, 0x95, 0xc8, 0x8a, 0xe2, 0xf0, 0x0c, 0xc5, 0x67, 0xcd, 0xe0, 0x2d, 0xa3, 0xfb, 0x2a,
	0x5a, 0x9f, 0x85, 0x94, 0x4b, 0x76, 0xca, 0x68, 0xba, 0x7d, 0x69, 0xa1, 0x95, 0x23, 0x11, 0x40,
	0xcc, 0xe4, 0xf9, 0x28, 0x69, 0x3d, 0x7e, 0x9c, 0x5f, 0xe7, 0x51, 0xcf, 0x1a, 0xdf, 0x74, 0x92,
	0x02, 0x8c, 0x0f, 0xd0, 0x62, 0xdc, 0xf1, 0x4d, 0xa0, 0x2a, 0xbc, 0x72, 0xab, 0xd1, 0x7e, 0x7c,
	0xcb, 0x3b, 0x8d, 0x24, 0x49, 0x23, 0x1e, 0x93, 0xbf, 0x35, 0x84, 0xf2, 0xff, 0x86, 0x80, 0x77,
	0x50, 0x2d, 0x49, 0x99, 0x48, 0x95, 0xb3, 0x8a, 0xde, 0xcb, 0xda, 0x95, 0x57, 0xd9, 0x2d, 0xdb,
	0x17, 0x25, 0x52, 0x54, 0xf0, 0x4b, 0x54, 0x4f, 0x52, 0xf1, 0x95, 0xf5, 0x14, 0xac, 0x7a, 0xb7,
	0x28, 0x19, 0xc1, 0xbd, 0xe8, 0xcf, 0xf7, 0xbf, 0xdf, 0x2a, 0x4d, 0x6c, 0x67, 0x9e, 0xb3, 0x9f,
	0xe3, 0xc8, 0xf3, 0x70, 0xef, 0xe7, 0xc5, 0xaf, 0xcb, 0xaa, 0xf5, 0xc0, 0x42, 0x4f, 0x98, 0xc8,
	0x82, 0x51, 0xcd, 0xe7, 0xb3, 0x33, 0xf2, 0xd6, 0x4c, 0x2c, 0xfb, 0xc5, 0xff, 0xe5, 0xb1, 0x9a,
	0xe0, 0xb8, 0xd4, 0xa9, 0xea, 0x51, 0xf6, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xda, 0x18, 0xae,
	0x95, 0x6a, 0x05, 0x00, 0x00,
}
