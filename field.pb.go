// Code generated by protoc-gen-go. DO NOT EDIT.
// source: field.proto

package tyrellcorp

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

type Field struct {
	// the name of the field
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the description of what the field is - will be output as a comment
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// the data type of the field
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// this is the "field number" used for protobufs
	Sequence int32 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// is this field a list
	IsList bool `protobuf:"varint,5,opt,name=is_list,json=isList,proto3" json:"is_list,omitempty"`
	// is this a field used in identifying uniqueness
	IsKey                bool     `protobuf:"varint,6,opt,name=is_key,json=isKey,proto3" json:"is_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_04234ff7fdd53e6e, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Field) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Field) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Field) GetIsList() bool {
	if m != nil {
		return m.IsList
	}
	return false
}

func (m *Field) GetIsKey() bool {
	if m != nil {
		return m.IsKey
	}
	return false
}

func init() {
	proto.RegisterType((*Field)(nil), "tyrellcorp.Field")
}

func init() { proto.RegisterFile("field.proto", fileDescriptor_04234ff7fdd53e6e) }

var fileDescriptor_04234ff7fdd53e6e = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x68, 0x42, 0xb9, 0x6e, 0x96, 0x10, 0x56, 0xa7, 0x88, 0x29, 0x03, 0x4a, 0x07,
	0x46, 0xb6, 0x0e, 0x2c, 0x30, 0x54, 0x99, 0x10, 0x4b, 0x45, 0xdd, 0x03, 0x4e, 0x75, 0x72, 0x26,
	0x77, 0x1d, 0xf2, 0x63, 0xf8, 0xaf, 0xc8, 0x46, 0x22, 0xd9, 0x3e, 0xbf, 0xa7, 0x4f, 0x7a, 0x3e,
	0x58, 0x7d, 0x10, 0x86, 0x63, 0x13, 0x07, 0x56, 0xb6, 0xa0, 0xe3, 0x80, 0x21, 0x78, 0x1e, 0xe2,
	0xdd, 0x8f, 0x81, 0xe2, 0x29, 0x75, 0xd6, 0xc2, 0xa2, 0x7f, 0xef, 0xd0, 0x99, 0xca, 0xd4, 0xd7,
	0x6d, 0x66, 0x5b, 0xc1, 0xea, 0x88, 0xe2, 0x07, 0x8a, 0x4a, 0xdc, 0xbb, 0x8b, 0x5c, 0xcd, 0xa3,
	0x64, 0xe9, 0x18, 0xd1, 0x5d, 0xfe, 0x59, 0x89, 0xed, 0x1a, 0x96, 0x82, 0xdf, 0x67, 0xec, 0x3d,
	0xba, 0x45, 0x65, 0xea, 0xa2, 0xfd, 0x7f, 0xdb, 0x5b, 0xb8, 0x22, 0xd9, 0x07, 0x12, 0x75, 0x45,
	0x65, 0xea, 0x65, 0x5b, 0x92, 0xbc, 0x90, 0xa8, 0xbd, 0x81, 0x92, 0x64, 0x7f, 0xc2, 0xd1, 0x95,
	0x39, 0x2f, 0x48, 0x9e, 0x71, 0xdc, 0xbe, 0xc2, 0xda, 0x73, 0xd7, 0x60, 0x08, 0xc4, 0xaa, 0x91,
	0xc3, 0xa9, 0x99, 0xd6, 0x6f, 0x21, 0x4f, 0xdf, 0xa5, 0x5f, 0xed, 0xcc, 0xdb, 0xfd, 0x27, 0xe9,
	0xd7, 0xf9, 0xd0, 0x78, 0xee, 0x36, 0x33, 0x61, 0x33, 0x09, 0x8f, 0x13, 0x1e, 0xca, 0x7c, 0x8c,
	0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x6b, 0x53, 0xe6, 0x1b, 0x01, 0x00, 0x00,
}
