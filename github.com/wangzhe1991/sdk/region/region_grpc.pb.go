// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.0
// source: region.proto

package region

import (
	context "context"
	common "github.com/wangzhe1991/sdk/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RegionClient is the client API for Region service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionClient interface {
	List(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListRegionResp, error)
	ManualUpdate(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*common.EmptyResponse, error)
}

type regionClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionClient(cc grpc.ClientConnInterface) RegionClient {
	return &regionClient{cc}
}

func (c *regionClient) List(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListRegionResp, error) {
	out := new(ListRegionResp)
	err := c.cc.Invoke(ctx, "/grpc_sdk.region.Region/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) ManualUpdate(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*common.EmptyResponse, error) {
	out := new(common.EmptyResponse)
	err := c.cc.Invoke(ctx, "/grpc_sdk.region.Region/ManualUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionServer is the server API for Region service.
// All implementations must embed UnimplementedRegionServer
// for forward compatibility
type RegionServer interface {
	List(context.Context, *common.ListRequest) (*ListRegionResp, error)
	ManualUpdate(context.Context, *common.EmptyRequest) (*common.EmptyResponse, error)
	mustEmbedUnimplementedRegionServer()
}

// UnimplementedRegionServer must be embedded to have forward compatible implementations.
type UnimplementedRegionServer struct {
}

func (UnimplementedRegionServer) List(context.Context, *common.ListRequest) (*ListRegionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRegionServer) ManualUpdate(context.Context, *common.EmptyRequest) (*common.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManualUpdate not implemented")
}
func (UnimplementedRegionServer) mustEmbedUnimplementedRegionServer() {}

// UnsafeRegionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionServer will
// result in compilation errors.
type UnsafeRegionServer interface {
	mustEmbedUnimplementedRegionServer()
}

func RegisterRegionServer(s grpc.ServiceRegistrar, srv RegionServer) {
	s.RegisterService(&Region_ServiceDesc, srv)
}

func _Region_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_sdk.region.Region/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).List(ctx, req.(*common.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_ManualUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).ManualUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_sdk.region.Region/ManualUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).ManualUpdate(ctx, req.(*common.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Region_ServiceDesc is the grpc.ServiceDesc for Region service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Region_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_sdk.region.Region",
	HandlerType: (*RegionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Region_List_Handler,
		},
		{
			MethodName: "ManualUpdate",
			Handler:    _Region_ManualUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "region.proto",
}
