// Code generated by protoc-gen-go. DO NOT EDIT.
// source: global.proto

package pb

import (
	fmt "fmt"
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

type Id struct {
	// The name of the feature.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_4baa8fc7dedf329e, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Id) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func init() {
	proto.RegisterType((*Id)(nil), "pb.Id")
}

func init() { proto.RegisterFile("global.proto", fileDescriptor_4baa8fc7dedf329e) }

var fileDescriptor_4baa8fc7dedf329e = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe3, 0x62,
	0xf2, 0x4c, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0xca, 0x4c, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x2a, 0x2f, 0x91, 0x60, 0x02, 0x0b, 0x80, 0x98, 0x4e,
	0x86, 0x5c, 0xd2, 0x99, 0xf9, 0x7a, 0xe9, 0x45, 0x05, 0xc9, 0x7a, 0xa9, 0x15, 0x89, 0xb9, 0x05,
	0x39, 0xa9, 0xc5, 0x7a, 0x45, 0xf9, 0xa5, 0x25, 0xa9, 0xe9, 0xa5, 0x99, 0x29, 0xa9, 0x4e, 0x02,
	0xa1, 0xc5, 0xa9, 0x45, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x01, 0x20, 0xc3, 0x03, 0x18,
	0x93, 0xd8, 0xc0, 0xb6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x91, 0xeb, 0x9d, 0x54, 0x75,
	0x00, 0x00, 0x00,
}
