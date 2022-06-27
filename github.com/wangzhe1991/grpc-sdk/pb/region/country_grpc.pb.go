// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.0
// source: country.proto

package region

import (
	context "context"
	common "github.com/wangzhe1991/grpc-sdk/pb/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CountryClient is the client API for Country service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountryClient interface {
	List(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*CountryListResp, error)
	Detail(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*CountryListResp, error)
}

type countryClient struct {
	cc grpc.ClientConnInterface
}

func NewCountryClient(cc grpc.ClientConnInterface) CountryClient {
	return &countryClient{cc}
}

func (c *countryClient) List(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*CountryListResp, error) {
	out := new(CountryListResp)
	err := c.cc.Invoke(ctx, "/grpc_sdk.region.Country/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryClient) Detail(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*CountryListResp, error) {
	out := new(CountryListResp)
	err := c.cc.Invoke(ctx, "/grpc_sdk.region.Country/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryServer is the server API for Country service.
// All implementations must embed UnimplementedCountryServer
// for forward compatibility
type CountryServer interface {
	List(context.Context, *common.EmptyRequest) (*CountryListResp, error)
	Detail(context.Context, *common.EmptyRequest) (*CountryListResp, error)
	mustEmbedUnimplementedCountryServer()
}

// UnimplementedCountryServer must be embedded to have forward compatible implementations.
type UnimplementedCountryServer struct {
}

func (UnimplementedCountryServer) List(context.Context, *common.EmptyRequest) (*CountryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCountryServer) Detail(context.Context, *common.EmptyRequest) (*CountryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedCountryServer) mustEmbedUnimplementedCountryServer() {}

// UnsafeCountryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountryServer will
// result in compilation errorx.
type UnsafeCountryServer interface {
	mustEmbedUnimplementedCountryServer()
}

func RegisterCountryServer(s grpc.ServiceRegistrar, srv CountryServer) {
	s.RegisterService(&Country_ServiceDesc, srv)
}

func _Country_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_sdk.region.Country/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServer).List(ctx, req.(*common.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Country_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_sdk.region.Country/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServer).Detail(ctx, req.(*common.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Country_ServiceDesc is the grpc.ServiceDesc for Country service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Country_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_sdk.region.Country",
	HandlerType: (*CountryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Country_List_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Country_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "country.proto",
}
