// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edith.proto

package edith

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Filename             string   `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_31ce212e93cb3421, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Request) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Request) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_31ce212e93cb3421, []int{1}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RequestItems struct {
	Texts                []*Request `protobuf:"bytes,1,rep,name=texts,proto3" json:"texts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestItems) Reset()         { *m = RequestItems{} }
func (m *RequestItems) String() string { return proto.CompactTextString(m) }
func (*RequestItems) ProtoMessage()    {}
func (*RequestItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_31ce212e93cb3421, []int{2}
}

func (m *RequestItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestItems.Unmarshal(m, b)
}
func (m *RequestItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestItems.Marshal(b, m, deterministic)
}
func (m *RequestItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestItems.Merge(m, src)
}
func (m *RequestItems) XXX_Size() int {
	return xxx_messageInfo_RequestItems.Size(m)
}
func (m *RequestItems) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestItems.DiscardUnknown(m)
}

var xxx_messageInfo_RequestItems proto.InternalMessageInfo

func (m *RequestItems) GetTexts() []*Request {
	if m != nil {
		return m.Texts
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*RequestItems)(nil), "RequestItems")
}

func init() {
	proto.RegisterFile("edith.proto", fileDescriptor_31ce212e93cb3421)
}

var fileDescriptor_31ce212e93cb3421 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0x69, 0xda, 0x64, 0x5a, 0x41, 0xf6, 0x20, 0x4b, 0x29, 0xda, 0x06, 0x0f, 0x3d,
	0x6d, 0xa1, 0xfe, 0x03, 0x41, 0xc5, 0xeb, 0xea, 0xc9, 0x5b, 0xd2, 0x8c, 0xed, 0x82, 0xd9, 0x8d,
	0xd9, 0x29, 0x98, 0xb3, 0xbf, 0xc1, 0xff, 0xeb, 0x66, 0x4d, 0xa3, 0x08, 0x7a, 0x7b, 0xfb, 0xde,
	0xc7, 0xec, 0x1b, 0x06, 0xc6, 0x58, 0x28, 0xda, 0x89, 0xaa, 0x36, 0x64, 0xd2, 0xf7, 0x00, 0x46,
	0x12, 0x5f, 0xf7, 0x68, 0x89, 0x9d, 0xc1, 0xd0, 0xa2, 0x2e, 0xb0, 0xe6, 0xc1, 0x3c, 0x58, 0x26,
	0xb2, 0x7b, 0xb1, 0x19, 0x24, 0x35, 0x6e, 0x54, 0xa5, 0x50, 0x13, 0x3f, 0xf6, 0xd1, 0xb7, 0xc1,
	0xa6, 0x10, 0x3f, 0xab, 0x17, 0xd4, 0x59, 0x89, 0x3c, 0xf4, 0x61, 0xff, 0x66, 0x0c, 0x06, 0xb9,
	0x29, 0x1a, 0x3e, 0x70, 0xfe, 0x44, 0x7a, 0xdd, 0x7a, 0xd4, 0x54, 0xc8, 0x23, 0xcf, 0x7a, 0x9d,
	0xce, 0x20, 0x96, 0x68, 0x2b, 0xa3, 0x2d, 0xb2, 0x53, 0x08, 0x4b, 0xbb, 0xed, 0x2a, 0xb4, 0x32,
	0x15, 0x30, 0xe9, 0x2a, 0xde, 0x13, 0x96, 0x96, 0x9d, 0x43, 0x44, 0xf8, 0x46, 0xd6, 0x31, 0xe1,
	0x72, 0xbc, 0x8e, 0x45, 0x97, 0xca, 0x2f, 0x7b, 0xfd, 0x11, 0x40, 0x74, 0xd3, 0xee, 0xc8, 0x16,
	0x10, 0x3f, 0xb8, 0x1d, 0x1e, 0x9d, 0xcd, 0x7a, 0x6c, 0x9a, 0x88, 0xc3, 0x67, 0xe9, 0xd1, 0x01,
	0xb9, 0x75, 0x95, 0xff, 0x42, 0x2e, 0x61, 0x74, 0x87, 0xf4, 0x6b, 0xc8, 0x89, 0xf8, 0xd9, 0xc9,
	0x51, 0x73, 0x4f, 0xfd, 0x33, 0xe7, 0x7a, 0xf1, 0x74, 0xb1, 0x75, 0xad, 0xf6, 0xb9, 0xd8, 0x98,
	0x72, 0xd5, 0x60, 0x6e, 0xb2, 0x9d, 0xce, 0x74, 0x66, 0x2c, 0xaa, 0x95, 0x3f, 0x4a, 0x3e, 0xf4,
	0x57, 0xb9, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x95, 0xd3, 0xf9, 0xa4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EdithClient is the client API for Edith service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EdithClient interface {
	SendText(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	SendFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetText(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RequestItems, error)
	GetFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type edithClient struct {
	cc grpc.ClientConnInterface
}

func NewEdithClient(cc grpc.ClientConnInterface) EdithClient {
	return &edithClient{cc}
}

func (c *edithClient) SendText(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Edith/SendText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edithClient) SendFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Edith/SendFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edithClient) GetText(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RequestItems, error) {
	out := new(RequestItems)
	err := c.cc.Invoke(ctx, "/Edith/GetText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edithClient) GetFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Edith/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdithServer is the server API for Edith service.
type EdithServer interface {
	SendText(context.Context, *Request) (*Response, error)
	SendFile(context.Context, *Request) (*Response, error)
	GetText(context.Context, *Request) (*RequestItems, error)
	GetFile(context.Context, *Request) (*Response, error)
}

// UnimplementedEdithServer can be embedded to have forward compatible implementations.
type UnimplementedEdithServer struct {
}

func (*UnimplementedEdithServer) SendText(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendText not implemented")
}
func (*UnimplementedEdithServer) SendFile(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFile not implemented")
}
func (*UnimplementedEdithServer) GetText(ctx context.Context, req *Request) (*RequestItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetText not implemented")
}
func (*UnimplementedEdithServer) GetFile(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}

func RegisterEdithServer(s *grpc.Server, srv EdithServer) {
	s.RegisterService(&_Edith_serviceDesc, srv)
}

func _Edith_SendText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdithServer).SendText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Edith/SendText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdithServer).SendText(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Edith_SendFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdithServer).SendFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Edith/SendFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdithServer).SendFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Edith_GetText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdithServer).GetText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Edith/GetText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdithServer).GetText(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Edith_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdithServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Edith/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdithServer).GetFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Edith_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Edith",
	HandlerType: (*EdithServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendText",
			Handler:    _Edith_SendText_Handler,
		},
		{
			MethodName: "SendFile",
			Handler:    _Edith_SendFile_Handler,
		},
		{
			MethodName: "GetText",
			Handler:    _Edith_GetText_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _Edith_GetFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edith.proto",
}
