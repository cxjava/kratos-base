// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: v1/catalog.proto

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
	Catalog_CreateBeer_FullMethodName = "/catalog.service.v1.Catalog/CreateBeer"
	Catalog_GetBeer_FullMethodName    = "/catalog.service.v1.Catalog/GetBeer"
)

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	CreateBeer(ctx context.Context, in *CreateBeerReq, opts ...grpc.CallOption) (*CreateBeerReply, error)
	GetBeer(ctx context.Context, in *GetBeerReq, opts ...grpc.CallOption) (*GetBeerReply, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) CreateBeer(ctx context.Context, in *CreateBeerReq, opts ...grpc.CallOption) (*CreateBeerReply, error) {
	out := new(CreateBeerReply)
	err := c.cc.Invoke(ctx, Catalog_CreateBeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetBeer(ctx context.Context, in *GetBeerReq, opts ...grpc.CallOption) (*GetBeerReply, error) {
	out := new(GetBeerReply)
	err := c.cc.Invoke(ctx, Catalog_GetBeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	CreateBeer(context.Context, *CreateBeerReq) (*CreateBeerReply, error)
	GetBeer(context.Context, *GetBeerReq) (*GetBeerReply, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) CreateBeer(context.Context, *CreateBeerReq) (*CreateBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBeer not implemented")
}
func (UnimplementedCatalogServer) GetBeer(context.Context, *GetBeerReq) (*GetBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeer not implemented")
}
func (UnimplementedCatalogServer) mustEmbedUnimplementedCatalogServer() {}

// UnsafeCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServer will
// result in compilation errors.
type UnsafeCatalogServer interface {
	mustEmbedUnimplementedCatalogServer()
}

func RegisterCatalogServer(s grpc.ServiceRegistrar, srv CatalogServer) {
	s.RegisterService(&Catalog_ServiceDesc, srv)
}

func _Catalog_CreateBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_CreateBeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateBeer(ctx, req.(*CreateBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_GetBeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetBeer(ctx, req.(*GetBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.service.v1.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBeer",
			Handler:    _Catalog_CreateBeer_Handler,
		},
		{
			MethodName: "GetBeer",
			Handler:    _Catalog_GetBeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/catalog.proto",
}