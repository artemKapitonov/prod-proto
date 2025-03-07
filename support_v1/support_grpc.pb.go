// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.3
// source: support_v1/support.proto

package support

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PongerClient is the client API for Ponger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PongerClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error)
}

type pongerClient struct {
	cc grpc.ClientConnInterface
}

func NewPongerClient(cc grpc.ClientConnInterface) PongerClient {
	return &pongerClient{cc}
}

func (c *pongerClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/proto.Ponger/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PongerServer is the server API for Ponger service.
// All implementations must embed UnimplementedPongerServer
// for forward compatibility
type PongerServer interface {
	Ping(context.Context, *emptypb.Empty) (*Pong, error)
	mustEmbedUnimplementedPongerServer()
}

// UnimplementedPongerServer must be embedded to have forward compatible implementations.
type UnimplementedPongerServer struct {
}

func (UnimplementedPongerServer) Ping(context.Context, *emptypb.Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPongerServer) mustEmbedUnimplementedPongerServer() {}

// UnsafePongerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PongerServer will
// result in compilation errors.
type UnsafePongerServer interface {
	mustEmbedUnimplementedPongerServer()
}

func RegisterPongerServer(s grpc.ServiceRegistrar, srv PongerServer) {
	s.RegisterService(&Ponger_ServiceDesc, srv)
}

func _Ponger_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PongerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ponger/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PongerServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Ponger_ServiceDesc is the grpc.ServiceDesc for Ponger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ponger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Ponger",
	HandlerType: (*PongerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Ponger_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "support_v1/support.proto",
}

// SupportChatServiceClient is the client API for SupportChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupportChatServiceClient interface {
	WriteMessage(ctx context.Context, in *WriteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*GetChatResponse, error)
	GetSupportChats(ctx context.Context, in *GetSupportChatsRequest, opts ...grpc.CallOption) (*GetSupportChatsResponse, error)
	ReadChat(ctx context.Context, in *ReadChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type supportChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupportChatServiceClient(cc grpc.ClientConnInterface) SupportChatServiceClient {
	return &supportChatServiceClient{cc}
}

func (c *supportChatServiceClient) WriteMessage(ctx context.Context, in *WriteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.SupportChatService/WriteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportChatServiceClient) GetChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*GetChatResponse, error) {
	out := new(GetChatResponse)
	err := c.cc.Invoke(ctx, "/proto.SupportChatService/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportChatServiceClient) GetSupportChats(ctx context.Context, in *GetSupportChatsRequest, opts ...grpc.CallOption) (*GetSupportChatsResponse, error) {
	out := new(GetSupportChatsResponse)
	err := c.cc.Invoke(ctx, "/proto.SupportChatService/GetSupportChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportChatServiceClient) ReadChat(ctx context.Context, in *ReadChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.SupportChatService/ReadChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupportChatServiceServer is the server API for SupportChatService service.
// All implementations must embed UnimplementedSupportChatServiceServer
// for forward compatibility
type SupportChatServiceServer interface {
	WriteMessage(context.Context, *WriteMessageRequest) (*emptypb.Empty, error)
	GetChat(context.Context, *ChatRequest) (*GetChatResponse, error)
	GetSupportChats(context.Context, *GetSupportChatsRequest) (*GetSupportChatsResponse, error)
	ReadChat(context.Context, *ReadChatRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSupportChatServiceServer()
}

// UnimplementedSupportChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupportChatServiceServer struct {
}

func (UnimplementedSupportChatServiceServer) WriteMessage(context.Context, *WriteMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMessage not implemented")
}
func (UnimplementedSupportChatServiceServer) GetChat(context.Context, *ChatRequest) (*GetChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedSupportChatServiceServer) GetSupportChats(context.Context, *GetSupportChatsRequest) (*GetSupportChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportChats not implemented")
}
func (UnimplementedSupportChatServiceServer) ReadChat(context.Context, *ReadChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadChat not implemented")
}
func (UnimplementedSupportChatServiceServer) mustEmbedUnimplementedSupportChatServiceServer() {}

// UnsafeSupportChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupportChatServiceServer will
// result in compilation errors.
type UnsafeSupportChatServiceServer interface {
	mustEmbedUnimplementedSupportChatServiceServer()
}

func RegisterSupportChatServiceServer(s grpc.ServiceRegistrar, srv SupportChatServiceServer) {
	s.RegisterService(&SupportChatService_ServiceDesc, srv)
}

func _SupportChatService_WriteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportChatServiceServer).WriteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SupportChatService/WriteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportChatServiceServer).WriteMessage(ctx, req.(*WriteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportChatService_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportChatServiceServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SupportChatService/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportChatServiceServer).GetChat(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportChatService_GetSupportChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupportChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportChatServiceServer).GetSupportChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SupportChatService/GetSupportChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportChatServiceServer).GetSupportChats(ctx, req.(*GetSupportChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportChatService_ReadChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportChatServiceServer).ReadChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SupportChatService/ReadChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportChatServiceServer).ReadChat(ctx, req.(*ReadChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupportChatService_ServiceDesc is the grpc.ServiceDesc for SupportChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupportChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SupportChatService",
	HandlerType: (*SupportChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteMessage",
			Handler:    _SupportChatService_WriteMessage_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _SupportChatService_GetChat_Handler,
		},
		{
			MethodName: "GetSupportChats",
			Handler:    _SupportChatService_GetSupportChats_Handler,
		},
		{
			MethodName: "ReadChat",
			Handler:    _SupportChatService_ReadChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "support_v1/support.proto",
}

// SupportManagerServiceClient is the client API for SupportManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupportManagerServiceClient interface {
	AddSupport(ctx context.Context, in *SupportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveSupport(ctx context.Context, in *SupportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type supportManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupportManagerServiceClient(cc grpc.ClientConnInterface) SupportManagerServiceClient {
	return &supportManagerServiceClient{cc}
}

func (c *supportManagerServiceClient) AddSupport(ctx context.Context, in *SupportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.SupportManagerService/AddSupport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supportManagerServiceClient) RemoveSupport(ctx context.Context, in *SupportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.SupportManagerService/RemoveSupport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupportManagerServiceServer is the server API for SupportManagerService service.
// All implementations must embed UnimplementedSupportManagerServiceServer
// for forward compatibility
type SupportManagerServiceServer interface {
	AddSupport(context.Context, *SupportRequest) (*emptypb.Empty, error)
	RemoveSupport(context.Context, *SupportRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSupportManagerServiceServer()
}

// UnimplementedSupportManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupportManagerServiceServer struct {
}

func (UnimplementedSupportManagerServiceServer) AddSupport(context.Context, *SupportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSupport not implemented")
}
func (UnimplementedSupportManagerServiceServer) RemoveSupport(context.Context, *SupportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSupport not implemented")
}
func (UnimplementedSupportManagerServiceServer) mustEmbedUnimplementedSupportManagerServiceServer() {}

// UnsafeSupportManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupportManagerServiceServer will
// result in compilation errors.
type UnsafeSupportManagerServiceServer interface {
	mustEmbedUnimplementedSupportManagerServiceServer()
}

func RegisterSupportManagerServiceServer(s grpc.ServiceRegistrar, srv SupportManagerServiceServer) {
	s.RegisterService(&SupportManagerService_ServiceDesc, srv)
}

func _SupportManagerService_AddSupport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportManagerServiceServer).AddSupport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SupportManagerService/AddSupport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportManagerServiceServer).AddSupport(ctx, req.(*SupportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupportManagerService_RemoveSupport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportManagerServiceServer).RemoveSupport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SupportManagerService/RemoveSupport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportManagerServiceServer).RemoveSupport(ctx, req.(*SupportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupportManagerService_ServiceDesc is the grpc.ServiceDesc for SupportManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupportManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SupportManagerService",
	HandlerType: (*SupportManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSupport",
			Handler:    _SupportManagerService_AddSupport_Handler,
		},
		{
			MethodName: "RemoveSupport",
			Handler:    _SupportManagerService_RemoveSupport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "support_v1/support.proto",
}
