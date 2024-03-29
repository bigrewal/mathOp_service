// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

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

// Add request message type
type AddRequest struct {
	X                    int64    `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int64    `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *AddRequest) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type AddResponse struct {
	Answer               int64    `protobuf:"varint,1,opt,name=Answer,proto3" json:"Answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetAnswer() int64 {
	if m != nil {
		return m.Answer
	}
	return 0
}

// Sqaure Root message types
type SqrtRequest struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=X,proto3" json:"X,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtRequest) Reset()         { *m = SqrtRequest{} }
func (m *SqrtRequest) String() string { return proto.CompactTextString(m) }
func (*SqrtRequest) ProtoMessage()    {}
func (*SqrtRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *SqrtRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtRequest.Unmarshal(m, b)
}
func (m *SqrtRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtRequest.Marshal(b, m, deterministic)
}
func (m *SqrtRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtRequest.Merge(m, src)
}
func (m *SqrtRequest) XXX_Size() int {
	return xxx_messageInfo_SqrtRequest.Size(m)
}
func (m *SqrtRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtRequest proto.InternalMessageInfo

func (m *SqrtRequest) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

type SqrtResponse struct {
	Answer               float64  `protobuf:"fixed64,1,opt,name=Answer,proto3" json:"Answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtResponse) Reset()         { *m = SqrtResponse{} }
func (m *SqrtResponse) String() string { return proto.CompactTextString(m) }
func (*SqrtResponse) ProtoMessage()    {}
func (*SqrtResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *SqrtResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtResponse.Unmarshal(m, b)
}
func (m *SqrtResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtResponse.Marshal(b, m, deterministic)
}
func (m *SqrtResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtResponse.Merge(m, src)
}
func (m *SqrtResponse) XXX_Size() int {
	return xxx_messageInfo_SqrtResponse.Size(m)
}
func (m *SqrtResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtResponse proto.InternalMessageInfo

func (m *SqrtResponse) GetAnswer() float64 {
	if m != nil {
		return m.Answer
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "service.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "service.AddResponse")
	proto.RegisterType((*SqrtRequest)(nil), "service.SqrtRequest")
	proto.RegisterType((*SqrtResponse)(nil), "service.SqrtResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x34, 0xb8,
	0xb8, 0x1c, 0x53, 0x52, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x78, 0xb8, 0x18, 0x23,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x18, 0x23, 0x40, 0xbc, 0x48, 0x09, 0x26, 0x08, 0x2f,
	0x52, 0x49, 0x95, 0x8b, 0x1b, 0xac, 0xb2, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b,
	0xcd, 0x31, 0xaf, 0xb8, 0x3c, 0xb5, 0x08, 0xaa, 0x1e, 0xca, 0x53, 0x92, 0xe6, 0xe2, 0x0e, 0x2e,
	0x2c, 0x2a, 0xc1, 0x30, 0x91, 0x31, 0x88, 0x31, 0x42, 0x49, 0x8d, 0x8b, 0x07, 0x22, 0x89, 0xd5,
	0x10, 0x46, 0x98, 0x21, 0x46, 0x65, 0x5c, 0xbc, 0xbe, 0x89, 0x25, 0x19, 0xfe, 0x05, 0xc1, 0x10,
	0x67, 0x0a, 0x19, 0x70, 0x31, 0x3b, 0xa6, 0xa4, 0x08, 0x09, 0xeb, 0xc1, 0xbc, 0x81, 0x70, 0xb4,
	0x94, 0x08, 0xaa, 0x20, 0xd4, 0x68, 0x63, 0x2e, 0x16, 0x90, 0x55, 0x42, 0x08, 0x59, 0x24, 0x67,
	0x49, 0x89, 0xa2, 0x89, 0x42, 0x34, 0x25, 0xb1, 0x81, 0x43, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x71, 0xad, 0xd8, 0x4e, 0x2e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MathOpServiceClient is the client API for MathOpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathOpServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type mathOpServiceClient struct {
	cc *grpc.ClientConn
}

func NewMathOpServiceClient(cc *grpc.ClientConn) MathOpServiceClient {
	return &mathOpServiceClient{cc}
}

func (c *mathOpServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/service.MathOpService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathOpServiceClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/service.MathOpService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathOpServiceServer is the server API for MathOpService service.
type MathOpServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
}

// UnimplementedMathOpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMathOpServiceServer struct {
}

func (*UnimplementedMathOpServiceServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedMathOpServiceServer) Sqrt(ctx context.Context, req *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}

func RegisterMathOpServiceServer(s *grpc.Server, srv MathOpServiceServer) {
	s.RegisterService(&_MathOpService_serviceDesc, srv)
}

func _MathOpService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathOpServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MathOpService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathOpServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathOpService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathOpServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MathOpService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathOpServiceServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathOpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.MathOpService",
	HandlerType: (*MathOpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MathOpService_Add_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _MathOpService_Sqrt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
