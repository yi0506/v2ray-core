package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_app_proxyman "v2ray.com/core/app/proxyman"
import v2ray_core_common_serial "v2ray.com/core/common/serial"
import v2ray_core_transport "v2ray.com/core/transport"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration serialization format.
type ConfigFormat int32

const (
	ConfigFormat_Protobuf ConfigFormat = 0
	ConfigFormat_JSON     ConfigFormat = 1
)

var ConfigFormat_name = map[int32]string{
	0: "Protobuf",
	1: "JSON",
}
var ConfigFormat_value = map[string]int32{
	"Protobuf": 0,
	"JSON":     1,
}

func (x ConfigFormat) String() string {
	return proto.EnumName(ConfigFormat_name, int32(x))
}
func (ConfigFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Master config of V2Ray. V2Ray Core takes this config as input and functions accordingly.
type Config struct {
	// Inbound handler configurations. Must have at least one item.
	Inbound []*v2ray_core_app_proxyman.InboundHandlerConfig `protobuf:"bytes,1,rep,name=inbound" json:"inbound,omitempty"`
	// Outbound handler configurations. Must have at least one item. The first item is used as default for routing.
	Outbound []*v2ray_core_app_proxyman.OutboundHandlerConfig `protobuf:"bytes,2,rep,name=outbound" json:"outbound,omitempty"`
	// App configuration. Must be one in the app directory.
	App []*v2ray_core_common_serial.TypedMessage `protobuf:"bytes,4,rep,name=app" json:"app,omitempty"`
	// Transport settings.
	Transport *v2ray_core_transport.Config `protobuf:"bytes,5,opt,name=transport" json:"transport,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetInbound() []*v2ray_core_app_proxyman.InboundHandlerConfig {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *Config) GetOutbound() []*v2ray_core_app_proxyman.OutboundHandlerConfig {
	if m != nil {
		return m.Outbound
	}
	return nil
}

func (m *Config) GetApp() []*v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Config) GetTransport() *v2ray_core_transport.Config {
	if m != nil {
		return m.Transport
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.Config")
	proto.RegisterEnum("v2ray.core.ConfigFormat", ConfigFormat_name, ConfigFormat_value)
}

func init() { proto.RegisterFile("v2ray.com/core/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xdf, 0x6e, 0x7d, 0x67, 0x3d, 0x1b, 0x32, 0x72, 0x35, 0xa6, 0x17, 0x43, 0xd8, 0x18,
	0x82, 0xa9, 0xd4, 0x1b, 0xf1, 0xd2, 0x81, 0x7f, 0x06, 0xba, 0x31, 0xc5, 0x0b, 0x6f, 0x24, 0xeb,
	0xb2, 0x31, 0x58, 0x72, 0x42, 0x9a, 0x89, 0xfd, 0x28, 0x7e, 0x05, 0x3f, 0xa5, 0xb4, 0x69, 0xbb,
	0x76, 0xe2, 0x55, 0x20, 0xe7, 0xf9, 0xfd, 0x72, 0xf2, 0xc0, 0xf1, 0x47, 0xa0, 0x59, 0x4c, 0x43,
	0x14, 0x7e, 0x88, 0x9a, 0xfb, 0x21, 0xca, 0xe5, 0x7a, 0x45, 0x95, 0x46, 0x83, 0x04, 0xf2, 0xa1,
	0xe6, 0xdd, 0xe1, 0x5e, 0x90, 0x29, 0xe5, 0x2b, 0x8d, 0x9f, 0xb1, 0x60, 0xb2, 0x42, 0x75, 0x2f,
	0x7e, 0x29, 0x85, 0x40, 0xe9, 0x47, 0x5c, 0xaf, 0xd9, 0xc6, 0x37, 0xb1, 0xe2, 0x8b, 0x77, 0xc1,
	0xa3, 0x88, 0xad, 0x78, 0x46, 0xf4, 0xf7, 0x08, 0xa3, 0x99, 0x8c, 0x14, 0x6a, 0x53, 0x11, 0x9f,
	0x7e, 0xd5, 0xa0, 0x31, 0x4a, 0x2f, 0xc8, 0x1d, 0x1c, 0xac, 0xe5, 0x1c, 0xb7, 0x72, 0xd1, 0x71,
	0x7a, 0xf5, 0x61, 0x33, 0x38, 0xa7, 0xbb, 0x5d, 0x29, 0x53, 0x8a, 0xe6, 0xbb, 0xd1, 0x07, 0x9b,
	0xbb, 0x67, 0x72, 0xb1, 0xe1, 0xda, 0xf2, 0xb3, 0x9c, 0x26, 0x63, 0xf0, 0x70, 0x6b, 0xac, 0xa9,
	0x96, 0x9a, 0xe8, 0x9f, 0xa6, 0x49, 0x16, 0xac, 0xaa, 0x0a, 0x9e, 0x5c, 0x41, 0x9d, 0x29, 0xd5,
	0x71, 0x53, 0xcd, 0xa0, 0xac, 0xb1, 0x15, 0x50, 0x5b, 0x01, 0x7d, 0x49, 0x2a, 0x78, 0xb4, 0x0d,
	0xcc, 0x12, 0x84, 0x5c, 0xc3, 0x61, 0xf1, 0xe7, 0xce, 0xff, 0x9e, 0x33, 0x6c, 0x06, 0x27, 0x65,
	0xbe, 0x18, 0xd2, 0xec, 0xd1, 0x5d, 0x7c, 0xec, 0x7a, 0xf5, 0xb6, 0x7b, 0x36, 0x80, 0x96, 0x1d,
	0xdd, 0xa2, 0x16, 0xcc, 0x90, 0x16, 0x78, 0xd3, 0xa4, 0xb4, 0xf9, 0x76, 0xd9, 0xfe, 0x47, 0x3c,
	0x70, 0xc7, 0xcf, 0x93, 0xa7, 0xb6, 0x73, 0xd3, 0x87, 0xa3, 0x10, 0x45, 0xc9, 0x3d, 0x75, 0xde,
	0xdc, 0xe4, 0xfc, 0xae, 0xc1, 0x6b, 0x30, 0x63, 0x31, 0x1d, 0xa1, 0xe6, 0xf3, 0x46, 0xda, 0xf8,
	0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xbe, 0x63, 0xd9, 0x1f, 0x02, 0x00, 0x00,
}
