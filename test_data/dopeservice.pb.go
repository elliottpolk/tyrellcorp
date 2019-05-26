// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dopeservice.proto

package super

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	// unique identifier of the original incoming request to help troubleshoot
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6d99603c0f220f, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type DopeRequest struct {
	// unique identifier to help troubleshoot each request
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// username of the one making the request
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// unique identifier of the super.Dope
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// dataset to process
	Payload              []*Dope  `protobuf:"bytes,4,rep,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DopeRequest) Reset()         { *m = DopeRequest{} }
func (m *DopeRequest) String() string { return proto.CompactTextString(m) }
func (*DopeRequest) ProtoMessage()    {}
func (*DopeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6d99603c0f220f, []int{1}
}

func (m *DopeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DopeRequest.Unmarshal(m, b)
}
func (m *DopeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DopeRequest.Marshal(b, m, deterministic)
}
func (m *DopeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DopeRequest.Merge(m, src)
}
func (m *DopeRequest) XXX_Size() int {
	return xxx_messageInfo_DopeRequest.Size(m)
}
func (m *DopeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DopeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DopeRequest proto.InternalMessageInfo

func (m *DopeRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *DopeRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *DopeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DopeRequest) GetPayload() []*Dope {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DopeResponse struct {
	// unique identifier of the original incoming request to help troubleshoot
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Payload              []*Dope  `protobuf:"bytes,2,rep,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DopeResponse) Reset()         { *m = DopeResponse{} }
func (m *DopeResponse) String() string { return proto.CompactTextString(m) }
func (*DopeResponse) ProtoMessage()    {}
func (*DopeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6d99603c0f220f, []int{2}
}

func (m *DopeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DopeResponse.Unmarshal(m, b)
}
func (m *DopeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DopeResponse.Marshal(b, m, deterministic)
}
func (m *DopeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DopeResponse.Merge(m, src)
}
func (m *DopeResponse) XXX_Size() int {
	return xxx_messageInfo_DopeResponse.Size(m)
}
func (m *DopeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DopeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DopeResponse proto.InternalMessageInfo

func (m *DopeResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *DopeResponse) GetPayload() []*Dope {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "super.Empty")
	proto.RegisterType((*DopeRequest)(nil), "super.DopeRequest")
	proto.RegisterType((*DopeResponse)(nil), "super.DopeResponse")
}

func init() { proto.RegisterFile("dopeservice.proto", fileDescriptor_fb6d99603c0f220f) }

var fileDescriptor_fb6d99603c0f220f = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x4a, 0xe3, 0x40,
	0x18, 0x85, 0x49, 0xba, 0xcd, 0xb6, 0xd3, 0xee, 0xb2, 0x3b, 0x5b, 0xd8, 0x10, 0x5a, 0x28, 0x81,
	0x4a, 0xa9, 0x90, 0x60, 0xbd, 0xf3, 0xb2, 0x46, 0xd0, 0xbb, 0x52, 0xf5, 0xa6, 0x37, 0x9a, 0x66,
	0x7e, 0xea, 0xe0, 0x24, 0x33, 0x4e, 0x26, 0x85, 0x22, 0x22, 0xf8, 0x0a, 0x3e, 0x97, 0x57, 0xbe,
	0x82, 0x0f, 0x22, 0x9d, 0xa9, 0xd2, 0x62, 0x4b, 0xf1, 0x32, 0xe7, 0x9f, 0xf9, 0x72, 0xce, 0x99,
	0x1f, 0xfd, 0x25, 0x5c, 0x40, 0x0e, 0x72, 0x46, 0x13, 0x08, 0x84, 0xe4, 0x8a, 0xe3, 0x72, 0x5e,
	0x08, 0x90, 0x1e, 0x5a, 0x4c, 0x8c, 0xe4, 0x35, 0xa7, 0x9c, 0x4f, 0x19, 0x84, 0xb1, 0xa0, 0x61,
	0x9c, 0x65, 0x5c, 0xc5, 0x8a, 0xf2, 0x2c, 0x37, 0x53, 0x7f, 0x0f, 0x95, 0x4f, 0x52, 0xa1, 0xe6,
	0xb8, 0x85, 0x90, 0x84, 0xbb, 0x02, 0x72, 0x75, 0x45, 0x89, 0x6b, 0xb5, 0xad, 0x6e, 0x75, 0x54,
	0x5d, 0x2a, 0x67, 0xc4, 0x7f, 0x44, 0xb5, 0x88, 0x0b, 0x18, 0x19, 0x61, 0xc7, 0x69, 0xec, 0xa1,
	0x4a, 0x91, 0x83, 0xcc, 0xe2, 0x14, 0x5c, 0x5b, 0x0f, 0x3f, 0xbf, 0xf1, 0x6f, 0x64, 0x53, 0xe2,
	0x96, 0xb4, 0x6a, 0x53, 0x82, 0x3b, 0xe8, 0xa7, 0x88, 0xe7, 0x8c, 0xc7, 0xc4, 0xfd, 0xd1, 0x2e,
	0x75, 0x6b, 0xfd, 0x5a, 0xa0, 0x43, 0x04, 0xfa, 0x7f, 0x1f, 0x33, 0xff, 0x02, 0xd5, 0x8d, 0x81,
	0x5c, 0xf0, 0x2c, 0x87, 0x5d, 0x0e, 0x56, 0xa8, 0xf6, 0x76, 0x6a, 0xff, 0xc5, 0x36, 0xb9, 0xce,
	0x4d, 0x8b, 0x38, 0x42, 0xce, 0xb1, 0x84, 0x58, 0x01, 0xc6, 0xab, 0xe7, 0x0d, 0xd6, 0xab, 0x2f,
	0x35, 0xdd, 0x98, 0xef, 0x3e, 0xbd, 0xbe, 0x3d, 0xdb, 0xd8, 0xff, 0xa5, 0xab, 0x9d, 0x1d, 0x84,
	0xfa, 0x39, 0x8e, 0xac, 0x1e, 0xbe, 0x46, 0x95, 0x11, 0x28, 0x49, 0x61, 0xb6, 0x99, 0xf3, 0x6f,
	0x4d, 0x33, 0x81, 0xfc, 0x7d, 0x8d, 0xeb, 0xe0, 0x75, 0xdc, 0xb8, 0x81, 0xf1, 0x9a, 0x10, 0xde,
	0x53, 0xf2, 0x80, 0x87, 0xc8, 0xb9, 0x14, 0x64, 0x9b, 0xcf, 0x8d, 0xfc, 0x96, 0xe6, 0xff, 0xf7,
	0x36, 0xe0, 0x16, 0x9e, 0x23, 0xe4, 0x44, 0xc0, 0xe0, 0x3b, 0xc9, 0x7b, 0x5f, 0x92, 0x0f, 0x4e,
	0x51, 0x23, 0xe1, 0x69, 0xa0, 0xe6, 0x12, 0x18, 0x4b, 0xb8, 0x14, 0xe6, 0xde, 0xe0, 0xcf, 0x4a,
	0xc9, 0xc3, 0xc5, 0xe2, 0x0d, 0xad, 0x71, 0x73, 0x4a, 0xd5, 0x4d, 0x31, 0x09, 0x12, 0x9e, 0x86,
	0xc0, 0x18, 0xe5, 0x4a, 0x09, 0xce, 0x6e, 0x43, 0x7d, 0x63, 0xe2, 0xe8, 0xfd, 0x3c, 0x7c, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x94, 0xa9, 0x32, 0x4c, 0xe5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DopeServiceClient is the client API for DopeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DopeServiceClient interface {
	// create new Dope item(s)
	Create(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*Empty, error)
	// retrieve a list of Dope items
	Retrieve(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*DopeResponse, error)
	// update Dope item(s)
	Update(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*DopeResponse, error)
	// delete Dope item(s)
	Delete(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*Empty, error)
}

type dopeServiceClient struct {
	cc *grpc.ClientConn
}

func NewDopeServiceClient(cc *grpc.ClientConn) DopeServiceClient {
	return &dopeServiceClient{cc}
}

func (c *dopeServiceClient) Create(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/super.DopeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopeServiceClient) Retrieve(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*DopeResponse, error) {
	out := new(DopeResponse)
	err := c.cc.Invoke(ctx, "/super.DopeService/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopeServiceClient) Update(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*DopeResponse, error) {
	out := new(DopeResponse)
	err := c.cc.Invoke(ctx, "/super.DopeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopeServiceClient) Delete(ctx context.Context, in *DopeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/super.DopeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DopeServiceServer is the server API for DopeService service.
type DopeServiceServer interface {
	// create new Dope item(s)
	Create(context.Context, *DopeRequest) (*Empty, error)
	// retrieve a list of Dope items
	Retrieve(context.Context, *DopeRequest) (*DopeResponse, error)
	// update Dope item(s)
	Update(context.Context, *DopeRequest) (*DopeResponse, error)
	// delete Dope item(s)
	Delete(context.Context, *DopeRequest) (*Empty, error)
}

// UnimplementedDopeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDopeServiceServer struct {
}

func (*UnimplementedDopeServiceServer) Create(ctx context.Context, req *DopeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDopeServiceServer) Retrieve(ctx context.Context, req *DopeRequest) (*DopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (*UnimplementedDopeServiceServer) Update(ctx context.Context, req *DopeRequest) (*DopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDopeServiceServer) Delete(ctx context.Context, req *DopeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterDopeServiceServer(s *grpc.Server, srv DopeServiceServer) {
	s.RegisterService(&_DopeService_serviceDesc, srv)
}

func _DopeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/super.DopeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopeServiceServer).Create(ctx, req.(*DopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopeService_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopeServiceServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/super.DopeService/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopeServiceServer).Retrieve(ctx, req.(*DopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/super.DopeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopeServiceServer).Update(ctx, req.(*DopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/super.DopeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopeServiceServer).Delete(ctx, req.(*DopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DopeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "super.DopeService",
	HandlerType: (*DopeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DopeService_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _DopeService_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DopeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DopeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dopeservice.proto",
}