// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/taa_engine/v1/taa_engine.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaaEngine_CreateTaaEngine_FullMethodName = "/api.taa_engine.v1.TaaEngine/CreateTaaEngine"
)

// TaaEngineClient is the client API for TaaEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaaEngineClient interface {
	// Sample function
	CreateTaaEngine(ctx context.Context, in *CreateTaaEngineRequest, opts ...grpc.CallOption) (*CreateTaaEngineResponse, error)
}

type taaEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewTaaEngineClient(cc grpc.ClientConnInterface) TaaEngineClient {
	return &taaEngineClient{cc}
}

func (c *taaEngineClient) CreateTaaEngine(ctx context.Context, in *CreateTaaEngineRequest, opts ...grpc.CallOption) (*CreateTaaEngineResponse, error) {
	out := new(CreateTaaEngineResponse)
	err := c.cc.Invoke(ctx, TaaEngine_CreateTaaEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaaEngineServer is the server API for TaaEngine service.
// All implementations must embed UnimplementedTaaEngineServer
// for forward compatibility
type TaaEngineServer interface {
	// Sample function
	CreateTaaEngine(context.Context, *CreateTaaEngineRequest) (*CreateTaaEngineResponse, error)
	mustEmbedUnimplementedTaaEngineServer()
}

// UnimplementedTaaEngineServer must be embedded to have forward compatible implementations.
type UnimplementedTaaEngineServer struct {
}

func (UnimplementedTaaEngineServer) CreateTaaEngine(context.Context, *CreateTaaEngineRequest) (*CreateTaaEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaaEngine not implemented")
}
func (UnimplementedTaaEngineServer) mustEmbedUnimplementedTaaEngineServer() {}

// UnsafeTaaEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaaEngineServer will
// result in compilation errors.
type UnsafeTaaEngineServer interface {
	mustEmbedUnimplementedTaaEngineServer()
}

func RegisterTaaEngineServer(s grpc.ServiceRegistrar, srv TaaEngineServer) {
	s.RegisterService(&TaaEngine_ServiceDesc, srv)
}

func _TaaEngine_CreateTaaEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaaEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaaEngineServer).CreateTaaEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaaEngine_CreateTaaEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaaEngineServer).CreateTaaEngine(ctx, req.(*CreateTaaEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaaEngine_ServiceDesc is the grpc.ServiceDesc for TaaEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaaEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.taa_engine.v1.TaaEngine",
	HandlerType: (*TaaEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaaEngine",
			Handler:    _TaaEngine_CreateTaaEngine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/taa_engine/v1/taa_engine.proto",
}