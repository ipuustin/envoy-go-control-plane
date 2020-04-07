// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3/clusters.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	proto "github.com/golang/protobuf/proto"
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

type Clusters struct {
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{0}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

type ClusterStatus struct {
	Name                                    string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AddedViaApi                             bool          `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	SuccessRateEjectionThreshold            *v3.Percent   `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	HostStatuses                            []*HostStatus `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	LocalOriginSuccessRateEjectionThreshold *v3.Percent   `protobuf:"bytes,5,opt,name=local_origin_success_rate_ejection_threshold,json=localOriginSuccessRateEjectionThreshold,proto3" json:"local_origin_success_rate_ejection_threshold,omitempty"`
	XXX_NoUnkeyedLiteral                    struct{}      `json:"-"`
	XXX_unrecognized                        []byte        `json:"-"`
	XXX_sizecache                           int32         `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{1}
}

func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStatus.Unmarshal(m, b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterStatus.Size(m)
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetAddedViaApi() bool {
	if m != nil {
		return m.AddedViaApi
	}
	return false
}

func (m *ClusterStatus) GetSuccessRateEjectionThreshold() *v3.Percent {
	if m != nil {
		return m.SuccessRateEjectionThreshold
	}
	return nil
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

func (m *ClusterStatus) GetLocalOriginSuccessRateEjectionThreshold() *v3.Percent {
	if m != nil {
		return m.LocalOriginSuccessRateEjectionThreshold
	}
	return nil
}

type HostStatus struct {
	Address                *v31.Address      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stats                  []*SimpleMetric   `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	HealthStatus           *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	SuccessRate            *v3.Percent       `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	Weight                 uint32            `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Hostname               string            `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Priority               uint32            `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	LocalOriginSuccessRate *v3.Percent       `protobuf:"bytes,8,opt,name=local_origin_success_rate,json=localOriginSuccessRate,proto3" json:"local_origin_success_rate,omitempty"`
	Locality               *v31.Locality     `protobuf:"bytes,9,opt,name=locality,proto3" json:"locality,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{2}
}

func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatus.Unmarshal(m, b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return xxx_messageInfo_HostStatus.Size(m)
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() *v31.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *HostStatus) GetStats() []*SimpleMetric {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *HostStatus) GetHealthStatus() *HostHealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return nil
}

func (m *HostStatus) GetSuccessRate() *v3.Percent {
	if m != nil {
		return m.SuccessRate
	}
	return nil
}

func (m *HostStatus) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *HostStatus) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostStatus) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *HostStatus) GetLocalOriginSuccessRate() *v3.Percent {
	if m != nil {
		return m.LocalOriginSuccessRate
	}
	return nil
}

func (m *HostStatus) GetLocality() *v31.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

type HostHealthStatus struct {
	FailedActiveHealthCheck   bool             `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	FailedOutlierCheck        bool             `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	FailedActiveDegradedCheck bool             `protobuf:"varint,4,opt,name=failed_active_degraded_check,json=failedActiveDegradedCheck,proto3" json:"failed_active_degraded_check,omitempty"`
	PendingDynamicRemoval     bool             `protobuf:"varint,5,opt,name=pending_dynamic_removal,json=pendingDynamicRemoval,proto3" json:"pending_dynamic_removal,omitempty"`
	PendingActiveHc           bool             `protobuf:"varint,6,opt,name=pending_active_hc,json=pendingActiveHc,proto3" json:"pending_active_hc,omitempty"`
	EdsHealthStatus           v31.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.config.core.v3.HealthStatus" json:"eds_health_status,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}         `json:"-"`
	XXX_unrecognized          []byte           `json:"-"`
	XXX_sizecache             int32            `json:"-"`
}

func (m *HostHealthStatus) Reset()         { *m = HostHealthStatus{} }
func (m *HostHealthStatus) String() string { return proto.CompactTextString(m) }
func (*HostHealthStatus) ProtoMessage()    {}
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_703c0abdc0a96402, []int{3}
}

func (m *HostHealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostHealthStatus.Unmarshal(m, b)
}
func (m *HostHealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostHealthStatus.Marshal(b, m, deterministic)
}
func (m *HostHealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostHealthStatus.Merge(m, src)
}
func (m *HostHealthStatus) XXX_Size() int {
	return xxx_messageInfo_HostHealthStatus.Size(m)
}
func (m *HostHealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostHealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostHealthStatus proto.InternalMessageInfo

func (m *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if m != nil {
		return m.FailedActiveHealthCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedOutlierCheck() bool {
	if m != nil {
		return m.FailedOutlierCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedActiveDegradedCheck() bool {
	if m != nil {
		return m.FailedActiveDegradedCheck
	}
	return false
}

func (m *HostHealthStatus) GetPendingDynamicRemoval() bool {
	if m != nil {
		return m.PendingDynamicRemoval
	}
	return false
}

func (m *HostHealthStatus) GetPendingActiveHc() bool {
	if m != nil {
		return m.PendingActiveHc
	}
	return false
}

func (m *HostHealthStatus) GetEdsHealthStatus() v31.HealthStatus {
	if m != nil {
		return m.EdsHealthStatus
	}
	return v31.HealthStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*Clusters)(nil), "envoy.admin.v3.Clusters")
	proto.RegisterType((*ClusterStatus)(nil), "envoy.admin.v3.ClusterStatus")
	proto.RegisterType((*HostStatus)(nil), "envoy.admin.v3.HostStatus")
	proto.RegisterType((*HostHealthStatus)(nil), "envoy.admin.v3.HostHealthStatus")
}

func init() { proto.RegisterFile("envoy/admin/v3/clusters.proto", fileDescriptor_703c0abdc0a96402) }

var fileDescriptor_703c0abdc0a96402 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6f, 0x23, 0x35,
	0x18, 0x55, 0x9a, 0x6c, 0x76, 0xd6, 0xd9, 0x6c, 0xbb, 0x16, 0x74, 0x67, 0x43, 0x9a, 0x4d, 0x53,
	0xa0, 0x11, 0xa0, 0x04, 0xa5, 0x12, 0x15, 0xe1, 0x50, 0xf5, 0x97, 0xd4, 0x03, 0xd0, 0x32, 0x45,
	0xdc, 0xd0, 0xc8, 0xf5, 0x7c, 0xcd, 0x18, 0x26, 0xe3, 0x91, 0xed, 0x04, 0x72, 0x83, 0x5b, 0xff,
	0x06, 0xae, 0xfc, 0x17, 0xdc, 0x91, 0xb8, 0xf2, 0x1f, 0x21, 0xff, 0x98, 0x64, 0x92, 0xa6, 0xe1,
	0x36, 0x9e, 0xf7, 0xde, 0xe7, 0xcf, 0xef, 0x7d, 0x36, 0xda, 0x83, 0x74, 0xca, 0x67, 0x7d, 0x12,
	0x8d, 0x59, 0xda, 0x9f, 0x1e, 0xf5, 0x69, 0x32, 0x91, 0x0a, 0x84, 0xec, 0x65, 0x82, 0x2b, 0x8e,
	0x5f, 0x19, 0xb8, 0x67, 0xe0, 0xde, 0xf4, 0xa8, 0xd1, 0x5c, 0xa1, 0x8f, 0x41, 0x09, 0x46, 0x1d,
	0xbb, 0xd1, 0xb1, 0x28, 0xe5, 0xe9, 0x3d, 0x1b, 0xf5, 0x29, 0x17, 0xa0, 0x39, 0x24, 0x8a, 0x04,
	0xc8, 0x9c, 0xf3, 0x6e, 0x2d, 0xe7, 0x8e, 0x48, 0x70, 0x84, 0xc3, 0xb5, 0x84, 0x18, 0x48, 0xa2,
	0xe2, 0x90, 0xc6, 0x40, 0x7f, 0x76, 0xc4, 0x0f, 0x2c, 0x51, 0xcd, 0x32, 0xc3, 0xc8, 0x40, 0x50,
	0x48, 0x95, 0x03, 0xf7, 0x26, 0x51, 0x46, 0xfa, 0x24, 0x4d, 0xb9, 0x22, 0x8a, 0xf1, 0x54, 0xf6,
	0xa5, 0x22, 0x6a, 0x92, 0x77, 0xb1, 0xff, 0x08, 0x9e, 0x82, 0x90, 0x8c, 0xa7, 0x2c, 0x1d, 0x59,
	0x4a, 0x67, 0x86, 0xbc, 0x73, 0x67, 0x06, 0xbe, 0x42, 0x3b, 0xce, 0x98, 0xd0, 0x96, 0x01, 0xe9,
	0x97, 0xda, 0xe5, 0x6e, 0x6d, 0xb0, 0xd7, 0x5b, 0x76, 0xa8, 0xe7, 0x34, 0xb7, 0x86, 0x16, 0x6c,
	0xd3, 0xe2, 0x12, 0xe4, 0xf0, 0xe0, 0x8f, 0xbf, 0x1f, 0x5a, 0x2d, 0xd4, 0x5c, 0x52, 0x0d, 0x48,
	0x92, 0xc5, 0x24, 0x97, 0xca, 0xce, 0x43, 0x19, 0xd5, 0x97, 0xea, 0x60, 0x8c, 0x2a, 0x29, 0x19,
	0x83, 0x5f, 0x6a, 0x97, 0xba, 0x2f, 0x02, 0xf3, 0x8d, 0x3b, 0xa8, 0x4e, 0xa2, 0x08, 0xa2, 0x70,
	0xca, 0x48, 0x48, 0x32, 0xe6, 0x6f, 0xb5, 0x4b, 0x5d, 0x2f, 0xa8, 0x99, 0x9f, 0x3f, 0x30, 0x72,
	0x9a, 0x31, 0xfc, 0x23, 0x7a, 0x27, 0x27, 0x94, 0x82, 0x94, 0xa1, 0x20, 0x0a, 0x42, 0xf8, 0x09,
	0xa8, 0x3e, 0x6f, 0xa8, 0x62, 0x01, 0x32, 0xe6, 0x49, 0xe4, 0x97, 0xdb, 0xa5, 0x6e, 0x6d, 0xb0,
	0xeb, 0xce, 0xa1, 0xdd, 0xd4, 0xc7, 0xb8, 0xb1, 0x6e, 0x06, 0x4d, 0x27, 0x0f, 0x88, 0x82, 0x4b,
	0x27, 0xfe, 0x3e, 0xd7, 0xe2, 0x13, 0x54, 0x8f, 0xb9, 0x54, 0x0b, 0x53, 0x2a, 0xc6, 0x94, 0xc6,
	0xaa, 0x29, 0x57, 0x5c, 0x2a, 0xe7, 0xc8, 0xcb, 0x78, 0xfe, 0x0d, 0x12, 0x4f, 0xd0, 0x67, 0x09,
	0xa7, 0x24, 0x09, 0xb9, 0x60, 0x23, 0x96, 0x86, 0xff, 0xd7, 0xec, 0xb3, 0x8d, 0xcd, 0x1e, 0x9a,
	0x5a, 0xd7, 0xa6, 0xd4, 0xed, 0x86, 0xbe, 0x87, 0x5d, 0x9d, 0xc2, 0x01, 0xda, 0xdf, 0x90, 0x82,
	0x6d, 0xb1, 0xf3, 0x7b, 0x05, 0xa1, 0x45, 0xf7, 0xf8, 0x18, 0x3d, 0x77, 0xe3, 0x6c, 0xa2, 0x58,
	0xe4, 0x6f, 0xc7, 0xb5, 0xa7, 0xc7, 0x55, 0x77, 0x74, 0x6a, 0x49, 0x41, 0xce, 0xc6, 0x03, 0xf4,
	0x4c, 0x9b, 0x24, 0xfd, 0x2d, 0xe3, 0x50, 0x73, 0xd5, 0xa1, 0x5b, 0x36, 0xce, 0x12, 0xf8, 0xc6,
	0x5c, 0xa7, 0xc0, 0x52, 0xf1, 0x25, 0xaa, 0xbb, 0xb1, 0xb7, 0xfe, 0xba, 0xa8, 0xda, 0xeb, 0xdc,
	0xbd, 0x32, 0xc4, 0xb9, 0xc7, 0x85, 0x15, 0xfe, 0x12, 0xbd, 0x2c, 0xda, 0xea, 0x57, 0x36, 0x7a,
	0x58, 0x2b, 0x04, 0x8e, 0x77, 0x51, 0xf5, 0x17, 0x60, 0xa3, 0x58, 0x19, 0xe3, 0xeb, 0x81, 0x5b,
	0xe1, 0x06, 0xf2, 0x74, 0x8c, 0x66, 0x24, 0xab, 0x66, 0x24, 0xe7, 0x6b, 0x8d, 0x65, 0x82, 0x71,
	0xc1, 0xd4, 0xcc, 0x7f, 0x6e, 0x54, 0xf3, 0x35, 0xfe, 0x0e, 0xbd, 0x7d, 0x32, 0x6e, 0xdf, 0xdb,
	0xd8, 0xd7, 0xee, 0xfa, 0x6c, 0xf1, 0x10, 0x79, 0x06, 0xd1, 0xdb, 0xbd, 0x30, 0x15, 0x5a, 0xeb,
	0x23, 0xf9, 0xda, 0xb1, 0x82, 0x39, 0x7f, 0xf8, 0x91, 0x1e, 0x83, 0x36, 0x6a, 0xad, 0x1b, 0x83,
	0x45, 0xe8, 0x9d, 0x3f, 0xcb, 0x68, 0x67, 0xd5, 0x63, 0xfc, 0x15, 0x6a, 0xdc, 0x13, 0x96, 0x40,
	0x14, 0x12, 0xaa, 0xd8, 0x14, 0xc2, 0xe2, 0x0b, 0x65, 0x86, 0xc3, 0x0b, 0xde, 0x58, 0xc6, 0xa9,
	0x21, 0x58, 0xf5, 0xb9, 0x86, 0xf1, 0xe7, 0xe8, 0x3d, 0x27, 0xe6, 0x13, 0x95, 0x30, 0x10, 0x4e,
	0x66, 0x6f, 0x30, 0xb6, 0xd8, 0xb5, 0x85, 0xac, 0xe2, 0x04, 0x35, 0x97, 0xb7, 0x8b, 0x60, 0x24,
	0x88, 0xbe, 0xfd, 0x56, 0x59, 0x31, 0xca, 0xb7, 0xc5, 0x0d, 0x2f, 0x1c, 0xc3, 0x16, 0xf8, 0x02,
	0xbd, 0xc9, 0x20, 0x8d, 0x58, 0x3a, 0x0a, 0xa3, 0x59, 0x4a, 0xc6, 0x8c, 0x86, 0x02, 0xc6, 0x7c,
	0x4a, 0x12, 0x93, 0xad, 0x17, 0xbc, 0xef, 0xe0, 0x0b, 0x8b, 0x06, 0x16, 0xc4, 0x9f, 0xa0, 0xd7,
	0xb9, 0x2e, 0x3f, 0x28, 0x35, 0x99, 0x7b, 0xc1, 0xb6, 0x03, 0xdc, 0xf9, 0x28, 0xfe, 0x16, 0xbd,
	0x86, 0x48, 0x86, 0x8f, 0x87, 0xf6, 0xd5, 0xa0, 0xb3, 0x3e, 0x94, 0xa5, 0xb1, 0xdd, 0x86, 0x48,
	0x16, 0x7f, 0x0c, 0x3f, 0xd5, 0xf9, 0x7c, 0x8c, 0x3e, 0x7c, 0x2a, 0x9f, 0x22, 0xf9, 0xec, 0xf8,
	0xaf, 0xdf, 0xfe, 0xf9, 0xb7, 0xba, 0xb5, 0x53, 0x46, 0x4d, 0xc6, 0xed, 0x6e, 0x99, 0xe0, 0xbf,
	0xce, 0x56, 0x6e, 0xcb, 0x59, 0xfe, 0xb2, 0xca, 0x1b, 0xfd, 0xcc, 0xdf, 0x94, 0xee, 0xaa, 0xe6,
	0xbd, 0x3f, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x02, 0x30, 0x11, 0x0b, 0x07, 0x00, 0x00,
}
