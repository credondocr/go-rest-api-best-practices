// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.3
// source: proto/service.proto

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
	ExampleService_GetExample_FullMethodName     = "/miapi.ExampleService/GetExample"
	ExampleService_GetAllProducts_FullMethodName = "/miapi.ExampleService/GetAllProducts"
	ExampleService_SearchProducts_FullMethodName = "/miapi.ExampleService/SearchProducts"
	ExampleService_GetProduct_FullMethodName     = "/miapi.ExampleService/GetProduct"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	GetExample(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error)
	GetAllProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ExampleService_GetAllProductsClient, error)
	SearchProducts(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (ExampleService_SearchProductsClient, error)
	GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) GetExample(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleResponse)
	err := c.cc.Invoke(ctx, ExampleService_GetExample_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) GetAllProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ExampleService_GetAllProductsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[0], ExampleService_GetAllProducts_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceGetAllProductsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService_GetAllProductsClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type exampleServiceGetAllProductsClient struct {
	grpc.ClientStream
}

func (x *exampleServiceGetAllProductsClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) SearchProducts(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (ExampleService_SearchProductsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[1], ExampleService_SearchProducts_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceSearchProductsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService_SearchProductsClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type exampleServiceSearchProductsClient struct {
	grpc.ClientStream
}

func (x *exampleServiceSearchProductsClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, ExampleService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	GetExample(context.Context, *ExampleRequest) (*ExampleResponse, error)
	GetAllProducts(*Empty, ExampleService_GetAllProductsServer) error
	SearchProducts(*SearchRequest, ExampleService_SearchProductsServer) error
	GetProduct(context.Context, *ProductRequest) (*Product, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) GetExample(context.Context, *ExampleRequest) (*ExampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExample not implemented")
}
func (UnimplementedExampleServiceServer) GetAllProducts(*Empty, ExampleService_GetAllProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedExampleServiceServer) SearchProducts(*SearchRequest, ExampleService_SearchProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedExampleServiceServer) GetProduct(context.Context, *ProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_GetExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).GetExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_GetExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).GetExample(ctx, req.(*ExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_GetAllProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).GetAllProducts(m, &exampleServiceGetAllProductsServer{ServerStream: stream})
}

type ExampleService_GetAllProductsServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type exampleServiceGetAllProductsServer struct {
	grpc.ServerStream
}

func (x *exampleServiceGetAllProductsServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _ExampleService_SearchProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).SearchProducts(m, &exampleServiceSearchProductsServer{ServerStream: stream})
}

type ExampleService_SearchProductsServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type exampleServiceSearchProductsServer struct {
	grpc.ServerStream
}

func (x *exampleServiceSearchProductsServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _ExampleService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).GetProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miapi.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExample",
			Handler:    _ExampleService_GetExample_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ExampleService_GetProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllProducts",
			Handler:       _ExampleService_GetAllProducts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchProducts",
			Handler:       _ExampleService_SearchProducts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
