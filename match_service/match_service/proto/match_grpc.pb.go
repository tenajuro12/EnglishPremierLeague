// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/match.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MatchService_CreateMatch_FullMethodName = "/match.MatchService/CreateMatch"
	MatchService_GetMatch_FullMethodName    = "/match.MatchService/GetMatch"
	MatchService_UpdateMatch_FullMethodName = "/match.MatchService/UpdateMatch"
	MatchService_DeleteMatch_FullMethodName = "/match.MatchService/DeleteMatch"
	MatchService_ListMatches_FullMethodName = "/match.MatchService/ListMatches"
)

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchServiceClient interface {
	CreateMatch(ctx context.Context, in *CreateMatchRequest, opts ...grpc.CallOption) (*CreateMatchResponse, error)
	GetMatch(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*GetMatchResponse, error)
	UpdateMatch(ctx context.Context, in *UpdateMatchRequest, opts ...grpc.CallOption) (*UpdateMatchResponse, error)
	DeleteMatch(ctx context.Context, in *DeleteMatchRequest, opts ...grpc.CallOption) (*DeleteMatchResponse, error)
	ListMatches(ctx context.Context, in *ListMatchesRequest, opts ...grpc.CallOption) (*ListMatchesResponse, error)
}

type matchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchServiceClient(cc grpc.ClientConnInterface) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) CreateMatch(ctx context.Context, in *CreateMatchRequest, opts ...grpc.CallOption) (*CreateMatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMatchResponse)
	err := c.cc.Invoke(ctx, MatchService_CreateMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) GetMatch(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*GetMatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMatchResponse)
	err := c.cc.Invoke(ctx, MatchService_GetMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) UpdateMatch(ctx context.Context, in *UpdateMatchRequest, opts ...grpc.CallOption) (*UpdateMatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMatchResponse)
	err := c.cc.Invoke(ctx, MatchService_UpdateMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) DeleteMatch(ctx context.Context, in *DeleteMatchRequest, opts ...grpc.CallOption) (*DeleteMatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMatchResponse)
	err := c.cc.Invoke(ctx, MatchService_DeleteMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) ListMatches(ctx context.Context, in *ListMatchesRequest, opts ...grpc.CallOption) (*ListMatchesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMatchesResponse)
	err := c.cc.Invoke(ctx, MatchService_ListMatches_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
// All implementations must embed UnimplementedMatchServiceServer
// for forward compatibility
type MatchServiceServer interface {
	CreateMatch(context.Context, *CreateMatchRequest) (*CreateMatchResponse, error)
	GetMatch(context.Context, *GetMatchRequest) (*GetMatchResponse, error)
	UpdateMatch(context.Context, *UpdateMatchRequest) (*UpdateMatchResponse, error)
	DeleteMatch(context.Context, *DeleteMatchRequest) (*DeleteMatchResponse, error)
	ListMatches(context.Context, *ListMatchesRequest) (*ListMatchesResponse, error)
	mustEmbedUnimplementedMatchServiceServer()
}

// UnimplementedMatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (UnimplementedMatchServiceServer) CreateMatch(context.Context, *CreateMatchRequest) (*CreateMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatch not implemented")
}
func (UnimplementedMatchServiceServer) GetMatch(context.Context, *GetMatchRequest) (*GetMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatch not implemented")
}
func (UnimplementedMatchServiceServer) UpdateMatch(context.Context, *UpdateMatchRequest) (*UpdateMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMatch not implemented")
}
func (UnimplementedMatchServiceServer) DeleteMatch(context.Context, *DeleteMatchRequest) (*DeleteMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMatch not implemented")
}
func (UnimplementedMatchServiceServer) ListMatches(context.Context, *ListMatchesRequest) (*ListMatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMatches not implemented")
}
func (UnimplementedMatchServiceServer) mustEmbedUnimplementedMatchServiceServer() {}

// UnsafeMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchServiceServer will
// result in compilation errors.
type UnsafeMatchServiceServer interface {
	mustEmbedUnimplementedMatchServiceServer()
}

func RegisterMatchServiceServer(s grpc.ServiceRegistrar, srv MatchServiceServer) {
	s.RegisterService(&MatchService_ServiceDesc, srv)
}

func _MatchService_CreateMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).CreateMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchService_CreateMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).CreateMatch(ctx, req.(*CreateMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_GetMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).GetMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchService_GetMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).GetMatch(ctx, req.(*GetMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_UpdateMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).UpdateMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchService_UpdateMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).UpdateMatch(ctx, req.(*UpdateMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_DeleteMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).DeleteMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchService_DeleteMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).DeleteMatch(ctx, req.(*DeleteMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_ListMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).ListMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchService_ListMatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).ListMatches(ctx, req.(*ListMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchService_ServiceDesc is the grpc.ServiceDesc for MatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "match.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMatch",
			Handler:    _MatchService_CreateMatch_Handler,
		},
		{
			MethodName: "GetMatch",
			Handler:    _MatchService_GetMatch_Handler,
		},
		{
			MethodName: "UpdateMatch",
			Handler:    _MatchService_UpdateMatch_Handler,
		},
		{
			MethodName: "DeleteMatch",
			Handler:    _MatchService_DeleteMatch_Handler,
		},
		{
			MethodName: "ListMatches",
			Handler:    _MatchService_ListMatches_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/match.proto",
}