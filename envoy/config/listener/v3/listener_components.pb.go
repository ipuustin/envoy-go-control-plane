// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3/listener_components.proto

package envoy_config_listener_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type FilterChainMatch_ConnectionSourceType int32

const (
	FilterChainMatch_ANY                 FilterChainMatch_ConnectionSourceType = 0
	FilterChainMatch_SAME_IP_OR_LOOPBACK FilterChainMatch_ConnectionSourceType = 1
	FilterChainMatch_EXTERNAL            FilterChainMatch_ConnectionSourceType = 2
)

var FilterChainMatch_ConnectionSourceType_name = map[int32]string{
	0: "ANY",
	1: "SAME_IP_OR_LOOPBACK",
	2: "EXTERNAL",
}

var FilterChainMatch_ConnectionSourceType_value = map[string]int32{
	"ANY":                 0,
	"SAME_IP_OR_LOOPBACK": 1,
	"EXTERNAL":            2,
}

func (x FilterChainMatch_ConnectionSourceType) String() string {
	return proto.EnumName(FilterChainMatch_ConnectionSourceType_name, int32(x))
}

func (FilterChainMatch_ConnectionSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{1, 0}
}

type Filter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_HiddenEnvoyDeprecatedConfig
	//	*Filter_TypedConfig
	ConfigType           isFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Filter_HiddenEnvoyDeprecatedConfig) isFilter_ConfigType() {}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *Filter) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Filter_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_HiddenEnvoyDeprecatedConfig)(nil),
		(*Filter_TypedConfig)(nil),
	}
}

type FilterChainMatch struct {
	DestinationPort      *wrappers.UInt32Value                 `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	PrefixRanges         []*v3.CidrRange                       `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	AddressSuffix        string                                `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	SuffixLen            *wrappers.UInt32Value                 `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	SourceType           FilterChainMatch_ConnectionSourceType `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=envoy.config.listener.v3.FilterChainMatch_ConnectionSourceType" json:"source_type,omitempty"`
	SourcePrefixRanges   []*v3.CidrRange                       `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	SourcePorts          []uint32                              `protobuf:"varint,7,rep,packed,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	ServerNames          []string                              `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	TransportProtocol    string                                `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	ApplicationProtocols []string                              `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*v3.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourceType() FilterChainMatch_ConnectionSourceType {
	if m != nil {
		return m.SourceType
	}
	return FilterChainMatch_ANY
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*v3.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []uint32 {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

type FilterChain struct {
	FilterChainMatch                *FilterChainMatch         `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	HiddenEnvoyDeprecatedTlsContext *v31.DownstreamTlsContext `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_tls_context,json=hiddenEnvoyDeprecatedTlsContext,proto3" json:"hidden_envoy_deprecated_tls_context,omitempty"` // Deprecated: Do not use.
	Filters                         []*Filter                 `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	UseProxyProto                   *wrappers.BoolValue       `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	Metadata                        *v3.Metadata              `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	TransportSocket                 *v3.TransportSocket       `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	Name                            string                    `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                  `json:"-"`
	XXX_unrecognized                []byte                    `json:"-"`
	XXX_sizecache                   int32                     `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

// Deprecated: Do not use.
func (m *FilterChain) GetHiddenEnvoyDeprecatedTlsContext() *v31.DownstreamTlsContext {
	if m != nil {
		return m.HiddenEnvoyDeprecatedTlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *v3.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *v3.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func (m *FilterChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListenerFilterChainMatchPredicate struct {
	// Types that are valid to be assigned to Rule:
	//	*ListenerFilterChainMatchPredicate_OrMatch
	//	*ListenerFilterChainMatchPredicate_AndMatch
	//	*ListenerFilterChainMatchPredicate_NotMatch
	//	*ListenerFilterChainMatchPredicate_AnyMatch
	//	*ListenerFilterChainMatchPredicate_DestinationPortRange
	Rule                 isListenerFilterChainMatchPredicate_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ListenerFilterChainMatchPredicate) Reset()         { *m = ListenerFilterChainMatchPredicate{} }
func (m *ListenerFilterChainMatchPredicate) String() string { return proto.CompactTextString(m) }
func (*ListenerFilterChainMatchPredicate) ProtoMessage()    {}
func (*ListenerFilterChainMatchPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{3}
}

func (m *ListenerFilterChainMatchPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate.Unmarshal(m, b)
}
func (m *ListenerFilterChainMatchPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate.Marshal(b, m, deterministic)
}
func (m *ListenerFilterChainMatchPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilterChainMatchPredicate.Merge(m, src)
}
func (m *ListenerFilterChainMatchPredicate) XXX_Size() int {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate.Size(m)
}
func (m *ListenerFilterChainMatchPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilterChainMatchPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilterChainMatchPredicate proto.InternalMessageInfo

type isListenerFilterChainMatchPredicate_Rule interface {
	isListenerFilterChainMatchPredicate_Rule()
}

type ListenerFilterChainMatchPredicate_OrMatch struct {
	OrMatch *ListenerFilterChainMatchPredicate_MatchSet `protobuf:"bytes,1,opt,name=or_match,json=orMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_AndMatch struct {
	AndMatch *ListenerFilterChainMatchPredicate_MatchSet `protobuf:"bytes,2,opt,name=and_match,json=andMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_NotMatch struct {
	NotMatch *ListenerFilterChainMatchPredicate `protobuf:"bytes,3,opt,name=not_match,json=notMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_AnyMatch struct {
	AnyMatch bool `protobuf:"varint,4,opt,name=any_match,json=anyMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_DestinationPortRange struct {
	DestinationPortRange *v32.Int32Range `protobuf:"bytes,5,opt,name=destination_port_range,json=destinationPortRange,proto3,oneof"`
}

func (*ListenerFilterChainMatchPredicate_OrMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_AndMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_NotMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_AnyMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_DestinationPortRange) isListenerFilterChainMatchPredicate_Rule() {
}

func (m *ListenerFilterChainMatchPredicate) GetRule() isListenerFilterChainMatchPredicate_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetOrMatch() *ListenerFilterChainMatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_OrMatch); ok {
		return x.OrMatch
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetAndMatch() *ListenerFilterChainMatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_AndMatch); ok {
		return x.AndMatch
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetNotMatch() *ListenerFilterChainMatchPredicate {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_NotMatch); ok {
		return x.NotMatch
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetAnyMatch() bool {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_AnyMatch); ok {
		return x.AnyMatch
	}
	return false
}

func (m *ListenerFilterChainMatchPredicate) GetDestinationPortRange() *v32.Int32Range {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_DestinationPortRange); ok {
		return x.DestinationPortRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilterChainMatchPredicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilterChainMatchPredicate_OrMatch)(nil),
		(*ListenerFilterChainMatchPredicate_AndMatch)(nil),
		(*ListenerFilterChainMatchPredicate_NotMatch)(nil),
		(*ListenerFilterChainMatchPredicate_AnyMatch)(nil),
		(*ListenerFilterChainMatchPredicate_DestinationPortRange)(nil),
	}
}

type ListenerFilterChainMatchPredicate_MatchSet struct {
	Rules                []*ListenerFilterChainMatchPredicate `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ListenerFilterChainMatchPredicate_MatchSet) Reset() {
	*m = ListenerFilterChainMatchPredicate_MatchSet{}
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) String() string {
	return proto.CompactTextString(m)
}
func (*ListenerFilterChainMatchPredicate_MatchSet) ProtoMessage() {}
func (*ListenerFilterChainMatchPredicate_MatchSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{3, 0}
}

func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Unmarshal(m, b)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Marshal(b, m, deterministic)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Merge(m, src)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Size() int {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Size(m)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet proto.InternalMessageInfo

func (m *ListenerFilterChainMatchPredicate_MatchSet) GetRules() []*ListenerFilterChainMatchPredicate {
	if m != nil {
		return m.Rules
	}
	return nil
}

type ListenerFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_HiddenEnvoyDeprecatedConfig
	//	*ListenerFilter_TypedConfig
	ConfigType           isListenerFilter_ConfigType        `protobuf_oneof:"config_type"`
	FilterDisabled       *ListenerFilterChainMatchPredicate `protobuf:"bytes,4,opt,name=filter_disabled,json=filterDisabled,proto3" json:"filter_disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{4}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ListenerFilter_HiddenEnvoyDeprecatedConfig) isListenerFilter_ConfigType() {}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *ListenerFilter) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ListenerFilter_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

func (m *ListenerFilter) GetFilterDisabled() *ListenerFilterChainMatchPredicate {
	if m != nil {
		return m.FilterDisabled
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilter_HiddenEnvoyDeprecatedConfig)(nil),
		(*ListenerFilter_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.listener.v3.FilterChainMatch_ConnectionSourceType", FilterChainMatch_ConnectionSourceType_name, FilterChainMatch_ConnectionSourceType_value)
	proto.RegisterType((*Filter)(nil), "envoy.config.listener.v3.Filter")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.config.listener.v3.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.config.listener.v3.FilterChain")
	proto.RegisterType((*ListenerFilterChainMatchPredicate)(nil), "envoy.config.listener.v3.ListenerFilterChainMatchPredicate")
	proto.RegisterType((*ListenerFilterChainMatchPredicate_MatchSet)(nil), "envoy.config.listener.v3.ListenerFilterChainMatchPredicate.MatchSet")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.config.listener.v3.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3/listener_components.proto", fileDescriptor_87f255d2eccc91b5)
}

var fileDescriptor_87f255d2eccc91b5 = []byte{
	// 1238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xce, 0xda, 0x8e, 0xb3, 0x1e, 0xe7, 0x87, 0x19, 0x02, 0xd9, 0xa6, 0x25, 0x75, 0x5c, 0x1a,
	0x59, 0x45, 0x5d, 0x4b, 0x36, 0x12, 0xaa, 0x2b, 0x01, 0x59, 0x27, 0x25, 0x6d, 0xd3, 0xd6, 0xdd,
	0x04, 0x54, 0xe0, 0xb0, 0x9a, 0xec, 0x8e, 0xd3, 0x2d, 0x9b, 0x99, 0xd5, 0xcc, 0xd8, 0x8d, 0x6f,
	0x88, 0x13, 0xe2, 0x46, 0x8f, 0xfc, 0x15, 0x9c, 0x39, 0x22, 0x21, 0xb8, 0xf2, 0x8f, 0x70, 0x46,
	0x39, 0x50, 0x34, 0x3f, 0xd6, 0x49, 0x9c, 0x84, 0x44, 0x84, 0x03, 0x37, 0xcf, 0xbc, 0xf7, 0xbe,
	0x79, 0xef, 0xcd, 0xf7, 0xbe, 0x59, 0x83, 0x26, 0x26, 0x03, 0x3a, 0x6c, 0x84, 0x94, 0xf4, 0xe2,
	0xdd, 0x46, 0x12, 0x73, 0x81, 0x09, 0x66, 0x8d, 0x41, 0x6b, 0xf4, 0x3b, 0x08, 0xe9, 0x5e, 0x4a,
	0x09, 0x26, 0x82, 0xbb, 0x29, 0xa3, 0x82, 0x42, 0x47, 0xc5, 0xb8, 0x3a, 0xc6, 0xcd, 0xfc, 0xdc,
	0x41, 0x6b, 0xb1, 0x76, 0x0c, 0x2d, 0xa4, 0x0c, 0x4b, 0x24, 0x14, 0x45, 0x0c, 0x73, 0x13, 0xbd,
	0x78, 0xfd, 0x54, 0x9f, 0x1d, 0xc4, 0xb1, 0x71, 0x78, 0x5f, 0x3b, 0xe0, 0x7d, 0x81, 0x09, 0x8f,
	0x29, 0xe1, 0x0d, 0xc1, 0x10, 0xe1, 0x29, 0x65, 0x22, 0xe0, 0x34, 0xfc, 0x0a, 0x0b, 0xde, 0x10,
	0x09, 0x97, 0x51, 0x21, 0x66, 0xc2, 0x44, 0x5d, 0xd1, 0x51, 0x62, 0x98, 0x2a, 0x3c, 0x86, 0xc8,
	0x6e, 0x06, 0x78, 0x65, 0x97, 0xd2, 0xdd, 0x04, 0x37, 0xd4, 0x6a, 0xa7, 0xdf, 0x6b, 0x20, 0x32,
	0x34, 0xa6, 0x6b, 0xe3, 0x26, 0x2e, 0x58, 0x3f, 0xcc, 0x30, 0x97, 0xc6, 0xad, 0x2f, 0x19, 0x4a,
	0x53, 0xcc, 0xb2, 0x52, 0xde, 0xe9, 0x47, 0x29, 0x6a, 0x20, 0x42, 0xa8, 0x40, 0x42, 0x65, 0xca,
	0x05, 0x12, 0xfd, 0xcc, 0xbc, 0x7c, 0xc2, 0x3c, 0xc0, 0x4c, 0x56, 0x14, 0x93, 0x5d, 0xe3, 0xb2,
	0x30, 0x40, 0x49, 0x1c, 0x21, 0x81, 0x1b, 0xd9, 0x0f, 0x6d, 0xa8, 0xfd, 0x65, 0x81, 0xe2, 0xbd,
	0x38, 0x11, 0x98, 0xc1, 0xab, 0xa0, 0x40, 0xd0, 0x1e, 0x76, 0xac, 0xaa, 0x55, 0x2f, 0x79, 0x53,
	0x07, 0x5e, 0x81, 0xe5, 0xaa, 0x96, 0xaf, 0x36, 0xe1, 0x0e, 0x58, 0x7a, 0x1e, 0x47, 0x11, 0x26,
	0x81, 0xaa, 0x3f, 0x88, 0x70, 0xca, 0x70, 0x88, 0x04, 0x8e, 0x02, 0xdd, 0x61, 0x27, 0x57, 0xb5,
	0xea, 0xe5, 0xe6, 0x82, 0xab, 0x6b, 0x71, 0xb3, 0x5a, 0xdc, 0x2d, 0x55, 0xa9, 0x97, 0x73, 0xac,
	0x8d, 0x09, 0xff, 0xaa, 0x06, 0x59, 0x97, 0x18, 0x6b, 0x23, 0x88, 0x8e, 0x42, 0x80, 0x77, 0xc0,
	0xb4, 0x6c, 0xeb, 0x08, 0xb1, 0xa0, 0x10, 0xe7, 0x4f, 0x20, 0xae, 0x92, 0xe1, 0xc6, 0x84, 0x5f,
	0x56, 0xbe, 0x3a, 0xb4, 0x7d, 0xe3, 0x87, 0x5f, 0xbe, 0x5d, 0x5a, 0x02, 0xd7, 0x34, 0x63, 0x50,
	0x1a, 0xbb, 0x83, 0xe6, 0x21, 0x63, 0x74, 0x81, 0xde, 0x0c, 0x28, 0x6b, 0xe4, 0x40, 0x86, 0x3e,
	0x28, 0xd8, 0xf9, 0x4a, 0xa1, 0xf6, 0x5d, 0x11, 0x54, 0xb4, 0xbd, 0xf3, 0x1c, 0xc5, 0xe4, 0x11,
	0x12, 0xe1, 0x73, 0xb8, 0x0d, 0x2a, 0x11, 0xe6, 0x22, 0x26, 0xaa, 0x9f, 0x81, 0xa4, 0x84, 0x63,
	0xab, 0x6c, 0xae, 0x9d, 0xc8, 0xe6, 0xd3, 0xfb, 0x44, 0xb4, 0x9a, 0x9f, 0xa1, 0xa4, 0x8f, 0xbd,
	0xf2, 0x81, 0x67, 0xdf, 0x2a, 0x3a, 0xaf, 0x5f, 0xe7, 0xeb, 0x96, 0x3f, 0x77, 0x04, 0xa2, 0x4b,
	0x99, 0x80, 0x6b, 0x60, 0x26, 0x65, 0xb8, 0x17, 0xef, 0x07, 0x8a, 0x35, 0xdc, 0xc9, 0x57, 0xf3,
	0xf5, 0x72, 0xf3, 0xba, 0x7b, 0x8c, 0xe7, 0x92, 0xa9, 0xee, 0xa0, 0xe5, 0x76, 0xe2, 0x88, 0xf9,
	0xd2, 0xcf, 0x9f, 0xd6, 0x51, 0x6a, 0xc1, 0xe1, 0x4d, 0x30, 0x6b, 0x88, 0x1e, 0xf0, 0x7e, 0xaf,
	0x17, 0xef, 0xab, 0x3e, 0x95, 0xfc, 0x19, 0xb3, 0xbb, 0xa5, 0x36, 0xe1, 0x5d, 0x00, 0xb4, 0x39,
	0x48, 0x30, 0x71, 0x26, 0xcf, 0x4f, 0xde, 0x2f, 0x69, 0xff, 0x4d, 0x4c, 0xe0, 0x0b, 0x50, 0xe6,
	0xb4, 0xcf, 0x42, 0xac, 0x3a, 0xe5, 0x4c, 0x57, 0xad, 0xfa, 0x6c, 0xf3, 0x23, 0xf7, 0xac, 0x79,
	0x74, 0xc7, 0x1b, 0xe8, 0x76, 0x28, 0x21, 0x38, 0x94, 0x95, 0x6f, 0x29, 0x9c, 0xed, 0x61, 0x8a,
	0x3d, 0xfb, 0xc0, 0x9b, 0xfc, 0xc6, 0xca, 0x55, 0x2c, 0x1f, 0xf0, 0xd1, 0x2e, 0x7c, 0x0a, 0xe6,
	0xcd, 0x59, 0xc7, 0x9b, 0x53, 0xbc, 0x58, 0x73, 0xa0, 0x0e, 0xee, 0x1e, 0x6d, 0x51, 0x0b, 0x4c,
	0x67, 0x90, 0x94, 0x09, 0xee, 0x4c, 0x55, 0xf3, 0xf5, 0x19, 0xaf, 0x72, 0xe0, 0xcd, 0xbc, 0xb2,
	0x40, 0xed, 0xf0, 0x86, 0x4c, 0x91, 0xf2, 0x72, 0x38, 0x5c, 0x06, 0xd3, 0x1c, 0xb3, 0x01, 0x66,
	0x81, 0x24, 0x3c, 0x77, 0xca, 0xd5, 0x7c, 0xbd, 0xe4, 0x97, 0xf5, 0xde, 0x63, 0xb9, 0x05, 0x6f,
	0x03, 0x78, 0x28, 0x11, 0xaa, 0x87, 0x21, 0x4d, 0x9c, 0x92, 0x6a, 0xff, 0x1b, 0x23, 0x4b, 0xd7,
	0x18, 0x60, 0x0b, 0xbc, 0x85, 0xd2, 0x34, 0x89, 0x43, 0xc3, 0x22, 0xb3, 0xcf, 0x1d, 0xa0, 0xa0,
	0xe7, 0x8f, 0x18, 0xb3, 0x18, 0x5e, 0xbb, 0x07, 0xe6, 0x4f, 0x6b, 0x1e, 0x9c, 0x02, 0xf9, 0xd5,
	0xc7, 0x9f, 0x57, 0x26, 0xe0, 0x02, 0x78, 0x73, 0x6b, 0xf5, 0xd1, 0x7a, 0x70, 0xbf, 0x1b, 0x3c,
	0xf1, 0x83, 0xcd, 0x27, 0x4f, 0xba, 0xde, 0x6a, 0xe7, 0x61, 0xc5, 0x82, 0xd3, 0xc0, 0x5e, 0x7f,
	0xb6, 0xbd, 0xee, 0x3f, 0x5e, 0xdd, 0xac, 0xe4, 0xda, 0xb7, 0xe5, 0x44, 0xd4, 0xc1, 0xca, 0x3f,
	0x4d, 0xc4, 0xe1, 0x85, 0x3d, 0x28, 0xd8, 0x56, 0x25, 0x57, 0xfb, 0xb5, 0x00, 0xca, 0x47, 0x4c,
	0xf0, 0x19, 0x80, 0x3d, 0xb5, 0x0c, 0x42, 0xb9, 0x0e, 0xf6, 0xa4, 0xaf, 0x12, 0x88, 0x72, 0xf3,
	0xd6, 0xc5, 0xe9, 0xe0, 0x57, 0x7a, 0xe3, 0x13, 0xf6, 0xbd, 0x05, 0x6e, 0x9c, 0x25, 0x28, 0x22,
	0xe1, 0x52, 0x02, 0x04, 0xde, 0x17, 0x46, 0x55, 0x32, 0xea, 0x1d, 0x6a, 0xb5, 0x7b, 0x42, 0xab,
	0x5d, 0x91, 0x70, 0x79, 0xf8, 0x1a, 0x7d, 0x49, 0xb8, 0x60, 0x18, 0xed, 0x6d, 0x27, 0xbc, 0xa3,
	0x61, 0xa4, 0xfa, 0xf8, 0xd7, 0x4f, 0xd5, 0x9e, 0x43, 0x27, 0xd8, 0x06, 0x53, 0x3a, 0xcf, 0x6c,
	0x32, 0xab, 0xe7, 0x95, 0xe8, 0x67, 0x01, 0xd0, 0x03, 0x73, 0x7d, 0x2e, 0x29, 0x4c, 0xf7, 0x87,
	0xfa, 0xa6, 0x8d, 0x7c, 0x2d, 0x9e, 0x98, 0x39, 0x8f, 0xd2, 0x44, 0x4f, 0xdc, 0x4c, 0x9f, 0xe3,
	0xae, 0x8c, 0x50, 0xd7, 0x0f, 0xdb, 0xc0, 0xde, 0xc3, 0x02, 0x45, 0x48, 0x20, 0x33, 0xb0, 0x4b,
	0xa7, 0xb3, 0xff, 0x91, 0xf1, 0xf2, 0x47, 0xfe, 0xb0, 0x0b, 0x2a, 0xe3, 0x1d, 0x71, 0x8a, 0x0a,
	0xe3, 0xe6, 0xe9, 0x18, 0xdb, 0x99, 0xf7, 0x96, 0x72, 0xf6, 0xe7, 0xc4, 0xf1, 0x0d, 0x08, 0xcd,
	0x73, 0x30, 0xa5, 0xe8, 0xad, 0x7e, 0xb7, 0xeb, 0x92, 0x54, 0x37, 0xc0, 0xf2, 0xb9, 0xa4, 0xaa,
	0xfd, 0x3c, 0x09, 0x96, 0x37, 0x8d, 0x61, 0x9c, 0x0e, 0x5d, 0x86, 0x23, 0xc9, 0x7b, 0x0c, 0x11,
	0xb0, 0x29, 0x3b, 0xc6, 0xaa, 0xb5, 0xb3, 0x5b, 0x7e, 0x2e, 0x9c, 0xab, 0x96, 0x5b, 0x58, 0x6c,
	0x4c, 0xf8, 0x53, 0x94, 0x69, 0xa2, 0x85, 0xa0, 0x84, 0x48, 0x64, 0xce, 0xc8, 0xfd, 0xa7, 0x67,
	0xd8, 0x88, 0x44, 0xfa, 0x90, 0x2f, 0x40, 0x89, 0x50, 0x61, 0x0e, 0xc9, 0xab, 0x43, 0xee, 0x5e,
	0xe2, 0x10, 0x89, 0x4d, 0xa8, 0xd0, 0xd8, 0x2b, 0xb2, 0x80, 0xa1, 0xc1, 0x96, 0x9c, 0xb2, 0xd5,
	0xdb, 0xfc, 0x22, 0x67, 0x5b, 0x3a, 0x87, 0xa1, 0xf6, 0x7b, 0x0a, 0xde, 0x1e, 0x7f, 0xb3, 0xb4,
	0x94, 0x1a, 0x2e, 0x5d, 0x31, 0x09, 0x49, 0x45, 0x97, 0x59, 0x28, 0xe5, 0x57, 0x82, 0xb9, 0x31,
	0xe1, 0xcf, 0x8f, 0xbd, 0x55, 0x6a, 0x7f, 0xf1, 0x47, 0x0b, 0xd8, 0x59, 0xbd, 0xf0, 0x4b, 0x30,
	0xc9, 0xfa, 0x09, 0xe6, 0x8e, 0xa5, 0x66, 0xe3, 0x32, 0xf5, 0xa9, 0x97, 0xe0, 0x95, 0x95, 0xb3,
	0x73, 0xbe, 0xc6, 0x6c, 0x7f, 0x22, 0x89, 0xe5, 0x81, 0x8f, 0x4f, 0x27, 0xd6, 0xc5, 0x6f, 0xa5,
	0xfd, 0xa1, 0x04, 0xba, 0x03, 0x3e, 0xf8, 0x97, 0x40, 0x5e, 0x19, 0x14, 0x64, 0x46, 0x30, 0xff,
	0xa7, 0x67, 0xd5, 0xfe, 0xc8, 0x81, 0xd9, 0xe3, 0x21, 0xff, 0xbf, 0x8f, 0xa4, 0xfc, 0x85, 0x3f,
	0x92, 0x60, 0x04, 0xe6, 0x8c, 0x9a, 0x47, 0x31, 0x47, 0x3b, 0x09, 0x8e, 0x8c, 0x46, 0x5d, 0xe6,
	0x2e, 0xfd, 0x59, 0x8d, 0xb9, 0x66, 0x20, 0xdb, 0xef, 0xc9, 0x1b, 0x58, 0x01, 0xef, 0x5e, 0xe4,
	0x06, 0xc6, 0x3e, 0xc9, 0xbc, 0x87, 0x3f, 0x7d, 0xfd, 0xdb, 0xef, 0xc5, 0x5c, 0x25, 0x0f, 0x56,
	0x62, 0xaa, 0x93, 0x52, 0x8a, 0x7a, 0x66, 0x7e, 0xde, 0x42, 0x06, 0xd8, 0x19, 0xfd, 0x7b, 0x50,
	0x62, 0xda, 0xb5, 0x76, 0x8a, 0xaa, 0x17, 0xad, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x39, 0x66,
	0xe4, 0x6f, 0x7b, 0x0c, 0x00, 0x00,
}
