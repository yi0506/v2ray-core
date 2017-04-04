package tcp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_serial "v2ray.com/core/common/serial"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConnectionReuse struct {
	Enable bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
}

func (m *ConnectionReuse) Reset()                    { *m = ConnectionReuse{} }
func (m *ConnectionReuse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionReuse) ProtoMessage()               {}
func (*ConnectionReuse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnectionReuse) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type Config struct {
	ConnectionReuse *ConnectionReuse                       `protobuf:"bytes,1,opt,name=connection_reuse,json=connectionReuse" json:"connection_reuse,omitempty"`
	HeaderSettings  *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,2,opt,name=header_settings,json=headerSettings" json:"header_settings,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetConnectionReuse() *ConnectionReuse {
	if m != nil {
		return m.ConnectionReuse
	}
	return nil
}

func (m *Config) GetHeaderSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.HeaderSettings
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionReuse)(nil), "v2ray.core.transport.internet.tcp.ConnectionReuse")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.tcp.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/tcp/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0xfc, 0x50, 0x7e, 0x22, 0x58, 0xe9, 0x42, 0x06, 0x57, 0xce, 0x80, 0xa2, 0x9b,
	0x44, 0xea, 0x1b, 0xd8, 0x95, 0x0b, 0x51, 0x62, 0x71, 0x21, 0x48, 0xc9, 0xdc, 0xb9, 0xd6, 0xc2,
	0x34, 0x37, 0x24, 0x57, 0xa1, 0xaf, 0xe4, 0x13, 0xf8, 0x78, 0xd2, 0x76, 0x5a, 0xa4, 0x9b, 0x59,
	0x06, 0xbe, 0xef, 0xe4, 0x9c, 0x2b, 0xb2, 0xaf, 0xcc, 0x9b, 0x56, 0x02, 0x35, 0x0a, 0xc8, 0xa3,
	0x62, 0x6f, 0x6c, 0x70, 0xe4, 0x59, 0xd5, 0x96, 0xd1, 0x5b, 0x64, 0xc5, 0xe0, 0x14, 0x90, 0x7d,
	0xaf, 0x2b, 0xe9, 0x3c, 0x31, 0xa5, 0xab, 0xd1, 0xf1, 0x28, 0x27, 0x5e, 0x8e, 0xbc, 0x64, 0x70,
	0x67, 0x37, 0xb3, 0x58, 0xa0, 0xa6, 0x21, 0xab, 0x02, 0xfa, 0xda, 0xec, 0x14, 0xb7, 0x0e, 0xb7,
	0x65, 0x83, 0x21, 0x98, 0x0a, 0x87, 0xd0, 0xf5, 0xb5, 0x48, 0x72, 0xb2, 0x16, 0x81, 0x6b, 0xb2,
	0x1a, 0x3f, 0x03, 0xa6, 0xa7, 0x22, 0x46, 0x6b, 0x36, 0x3b, 0x5c, 0x46, 0xe7, 0xd1, 0xd5, 0x7f,
	0xbd, 0x7f, 0xad, 0x7f, 0x22, 0x11, 0xe7, 0x7d, 0xa1, 0xf4, 0x4d, 0x9c, 0xc0, 0x64, 0x95, 0xbe,
	0xd3, 0x7a, 0xf8, 0x28, 0xcb, 0xe4, 0xc1, 0x96, 0x72, 0xf6, 0xa1, 0x4e, 0x60, 0xd6, 0xe0, 0x51,
	0x24, 0x1f, 0x68, 0xb6, 0xe8, 0xcb, 0x80, 0xcc, 0xb5, 0xad, 0xc2, 0x72, 0xd1, 0xa7, 0x5f, 0xfe,
	0x4d, 0x1f, 0xc6, 0xc9, 0x61, 0x9c, 0x2c, 0xba, 0x71, 0x0f, 0xc3, 0x36, 0x7d, 0x3c, 0xe8, 0xcf,
	0x7b, 0xfb, 0x4e, 0x8b, 0x0b, 0xa0, 0xe6, 0x70, 0xb5, 0xa7, 0xe8, 0xf5, 0x1f, 0x83, 0xfb, 0x5e,
	0xac, 0x5e, 0x32, 0x6d, 0x5a, 0x99, 0x77, 0x68, 0x31, 0xa1, 0xf7, 0x23, 0x5a, 0x80, 0xdb, 0xc4,
	0xfd, 0x01, 0x6f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x17, 0x2c, 0x8e, 0x55, 0xcb, 0x01, 0x00,
	0x00,
}
