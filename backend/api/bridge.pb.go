// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bridge.proto

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

type Session struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3ed31acb30cd14, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Target struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3ed31acb30cd14, []int{1}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type Payload struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3ed31acb30cd14, []int{2}
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

func (m *Payload) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *Payload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Session)(nil), "api.Session")
	proto.RegisterType((*Target)(nil), "api.Target")
	proto.RegisterType((*Payload)(nil), "api.Payload")
}

func init() { proto.RegisterFile("bridge.proto", fileDescriptor_1d3ed31acb30cd14) }

var fileDescriptor_1d3ed31acb30cd14 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2a, 0xca, 0x4c,
	0x49, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x92, 0xe4,
	0x62, 0x0f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x13, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xf2, 0x74, 0x51, 0x52, 0xe1, 0x62, 0x0b, 0x49, 0x2c, 0x4a,
	0x4f, 0x2d, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0xc8, 0x2f, 0x2e, 0xc9, 0x4b, 0xcc, 0x4d, 0x85, 0xca,
	0xc3, 0xf9, 0x4a, 0xd6, 0x5c, 0xec, 0x01, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x42, 0x32, 0x5c,
	0x9c, 0x50, 0xb3, 0xe0, 0xe6, 0x20, 0x04, 0x84, 0x84, 0xb8, 0x58, 0x9c, 0xf2, 0x53, 0x2a, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0xa3, 0x58, 0x2e, 0x36, 0x27, 0xb0, 0x93, 0x84,
	0x34, 0xb9, 0x58, 0x9d, 0x8b, 0xf2, 0x8b, 0x8b, 0x85, 0x78, 0xf4, 0x12, 0x0b, 0x32, 0xf5, 0xa0,
	0x46, 0x4a, 0xa1, 0xf0, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x85, 0x34, 0xb8, 0x38, 0x5d, 0x8b,
	0x4b, 0x12, 0x93, 0x72, 0x32, 0x8b, 0x33, 0x84, 0xb8, 0xc1, 0x0a, 0x20, 0xee, 0x84, 0xaa, 0x86,
	0x5a, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x74, 0xf3,
	0x20, 0x6a, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BridgeClient is the client API for Bridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BridgeClient interface {
	Cross(ctx context.Context, opts ...grpc.CallOption) (Bridge_CrossClient, error)
	Establish(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Session, error)
}

type bridgeClient struct {
	cc *grpc.ClientConn
}

func NewBridgeClient(cc *grpc.ClientConn) BridgeClient {
	return &bridgeClient{cc}
}

func (c *bridgeClient) Cross(ctx context.Context, opts ...grpc.CallOption) (Bridge_CrossClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Bridge_serviceDesc.Streams[0], "/api.Bridge/Cross", opts...)
	if err != nil {
		return nil, err
	}
	x := &bridgeCrossClient{stream}
	return x, nil
}

type Bridge_CrossClient interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ClientStream
}

type bridgeCrossClient struct {
	grpc.ClientStream
}

func (x *bridgeCrossClient) Send(m *Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bridgeCrossClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bridgeClient) Establish(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/api.Bridge/Establish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeServer is the server API for Bridge service.
type BridgeServer interface {
	Cross(Bridge_CrossServer) error
	Establish(context.Context, *Target) (*Session, error)
}

// UnimplementedBridgeServer can be embedded to have forward compatible implementations.
type UnimplementedBridgeServer struct {
}

func (*UnimplementedBridgeServer) Cross(srv Bridge_CrossServer) error {
	return status.Errorf(codes.Unimplemented, "method Cross not implemented")
}
func (*UnimplementedBridgeServer) Establish(ctx context.Context, req *Target) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Establish not implemented")
}

func RegisterBridgeServer(s *grpc.Server, srv BridgeServer) {
	s.RegisterService(&_Bridge_serviceDesc, srv)
}

func _Bridge_Cross_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BridgeServer).Cross(&bridgeCrossServer{stream})
}

type Bridge_CrossServer interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ServerStream
}

type bridgeCrossServer struct {
	grpc.ServerStream
}

func (x *bridgeCrossServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bridgeCrossServer) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Bridge_Establish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).Establish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Bridge/Establish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).Establish(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bridge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Bridge",
	HandlerType: (*BridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Establish",
			Handler:    _Bridge_Establish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Cross",
			Handler:       _Bridge_Cross_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bridge.proto",
}
