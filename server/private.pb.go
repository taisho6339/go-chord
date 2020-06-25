// Code generated by protoc-gen-go. DO NOT EDIT.
// source: private.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type FindRequest struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2a91b51c7bdc125, []int{0}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*FindRequest)(nil), "server.FindRequest")
}

func init() {
	proto.RegisterFile("private.proto", fileDescriptor_d2a91b51c7bdc125)
}

var fileDescriptor_d2a91b51c7bdc125 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xdb, 0x1c, 0x02, 0x4e, 0xab, 0xc2, 0x56, 0xa4, 0x44, 0x04, 0x8d, 0x17, 0x4f, 0x1b,
	0x31, 0x28, 0xa8, 0x07, 0xa1, 0xa2, 0x20, 0x48, 0x29, 0xad, 0x27, 0x6f, 0x49, 0x76, 0x9a, 0x2e,
	0xa4, 0x99, 0xb8, 0x3b, 0x09, 0xe4, 0x23, 0xfb, 0x2d, 0x24, 0x8d, 0x15, 0xff, 0x20, 0xd4, 0xeb,
	0xbe, 0xf9, 0xbd, 0xdf, 0x5b, 0xd8, 0x2e, 0x8c, 0xae, 0x22, 0x46, 0x59, 0x18, 0x62, 0x12, 0xae,
	0x45, 0x53, 0xa1, 0xf1, 0x0e, 0x52, 0xa2, 0x34, 0xc3, 0x60, 0xf5, 0x1a, 0x97, 0xf3, 0x00, 0x97,
	0x05, 0xd7, 0xed, 0x91, 0x07, 0x39, 0xa9, 0x0f, 0xc0, 0x3f, 0x84, 0xde, 0x83, 0xce, 0xd5, 0x14,
	0x5f, 0x4b, 0xb4, 0x2c, 0x76, 0xc0, 0xd1, 0x6a, 0xd8, 0x3d, 0xea, 0x9e, 0xf6, 0xa7, 0x8e, 0x56,
	0xe7, 0x6f, 0x0e, 0xec, 0x3e, 0xe6, 0x8c, 0x26, 0x8f, 0xb2, 0x19, 0x9a, 0x4a, 0x27, 0x28, 0x42,
	0xd8, 0x9a, 0x95, 0x49, 0x82, 0xd6, 0x92, 0x11, 0xfb, 0xb2, 0x35, 0xc9, 0xb5, 0x49, 0xde, 0x37,
	0x26, 0xaf, 0x2f, 0xdb, 0x25, 0x72, 0x4c, 0x0a, 0xfd, 0x8e, 0xb8, 0x80, 0xde, 0xc4, 0xa0, 0xc2,
	0x7f, 0x62, 0x37, 0xb0, 0xd7, 0xcc, 0xfb, 0xf4, 0x8d, 0xea, 0xe7, 0x28, 0xce, 0x50, 0x0c, 0xd6,
	0x77, 0x5f, 0xc6, 0xff, 0x82, 0xaf, 0x61, 0xf0, 0x03, 0x7e, 0xd2, 0x96, 0x37, 0x63, 0x6f, 0x61,
	0xd8, 0xc4, 0x77, 0x19, 0x59, 0xb4, 0x3c, 0x31, 0x98, 0xa0, 0xd2, 0x79, 0xda, 0xa4, 0x9b, 0x15,
	0x9c, 0x81, 0x3b, 0x26, 0xd6, 0xf3, 0x5a, 0x7c, 0x4b, 0xbc, 0x3f, 0x7e, 0xee, 0x77, 0x46, 0x27,
	0x2f, 0xc7, 0xa9, 0xe6, 0x45, 0x19, 0xcb, 0x84, 0x96, 0x01, 0x47, 0xda, 0x2e, 0xe8, 0x32, 0x0c,
	0xaf, 0x82, 0x94, 0x8c, 0x0a, 0xda, 0x8e, 0xd8, 0x5d, 0x61, 0xe1, 0x7b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x87, 0x55, 0xa8, 0xca, 0xf8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InternalServiceClient is the client API for InternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalServiceClient interface {
	Successor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error)
	Predecessor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error)
	FindSuccessorByTable(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	FindSuccessorByList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	FindClosestPrecedingNode(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error)
}

type internalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalServiceClient(cc grpc.ClientConnInterface) InternalServiceClient {
	return &internalServiceClient{cc}
}

func (c *internalServiceClient) Successor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/Successor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Predecessor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/Predecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindSuccessorByTable(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindSuccessorByTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindSuccessorByList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindSuccessorByList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindClosestPrecedingNode(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindClosestPrecedingNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/server.InternalService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServiceServer is the server API for InternalService service.
type InternalServiceServer interface {
	Successor(context.Context, *empty.Empty) (*Node, error)
	Predecessor(context.Context, *empty.Empty) (*Node, error)
	FindSuccessorByTable(context.Context, *FindRequest) (*Node, error)
	FindSuccessorByList(context.Context, *FindRequest) (*Node, error)
	FindClosestPrecedingNode(context.Context, *FindRequest) (*Node, error)
	Notify(context.Context, *Node) (*empty.Empty, error)
}

// UnimplementedInternalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInternalServiceServer struct {
}

func (*UnimplementedInternalServiceServer) Successor(ctx context.Context, req *empty.Empty) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Successor not implemented")
}
func (*UnimplementedInternalServiceServer) Predecessor(ctx context.Context, req *empty.Empty) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predecessor not implemented")
}
func (*UnimplementedInternalServiceServer) FindSuccessorByTable(ctx context.Context, req *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorByTable not implemented")
}
func (*UnimplementedInternalServiceServer) FindSuccessorByList(ctx context.Context, req *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorByList not implemented")
}
func (*UnimplementedInternalServiceServer) FindClosestPrecedingNode(ctx context.Context, req *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindClosestPrecedingNode not implemented")
}
func (*UnimplementedInternalServiceServer) Notify(ctx context.Context, req *Node) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}

func RegisterInternalServiceServer(s *grpc.Server, srv InternalServiceServer) {
	s.RegisterService(&_InternalService_serviceDesc, srv)
}

func _InternalService_Successor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Successor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Successor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Successor(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Predecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Predecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Predecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Predecessor(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindSuccessorByTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindSuccessorByTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindSuccessorByTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindSuccessorByTable(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindSuccessorByList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindSuccessorByList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindSuccessorByList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindSuccessorByList(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindClosestPrecedingNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindClosestPrecedingNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindClosestPrecedingNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindClosestPrecedingNode(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Notify(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.InternalService",
	HandlerType: (*InternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Successor",
			Handler:    _InternalService_Successor_Handler,
		},
		{
			MethodName: "Predecessor",
			Handler:    _InternalService_Predecessor_Handler,
		},
		{
			MethodName: "FindSuccessorByTable",
			Handler:    _InternalService_FindSuccessorByTable_Handler,
		},
		{
			MethodName: "FindSuccessorByList",
			Handler:    _InternalService_FindSuccessorByList_Handler,
		},
		{
			MethodName: "FindClosestPrecedingNode",
			Handler:    _InternalService_FindClosestPrecedingNode_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _InternalService_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "private.proto",
}
