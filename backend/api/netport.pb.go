// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netport.proto

package api

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

type NewPortRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPortRequest) Reset()         { *m = NewPortRequest{} }
func (m *NewPortRequest) String() string { return proto.CompactTextString(m) }
func (*NewPortRequest) ProtoMessage()    {}
func (*NewPortRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c23dc8bc2c4b258, []int{0}
}

func (m *NewPortRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPortRequest.Unmarshal(m, b)
}
func (m *NewPortRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPortRequest.Marshal(b, m, deterministic)
}
func (m *NewPortRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPortRequest.Merge(m, src)
}
func (m *NewPortRequest) XXX_Size() int {
	return xxx_messageInfo_NewPortRequest.Size(m)
}
func (m *NewPortRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPortRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewPortRequest proto.InternalMessageInfo

type NewPortResponse struct {
	PortID               []byte   `protobuf:"bytes,1,opt,name=PortID,proto3" json:"PortID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPortResponse) Reset()         { *m = NewPortResponse{} }
func (m *NewPortResponse) String() string { return proto.CompactTextString(m) }
func (*NewPortResponse) ProtoMessage()    {}
func (*NewPortResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c23dc8bc2c4b258, []int{1}
}

func (m *NewPortResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPortResponse.Unmarshal(m, b)
}
func (m *NewPortResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPortResponse.Marshal(b, m, deterministic)
}
func (m *NewPortResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPortResponse.Merge(m, src)
}
func (m *NewPortResponse) XXX_Size() int {
	return xxx_messageInfo_NewPortResponse.Size(m)
}
func (m *NewPortResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPortResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewPortResponse proto.InternalMessageInfo

func (m *NewPortResponse) GetPortID() []byte {
	if m != nil {
		return m.PortID
	}
	return nil
}

type PopRequest struct {
	PortID               []byte   `protobuf:"bytes,1,opt,name=PortID,proto3" json:"PortID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopRequest) Reset()         { *m = PopRequest{} }
func (m *PopRequest) String() string { return proto.CompactTextString(m) }
func (*PopRequest) ProtoMessage()    {}
func (*PopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c23dc8bc2c4b258, []int{2}
}

func (m *PopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopRequest.Unmarshal(m, b)
}
func (m *PopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopRequest.Marshal(b, m, deterministic)
}
func (m *PopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopRequest.Merge(m, src)
}
func (m *PopRequest) XXX_Size() int {
	return xxx_messageInfo_PopRequest.Size(m)
}
func (m *PopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PopRequest proto.InternalMessageInfo

func (m *PopRequest) GetPortID() []byte {
	if m != nil {
		return m.PortID
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c23dc8bc2c4b258, []int{3}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type Payload struct {
	Source               []byte   `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Dest                 []byte   `protobuf:"bytes,2,opt,name=Dest,proto3" json:"Dest,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c23dc8bc2c4b258, []int{4}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetSource() []byte {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Payload) GetDest() []byte {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *Payload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*NewPortRequest)(nil), "api.NewPortRequest")
	proto.RegisterType((*NewPortResponse)(nil), "api.NewPortResponse")
	proto.RegisterType((*PopRequest)(nil), "api.PopRequest")
	proto.RegisterType((*Void)(nil), "api.Void")
	proto.RegisterType((*Payload)(nil), "api.Payload")
}

func init() { proto.RegisterFile("netport.proto", fileDescriptor_6c23dc8bc2c4b258) }

var fileDescriptor_6c23dc8bc2c4b258 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc4, 0x40,
	0x10, 0x84, 0x13, 0x13, 0xb2, 0xd8, 0xac, 0xae, 0xb4, 0x22, 0x21, 0x27, 0x1d, 0x44, 0xf4, 0x92,
	0x83, 0x82, 0x0f, 0x20, 0x7b, 0xd9, 0x8b, 0x0c, 0x11, 0xbc, 0x8f, 0xa6, 0xc1, 0x80, 0x6c, 0x8f,
	0xf3, 0x83, 0xec, 0x33, 0xf8, 0xd2, 0x32, 0x3f, 0x2e, 0xe4, 0xb0, 0xb7, 0xae, 0x6f, 0xaa, 0xa0,
	0x6a, 0xe0, 0x64, 0x4b, 0x4e, 0xb3, 0x71, 0xbd, 0x36, 0xec, 0x18, 0x2b, 0xa5, 0x27, 0x71, 0x06,
	0xa7, 0x2f, 0xf4, 0x23, 0xd9, 0xb8, 0x81, 0xbe, 0x3d, 0x59, 0x27, 0xee, 0x61, 0xb5, 0x27, 0x56,
	0xf3, 0xd6, 0x12, 0x5e, 0x42, 0x13, 0xf4, 0x66, 0xdd, 0x96, 0x57, 0xe5, 0xdd, 0x72, 0xc8, 0x4a,
	0xdc, 0x00, 0x48, 0xd6, 0x39, 0x78, 0xd0, 0xd5, 0x40, 0xfd, 0xc6, 0xd3, 0x28, 0x36, 0xb0, 0x90,
	0x6a, 0xf7, 0xc5, 0x6a, 0x0c, 0xd6, 0x57, 0xf6, 0xe6, 0x83, 0xfe, 0xad, 0x49, 0x21, 0x42, 0xbd,
	0x26, 0xeb, 0xda, 0xa3, 0x48, 0xe3, 0x1d, 0xd8, 0x33, 0x8f, 0xbb, 0xb6, 0x4a, 0x2c, 0xdc, 0x0f,
	0xbf, 0x25, 0x34, 0x03, 0x7b, 0x47, 0x06, 0x9f, 0x60, 0x91, 0xeb, 0xe2, 0x79, 0xaf, 0xf4, 0xd4,
	0xcf, 0xe7, 0x74, 0x17, 0x73, 0x98, 0x16, 0x89, 0x02, 0xaf, 0xa1, 0x96, 0xde, 0x7e, 0xe2, 0x32,
	0xbe, 0xe7, 0x62, 0xdd, 0x71, 0x54, 0xb1, 0x6e, 0x81, 0xb7, 0x50, 0x49, 0xd6, 0xb8, 0x4a, 0x8e,
	0xfd, 0xd0, 0x6e, 0x16, 0x11, 0xc5, 0x7b, 0x13, 0xff, 0xf3, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xcc, 0x6b, 0x9d, 0x08, 0x60, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	NewPort(ctx context.Context, in *NewPortRequest, opts ...grpc.CallOption) (*NewPortResponse, error)
	Push(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Void, error)
	Pop(ctx context.Context, in *PopRequest, opts ...grpc.CallOption) (*Payload, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) NewPort(ctx context.Context, in *NewPortRequest, opts ...grpc.CallOption) (*NewPortResponse, error) {
	out := new(NewPortResponse)
	err := c.cc.Invoke(ctx, "/api.Router/NewPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Push(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/api.Router/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Pop(ctx context.Context, in *PopRequest, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := c.cc.Invoke(ctx, "/api.Router/Pop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	NewPort(context.Context, *NewPortRequest) (*NewPortResponse, error)
	Push(context.Context, *Payload) (*Void, error)
	Pop(context.Context, *PopRequest) (*Payload, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) NewPort(ctx context.Context, req *NewPortRequest) (*NewPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewPort not implemented")
}
func (*UnimplementedRouterServer) Push(ctx context.Context, req *Payload) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (*UnimplementedRouterServer) Pop(ctx context.Context, req *PopRequest) (*Payload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pop not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_NewPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).NewPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Router/NewPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).NewPort(ctx, req.(*NewPortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Router/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Push(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Pop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Pop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Router/Pop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Pop(ctx, req.(*PopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPort",
			Handler:    _Router_NewPort_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Router_Push_Handler,
		},
		{
			MethodName: "Pop",
			Handler:    _Router_Pop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "netport.proto",
}