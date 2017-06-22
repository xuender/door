// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

/*
Package door is a generated protocol buffer package.

It is generated from these files:
	event.proto

It has these top-level messages:
	Event
*/
package door

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MethodEnum int32

const (
	MethodEnum_GET    MethodEnum = 0
	MethodEnum_POST   MethodEnum = 1
	MethodEnum_PUT    MethodEnum = 2
	MethodEnum_DELETE MethodEnum = 3
)

var MethodEnum_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "PUT",
	3: "DELETE",
}
var MethodEnum_value = map[string]int32{
	"GET":    0,
	"POST":   1,
	"PUT":    2,
	"DELETE": 3,
}

func (x MethodEnum) String() string {
	return proto.EnumName(MethodEnum_name, int32(x))
}
func (MethodEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event struct {
	Path   string     `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Method MethodEnum `protobuf:"varint,2,opt,name=method,enum=door.MethodEnum" json:"method,omitempty"`
	Data   []byte     `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Event) GetMethod() MethodEnum {
	if m != nil {
		return m.Method
	}
	return MethodEnum_GET
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "door.Event")
	proto.RegisterEnum("door.MethodEnum", MethodEnum_name, MethodEnum_value)
}

func init() { proto.RegisterFile("event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xc9, 0xcf, 0x2f, 0x52, 0x8a, 0xe4,
	0x62, 0x75, 0x05, 0x09, 0x0a, 0x09, 0x71, 0xb1, 0x14, 0x24, 0x96, 0x64, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x1a, 0x5c, 0x6c, 0xb9, 0xa9, 0x25, 0x19, 0xf9, 0x29, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x02, 0x7a, 0x20, 0x3d, 0x7a, 0xbe, 0x60, 0x31, 0xd7, 0xbc,
	0xd2, 0xdc, 0x20, 0xa8, 0x3c, 0x48, 0x77, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0x4f, 0x10, 0x98, 0xad, 0x65, 0xc2, 0xc5, 0x85, 0x50, 0x29, 0xc4, 0xce, 0xc5, 0xec, 0xee, 0x1a,
	0x22, 0xc0, 0x20, 0xc4, 0xc1, 0xc5, 0x12, 0xe0, 0x1f, 0x1c, 0x22, 0xc0, 0x08, 0x12, 0x0a, 0x08,
	0x0d, 0x11, 0x60, 0x12, 0xe2, 0xe2, 0x62, 0x73, 0x71, 0xf5, 0x71, 0x0d, 0x71, 0x15, 0x60, 0x4e,
	0x62, 0x03, 0xbb, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x31, 0x44, 0x9a, 0xe7, 0xac, 0x00,
	0x00, 0x00,
}