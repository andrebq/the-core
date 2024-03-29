// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

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

type ContentID struct {
	Alg                  string   `protobuf:"bytes,1,opt,name=Alg,proto3" json:"Alg,omitempty"`
	HexHash              string   `protobuf:"bytes,2,opt,name=HexHash,proto3" json:"HexHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentID) Reset()         { *m = ContentID{} }
func (m *ContentID) String() string { return proto.CompactTextString(m) }
func (*ContentID) ProtoMessage()    {}
func (*ContentID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *ContentID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentID.Unmarshal(m, b)
}
func (m *ContentID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentID.Marshal(b, m, deterministic)
}
func (m *ContentID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentID.Merge(m, src)
}
func (m *ContentID) XXX_Size() int {
	return xxx_messageInfo_ContentID.Size(m)
}
func (m *ContentID) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentID.DiscardUnknown(m)
}

var xxx_messageInfo_ContentID proto.InternalMessageInfo

func (m *ContentID) GetAlg() string {
	if m != nil {
		return m.Alg
	}
	return ""
}

func (m *ContentID) GetHexHash() string {
	if m != nil {
		return m.HexHash
	}
	return ""
}

type Content struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
	Tags                 []*Tag   `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Content) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Tag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ContentID)(nil), "api.ContentID")
	proto.RegisterType((*Content)(nil), "api.Content")
	proto.RegisterType((*Tag)(nil), "api.Tag")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x32,
	0xe7, 0xe2, 0x74, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0xf1, 0x74, 0x11, 0x12, 0xe0, 0x62, 0x76,
	0xcc, 0x49, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x24, 0xb8, 0xd8, 0x3d,
	0x52, 0x2b, 0x3c, 0x12, 0x8b, 0x33, 0x24, 0x98, 0xc0, 0xa2, 0x30, 0xae, 0x92, 0x35, 0x17, 0x3b,
	0x54, 0xa3, 0x90, 0x10, 0x17, 0x8b, 0x53, 0x7e, 0x4a, 0x25, 0x58, 0x1f, 0x4f, 0x10, 0x98, 0x2d,
	0x24, 0xc3, 0xc5, 0x52, 0x92, 0x98, 0x5e, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xa1,
	0x97, 0x58, 0x90, 0xa9, 0x17, 0x92, 0x98, 0x1e, 0x04, 0x16, 0x55, 0xd2, 0xe7, 0x62, 0x0e, 0x49,
	0x4c, 0x07, 0x69, 0xf4, 0x4b, 0xcc, 0x4d, 0x85, 0x5a, 0x08, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0x86,
	0x25, 0xe6, 0x94, 0xa6, 0x42, 0xed, 0x83, 0x70, 0x8c, 0xc2, 0xb9, 0xd8, 0x83, 0x21, 0x8e, 0x17,
	0x52, 0xe5, 0x62, 0x0e, 0x28, 0x2d, 0x11, 0xe2, 0x01, 0x1b, 0x09, 0x75, 0x82, 0x14, 0x1f, 0x32,
	0xcf, 0xd3, 0x45, 0x89, 0x01, 0xa4, 0xcc, 0x3d, 0xb5, 0x44, 0x08, 0x4d, 0x42, 0x0a, 0x45, 0x9b,
	0x12, 0x43, 0x12, 0x1b, 0x38, 0x2c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0xe2, 0xb5,
	0x8e, 0x1c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	Put(ctx context.Context, in *Content, opts ...grpc.CallOption) (*ContentID, error)
	Get(ctx context.Context, in *ContentID, opts ...grpc.CallOption) (*Content, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Put(ctx context.Context, in *Content, opts ...grpc.CallOption) (*ContentID, error) {
	out := new(ContentID)
	err := c.cc.Invoke(ctx, "/api.Storage/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Get(ctx context.Context, in *ContentID, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/api.Storage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	Put(context.Context, *Content) (*ContentID, error)
	Get(context.Context, *ContentID) (*Content, error)
}

// UnimplementedStorageServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (*UnimplementedStorageServer) Put(ctx context.Context, req *Content) (*ContentID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedStorageServer) Get(ctx context.Context, req *ContentID) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Storage/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Put(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Storage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*ContentID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Storage_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
