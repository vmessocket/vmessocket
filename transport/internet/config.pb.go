// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: transport/internet/config.proto

package internet

import (
	serial "github.com/vmessocket/vmessocket/common/serial"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransportProtocol int32

const (
	TransportProtocol_TCP       TransportProtocol = 0
	TransportProtocol_UDP       TransportProtocol = 1
	TransportProtocol_WebSocket TransportProtocol = 2
	TransportProtocol_HTTP      TransportProtocol = 3
)

// Enum value maps for TransportProtocol.
var (
	TransportProtocol_name = map[int32]string{
		0: "TCP",
		1: "UDP",
		2: "WebSocket",
		3: "HTTP",
	}
	TransportProtocol_value = map[string]int32{
		"TCP":       0,
		"UDP":       1,
		"WebSocket": 2,
		"HTTP":      3,
	}
)

func (x TransportProtocol) Enum() *TransportProtocol {
	p := new(TransportProtocol)
	*p = x
	return p
}

func (x TransportProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransportProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_config_proto_enumTypes[0].Descriptor()
}

func (TransportProtocol) Type() protoreflect.EnumType {
	return &file_transport_internet_config_proto_enumTypes[0]
}

func (x TransportProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransportProtocol.Descriptor instead.
func (TransportProtocol) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{0}
}

type SocketConfig_TCPFastOpenState int32

const (
	SocketConfig_AsIs    SocketConfig_TCPFastOpenState = 0
	SocketConfig_Enable  SocketConfig_TCPFastOpenState = 1
	SocketConfig_Disable SocketConfig_TCPFastOpenState = 2
)

// Enum value maps for SocketConfig_TCPFastOpenState.
var (
	SocketConfig_TCPFastOpenState_name = map[int32]string{
		0: "AsIs",
		1: "Enable",
		2: "Disable",
	}
	SocketConfig_TCPFastOpenState_value = map[string]int32{
		"AsIs":    0,
		"Enable":  1,
		"Disable": 2,
	}
)

func (x SocketConfig_TCPFastOpenState) Enum() *SocketConfig_TCPFastOpenState {
	p := new(SocketConfig_TCPFastOpenState)
	*p = x
	return p
}

func (x SocketConfig_TCPFastOpenState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketConfig_TCPFastOpenState) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_config_proto_enumTypes[1].Descriptor()
}

func (SocketConfig_TCPFastOpenState) Type() protoreflect.EnumType {
	return &file_transport_internet_config_proto_enumTypes[1]
}

func (x SocketConfig_TCPFastOpenState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketConfig_TCPFastOpenState.Descriptor instead.
func (SocketConfig_TCPFastOpenState) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{3, 0}
}

type SocketConfig_TProxyMode int32

const (
	SocketConfig_Off      SocketConfig_TProxyMode = 0
	SocketConfig_TProxy   SocketConfig_TProxyMode = 1
	SocketConfig_Redirect SocketConfig_TProxyMode = 2
)

// Enum value maps for SocketConfig_TProxyMode.
var (
	SocketConfig_TProxyMode_name = map[int32]string{
		0: "Off",
		1: "TProxy",
		2: "Redirect",
	}
	SocketConfig_TProxyMode_value = map[string]int32{
		"Off":      0,
		"TProxy":   1,
		"Redirect": 2,
	}
)

func (x SocketConfig_TProxyMode) Enum() *SocketConfig_TProxyMode {
	p := new(SocketConfig_TProxyMode)
	*p = x
	return p
}

func (x SocketConfig_TProxyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketConfig_TProxyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_config_proto_enumTypes[2].Descriptor()
}

func (SocketConfig_TProxyMode) Type() protoreflect.EnumType {
	return &file_transport_internet_config_proto_enumTypes[2]
}

func (x SocketConfig_TProxyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketConfig_TProxyMode.Descriptor instead.
func (SocketConfig_TProxyMode) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{3, 1}
}

type TransportConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Protocol     TransportProtocol    `protobuf:"varint,1,opt,name=protocol,proto3,enum=vmessocket.core.transport.internet.TransportProtocol" json:"protocol,omitempty"`
	ProtocolName string               `protobuf:"bytes,3,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	Settings     *serial.TypedMessage `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *TransportConfig) Reset() {
	*x = TransportConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportConfig) ProtoMessage() {}

func (x *TransportConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportConfig.ProtoReflect.Descriptor instead.
func (*TransportConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *TransportConfig) GetProtocol() TransportProtocol {
	if x != nil {
		return x.Protocol
	}
	return TransportProtocol_TCP
}

func (x *TransportConfig) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

func (x *TransportConfig) GetSettings() *serial.TypedMessage {
	if x != nil {
		return x.Settings
	}
	return nil
}

type StreamConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Protocol          TransportProtocol      `protobuf:"varint,1,opt,name=protocol,proto3,enum=vmessocket.core.transport.internet.TransportProtocol" json:"protocol,omitempty"`
	ProtocolName      string                 `protobuf:"bytes,5,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	TransportSettings []*TransportConfig     `protobuf:"bytes,2,rep,name=transport_settings,json=transportSettings,proto3" json:"transport_settings,omitempty"`
	SecurityType      string                 `protobuf:"bytes,3,opt,name=security_type,json=securityType,proto3" json:"security_type,omitempty"`
	SecuritySettings  []*serial.TypedMessage `protobuf:"bytes,4,rep,name=security_settings,json=securitySettings,proto3" json:"security_settings,omitempty"`
	SocketSettings    *SocketConfig          `protobuf:"bytes,6,opt,name=socket_settings,json=socketSettings,proto3" json:"socket_settings,omitempty"`
}

func (x *StreamConfig) Reset() {
	*x = StreamConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamConfig) ProtoMessage() {}

func (x *StreamConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamConfig.ProtoReflect.Descriptor instead.
func (*StreamConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Do not use.
func (x *StreamConfig) GetProtocol() TransportProtocol {
	if x != nil {
		return x.Protocol
	}
	return TransportProtocol_TCP
}

func (x *StreamConfig) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

func (x *StreamConfig) GetTransportSettings() []*TransportConfig {
	if x != nil {
		return x.TransportSettings
	}
	return nil
}

func (x *StreamConfig) GetSecurityType() string {
	if x != nil {
		return x.SecurityType
	}
	return ""
}

func (x *StreamConfig) GetSecuritySettings() []*serial.TypedMessage {
	if x != nil {
		return x.SecuritySettings
	}
	return nil
}

func (x *StreamConfig) GetSocketSettings() *SocketConfig {
	if x != nil {
		return x.SocketSettings
	}
	return nil
}

type ProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransportLayerProxy bool `protobuf:"varint,2,opt,name=transportLayerProxy,proto3" json:"transportLayerProxy,omitempty"`
}

func (x *ProxyConfig) Reset() {
	*x = ProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig) ProtoMessage() {}

func (x *ProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig.ProtoReflect.Descriptor instead.
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyConfig) GetTransportLayerProxy() bool {
	if x != nil {
		return x.TransportLayerProxy
	}
	return false
}

type SocketConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark                       uint32                        `protobuf:"varint,1,opt,name=mark,proto3" json:"mark,omitempty"`
	Tfo                        SocketConfig_TCPFastOpenState `protobuf:"varint,2,opt,name=tfo,proto3,enum=vmessocket.core.transport.internet.SocketConfig_TCPFastOpenState" json:"tfo,omitempty"`
	Tproxy                     SocketConfig_TProxyMode       `protobuf:"varint,3,opt,name=tproxy,proto3,enum=vmessocket.core.transport.internet.SocketConfig_TProxyMode" json:"tproxy,omitempty"`
	ReceiveOriginalDestAddress bool                          `protobuf:"varint,4,opt,name=receive_original_dest_address,json=receiveOriginalDestAddress,proto3" json:"receive_original_dest_address,omitempty"`
	BindAddress                []byte                        `protobuf:"bytes,5,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort                   uint32                        `protobuf:"varint,6,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	TcpKeepAliveInterval       int32                         `protobuf:"varint,7,opt,name=tcp_keep_alive_interval,json=tcpKeepAliveInterval,proto3" json:"tcp_keep_alive_interval,omitempty"`
	TfoQueueLength             uint32                        `protobuf:"varint,8,opt,name=tfo_queue_length,json=tfoQueueLength,proto3" json:"tfo_queue_length,omitempty"`
}

func (x *SocketConfig) Reset() {
	*x = SocketConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketConfig) ProtoMessage() {}

func (x *SocketConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketConfig.ProtoReflect.Descriptor instead.
func (*SocketConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{3}
}

func (x *SocketConfig) GetMark() uint32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *SocketConfig) GetTfo() SocketConfig_TCPFastOpenState {
	if x != nil {
		return x.Tfo
	}
	return SocketConfig_AsIs
}

func (x *SocketConfig) GetTproxy() SocketConfig_TProxyMode {
	if x != nil {
		return x.Tproxy
	}
	return SocketConfig_Off
}

func (x *SocketConfig) GetReceiveOriginalDestAddress() bool {
	if x != nil {
		return x.ReceiveOriginalDestAddress
	}
	return false
}

func (x *SocketConfig) GetBindAddress() []byte {
	if x != nil {
		return x.BindAddress
	}
	return nil
}

func (x *SocketConfig) GetBindPort() uint32 {
	if x != nil {
		return x.BindPort
	}
	return 0
}

func (x *SocketConfig) GetTcpKeepAliveInterval() int32 {
	if x != nil {
		return x.TcpKeepAliveInterval
	}
	return 0
}

func (x *SocketConfig) GetTfoQueueLength() uint32 {
	if x != nil {
		return x.TfoQueueLength
	}
	return 0
}

var File_transport_internet_config_proto protoreflect.FileDescriptor

var file_transport_internet_config_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35,
	0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x6d, 0x65,
	0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0xc8, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x55, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x62,
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x76, 0x6d, 0x65,
	0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x59, 0x0a, 0x0f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x76, 0x6d, 0x65,
	0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e,
	0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x3f, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x98, 0x04,
	0x0a, 0x0c, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x53, 0x0a, 0x03, 0x74, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x41, 0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x54, 0x43, 0x50, 0x46, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x03, 0x74, 0x66, 0x6f, 0x12, 0x53, 0x0a, 0x06, 0x74, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x41, 0x0a, 0x1d,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x35, 0x0a, 0x17, 0x74, 0x63, 0x70, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x14, 0x74, 0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x66, 0x6f, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x74, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x22, 0x35, 0x0a, 0x10, 0x54, 0x43, 0x50, 0x46, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x73, 0x49, 0x73, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x02, 0x22, 0x2f, 0x0a, 0x0a, 0x54, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x66, 0x66, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x10, 0x02, 0x2a, 0x3e, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x07, 0x0a,
	0x03, 0x54, 0x43, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x03, 0x42, 0x84, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x6d, 0x65,
	0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0xaa, 0x02, 0x22, 0x76, 0x6d, 0x65,
	0x73, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_config_proto_rawDescOnce sync.Once
	file_transport_internet_config_proto_rawDescData = file_transport_internet_config_proto_rawDesc
)

func file_transport_internet_config_proto_rawDescGZIP() []byte {
	file_transport_internet_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_config_proto_rawDescData)
	})
	return file_transport_internet_config_proto_rawDescData
}

var file_transport_internet_config_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_transport_internet_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transport_internet_config_proto_goTypes = []interface{}{
	(TransportProtocol)(0),             // 0: vmessocket.core.transport.internet.TransportProtocol
	(SocketConfig_TCPFastOpenState)(0), // 1: vmessocket.core.transport.internet.SocketConfig.TCPFastOpenState
	(SocketConfig_TProxyMode)(0),       // 2: vmessocket.core.transport.internet.SocketConfig.TProxyMode
	(*TransportConfig)(nil),            // 3: vmessocket.core.transport.internet.TransportConfig
	(*StreamConfig)(nil),               // 4: vmessocket.core.transport.internet.StreamConfig
	(*ProxyConfig)(nil),                // 5: vmessocket.core.transport.internet.ProxyConfig
	(*SocketConfig)(nil),               // 6: vmessocket.core.transport.internet.SocketConfig
	(*serial.TypedMessage)(nil),        // 7: vmessocket.core.common.serial.TypedMessage
}
var file_transport_internet_config_proto_depIdxs = []int32{
	0, // 0: vmessocket.core.transport.internet.TransportConfig.protocol:type_name -> vmessocket.core.transport.internet.TransportProtocol
	7, // 1: vmessocket.core.transport.internet.TransportConfig.settings:type_name -> vmessocket.core.common.serial.TypedMessage
	0, // 2: vmessocket.core.transport.internet.StreamConfig.protocol:type_name -> vmessocket.core.transport.internet.TransportProtocol
	3, // 3: vmessocket.core.transport.internet.StreamConfig.transport_settings:type_name -> vmessocket.core.transport.internet.TransportConfig
	7, // 4: vmessocket.core.transport.internet.StreamConfig.security_settings:type_name -> vmessocket.core.common.serial.TypedMessage
	6, // 5: vmessocket.core.transport.internet.StreamConfig.socket_settings:type_name -> vmessocket.core.transport.internet.SocketConfig
	1, // 6: vmessocket.core.transport.internet.SocketConfig.tfo:type_name -> vmessocket.core.transport.internet.SocketConfig.TCPFastOpenState
	2, // 7: vmessocket.core.transport.internet.SocketConfig.tproxy:type_name -> vmessocket.core.transport.internet.SocketConfig.TProxyMode
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_transport_internet_config_proto_init() }
func file_transport_internet_config_proto_init() {
	if File_transport_internet_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_internet_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_internet_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_internet_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transport_internet_config_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_config_proto_depIdxs,
		EnumInfos:         file_transport_internet_config_proto_enumTypes,
		MessageInfos:      file_transport_internet_config_proto_msgTypes,
	}.Build()
	File_transport_internet_config_proto = out.File
	file_transport_internet_config_proto_rawDesc = nil
	file_transport_internet_config_proto_goTypes = nil
	file_transport_internet_config_proto_depIdxs = nil
}
