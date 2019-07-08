// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

package hello

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

type Reqeust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reqeust) Reset()         { *m = Reqeust{} }
func (m *Reqeust) String() string { return proto.CompactTextString(m) }
func (*Reqeust) ProtoMessage()    {}
func (*Reqeust) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7116fbf19896b02, []int{0}
}

func (m *Reqeust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reqeust.Unmarshal(m, b)
}
func (m *Reqeust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reqeust.Marshal(b, m, deterministic)
}
func (m *Reqeust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reqeust.Merge(m, src)
}
func (m *Reqeust) XXX_Size() int {
	return xxx_messageInfo_Reqeust.Size(m)
}
func (m *Reqeust) XXX_DiscardUnknown() {
	xxx_messageInfo_Reqeust.DiscardUnknown(m)
}

var xxx_messageInfo_Reqeust proto.InternalMessageInfo

type Response struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7116fbf19896b02, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func init() {
	proto.RegisterType((*Reqeust)(nil), "hello.Reqeust")
	proto.RegisterType((*Response)(nil), "hello.Response")
}

func init() { proto.RegisterFile("proto/hello.proto", fileDescriptor_f7116fbf19896b02) }

var fileDescriptor_f7116fbf19896b02 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x1c, 0x25, 0x4e,
	0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd4, 0xd2, 0xe2, 0x12, 0x25, 0x19, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x2e, 0xe6, 0xe2, 0x92, 0x22, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x10, 0xd3, 0xc8, 0x90, 0x8b, 0xd5, 0x03, 0xa4, 0x43, 0x48, 0x83, 0x8b, 0x39,
	0x38, 0xb1, 0x52, 0x88, 0x4f, 0x0f, 0x62, 0x1a, 0x54, 0xb7, 0x14, 0x3f, 0x9c, 0x0f, 0x31, 0x42,
	0x89, 0x21, 0x89, 0x0d, 0x6c, 0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x63, 0x39, 0xc7, 0xb0,
	0x7e, 0x00, 0x00, 0x00,
}
