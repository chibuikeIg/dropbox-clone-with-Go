// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: upload_data.proto

package uploaddataservice

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
	UploadData_SaveUploadData_FullMethodName   = "/UploadData/SaveUploadData"
	UploadData_FetchUploadData_FullMethodName  = "/UploadData/FetchUploadData"
	UploadData_DeleteUploadData_FullMethodName = "/UploadData/DeleteUploadData"
)

// UploadDataClient is the client API for UploadData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadDataClient interface {
	SaveUploadData(ctx context.Context, in *SaveUploadDataRequest, opts ...grpc.CallOption) (*SaveUploadDataResponse, error)
	FetchUploadData(ctx context.Context, in *UploadDataRequest, opts ...grpc.CallOption) (*UploadDataResponse, error)
	DeleteUploadData(ctx context.Context, in *DeleteUploadDataRequest, opts ...grpc.CallOption) (*DeleteUploadDataResponse, error)
}

type uploadDataClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadDataClient(cc grpc.ClientConnInterface) UploadDataClient {
	return &uploadDataClient{cc}
}

func (c *uploadDataClient) SaveUploadData(ctx context.Context, in *SaveUploadDataRequest, opts ...grpc.CallOption) (*SaveUploadDataResponse, error) {
	out := new(SaveUploadDataResponse)
	err := c.cc.Invoke(ctx, UploadData_SaveUploadData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadDataClient) FetchUploadData(ctx context.Context, in *UploadDataRequest, opts ...grpc.CallOption) (*UploadDataResponse, error) {
	out := new(UploadDataResponse)
	err := c.cc.Invoke(ctx, UploadData_FetchUploadData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadDataClient) DeleteUploadData(ctx context.Context, in *DeleteUploadDataRequest, opts ...grpc.CallOption) (*DeleteUploadDataResponse, error) {
	out := new(DeleteUploadDataResponse)
	err := c.cc.Invoke(ctx, UploadData_DeleteUploadData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadDataServer is the server API for UploadData service.
// All implementations must embed UnimplementedUploadDataServer
// for forward compatibility
type UploadDataServer interface {
	SaveUploadData(context.Context, *SaveUploadDataRequest) (*SaveUploadDataResponse, error)
	FetchUploadData(context.Context, *UploadDataRequest) (*UploadDataResponse, error)
	DeleteUploadData(context.Context, *DeleteUploadDataRequest) (*DeleteUploadDataResponse, error)
	mustEmbedUnimplementedUploadDataServer()
}

// UnimplementedUploadDataServer must be embedded to have forward compatible implementations.
type UnimplementedUploadDataServer struct {
}

func (UnimplementedUploadDataServer) SaveUploadData(context.Context, *SaveUploadDataRequest) (*SaveUploadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUploadData not implemented")
}
func (UnimplementedUploadDataServer) FetchUploadData(context.Context, *UploadDataRequest) (*UploadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUploadData not implemented")
}
func (UnimplementedUploadDataServer) DeleteUploadData(context.Context, *DeleteUploadDataRequest) (*DeleteUploadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUploadData not implemented")
}
func (UnimplementedUploadDataServer) mustEmbedUnimplementedUploadDataServer() {}

// UnsafeUploadDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadDataServer will
// result in compilation errors.
type UnsafeUploadDataServer interface {
	mustEmbedUnimplementedUploadDataServer()
}

func RegisterUploadDataServer(s grpc.ServiceRegistrar, srv UploadDataServer) {
	s.RegisterService(&UploadData_ServiceDesc, srv)
}

func _UploadData_SaveUploadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUploadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadDataServer).SaveUploadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadData_SaveUploadData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadDataServer).SaveUploadData(ctx, req.(*SaveUploadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadData_FetchUploadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadDataServer).FetchUploadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadData_FetchUploadData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadDataServer).FetchUploadData(ctx, req.(*UploadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadData_DeleteUploadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUploadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadDataServer).DeleteUploadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadData_DeleteUploadData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadDataServer).DeleteUploadData(ctx, req.(*DeleteUploadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadData_ServiceDesc is the grpc.ServiceDesc for UploadData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UploadData",
	HandlerType: (*UploadDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUploadData",
			Handler:    _UploadData_SaveUploadData_Handler,
		},
		{
			MethodName: "FetchUploadData",
			Handler:    _UploadData_FetchUploadData_Handler,
		},
		{
			MethodName: "DeleteUploadData",
			Handler:    _UploadData_DeleteUploadData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload_data.proto",
}