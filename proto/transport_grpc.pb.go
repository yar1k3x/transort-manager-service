// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/transport.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TransportService_CreateTransport_FullMethodName      = "/proto.TransportService/CreateTransport"
	TransportService_UpdateTransport_FullMethodName      = "/proto.TransportService/UpdateTransport"
	TransportService_GetTransportInfo_FullMethodName     = "/proto.TransportService/GetTransportInfo"
	TransportService_CreateTransportLog_FullMethodName   = "/proto.TransportService/CreateTransportLog"
	TransportService_GetTransportLogsInfo_FullMethodName = "/proto.TransportService/GetTransportLogsInfo"
	TransportService_GetTransportType_FullMethodName     = "/proto.TransportService/GetTransportType"
	TransportService_GetServiceType_FullMethodName       = "/proto.TransportService/GetServiceType"
)

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportServiceClient interface {
	CreateTransport(ctx context.Context, in *CreateTransportRequest, opts ...grpc.CallOption) (*CreateTransportResponse, error)
	UpdateTransport(ctx context.Context, in *UpdateTransportRequest, opts ...grpc.CallOption) (*UpdateTransportResponse, error)
	GetTransportInfo(ctx context.Context, in *GetTransportInfoRequest, opts ...grpc.CallOption) (*GetTransportInfoResponse, error)
	CreateTransportLog(ctx context.Context, in *CreateTransportLogRequest, opts ...grpc.CallOption) (*CreateTransportLogResponse, error)
	GetTransportLogsInfo(ctx context.Context, in *GetTransportLogsInfoRequest, opts ...grpc.CallOption) (*GetTransportLogsInfoResponse, error)
	GetTransportType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTransportTypeResponse, error)
	GetServiceType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetServiceTypeResponse, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) CreateTransport(ctx context.Context, in *CreateTransportRequest, opts ...grpc.CallOption) (*CreateTransportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransportResponse)
	err := c.cc.Invoke(ctx, TransportService_CreateTransport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) UpdateTransport(ctx context.Context, in *UpdateTransportRequest, opts ...grpc.CallOption) (*UpdateTransportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTransportResponse)
	err := c.cc.Invoke(ctx, TransportService_UpdateTransport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) GetTransportInfo(ctx context.Context, in *GetTransportInfoRequest, opts ...grpc.CallOption) (*GetTransportInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransportInfoResponse)
	err := c.cc.Invoke(ctx, TransportService_GetTransportInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) CreateTransportLog(ctx context.Context, in *CreateTransportLogRequest, opts ...grpc.CallOption) (*CreateTransportLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransportLogResponse)
	err := c.cc.Invoke(ctx, TransportService_CreateTransportLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) GetTransportLogsInfo(ctx context.Context, in *GetTransportLogsInfoRequest, opts ...grpc.CallOption) (*GetTransportLogsInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransportLogsInfoResponse)
	err := c.cc.Invoke(ctx, TransportService_GetTransportLogsInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) GetTransportType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTransportTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransportTypeResponse)
	err := c.cc.Invoke(ctx, TransportService_GetTransportType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) GetServiceType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetServiceTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServiceTypeResponse)
	err := c.cc.Invoke(ctx, TransportService_GetServiceType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
// All implementations must embed UnimplementedTransportServiceServer
// for forward compatibility.
type TransportServiceServer interface {
	CreateTransport(context.Context, *CreateTransportRequest) (*CreateTransportResponse, error)
	UpdateTransport(context.Context, *UpdateTransportRequest) (*UpdateTransportResponse, error)
	GetTransportInfo(context.Context, *GetTransportInfoRequest) (*GetTransportInfoResponse, error)
	CreateTransportLog(context.Context, *CreateTransportLogRequest) (*CreateTransportLogResponse, error)
	GetTransportLogsInfo(context.Context, *GetTransportLogsInfoRequest) (*GetTransportLogsInfoResponse, error)
	GetTransportType(context.Context, *emptypb.Empty) (*GetTransportTypeResponse, error)
	GetServiceType(context.Context, *emptypb.Empty) (*GetServiceTypeResponse, error)
	mustEmbedUnimplementedTransportServiceServer()
}

// UnimplementedTransportServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransportServiceServer struct{}

func (UnimplementedTransportServiceServer) CreateTransport(context.Context, *CreateTransportRequest) (*CreateTransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransport not implemented")
}
func (UnimplementedTransportServiceServer) UpdateTransport(context.Context, *UpdateTransportRequest) (*UpdateTransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransport not implemented")
}
func (UnimplementedTransportServiceServer) GetTransportInfo(context.Context, *GetTransportInfoRequest) (*GetTransportInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransportInfo not implemented")
}
func (UnimplementedTransportServiceServer) CreateTransportLog(context.Context, *CreateTransportLogRequest) (*CreateTransportLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransportLog not implemented")
}
func (UnimplementedTransportServiceServer) GetTransportLogsInfo(context.Context, *GetTransportLogsInfoRequest) (*GetTransportLogsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransportLogsInfo not implemented")
}
func (UnimplementedTransportServiceServer) GetTransportType(context.Context, *emptypb.Empty) (*GetTransportTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransportType not implemented")
}
func (UnimplementedTransportServiceServer) GetServiceType(context.Context, *emptypb.Empty) (*GetServiceTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceType not implemented")
}
func (UnimplementedTransportServiceServer) mustEmbedUnimplementedTransportServiceServer() {}
func (UnimplementedTransportServiceServer) testEmbeddedByValue()                          {}

// UnsafeTransportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServiceServer will
// result in compilation errors.
type UnsafeTransportServiceServer interface {
	mustEmbedUnimplementedTransportServiceServer()
}

func RegisterTransportServiceServer(s grpc.ServiceRegistrar, srv TransportServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransportServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransportService_ServiceDesc, srv)
}

func _TransportService_CreateTransport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).CreateTransport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_CreateTransport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).CreateTransport(ctx, req.(*CreateTransportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_UpdateTransport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).UpdateTransport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_UpdateTransport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).UpdateTransport(ctx, req.(*UpdateTransportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_GetTransportInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransportInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetTransportInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_GetTransportInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetTransportInfo(ctx, req.(*GetTransportInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_CreateTransportLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransportLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).CreateTransportLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_CreateTransportLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).CreateTransportLog(ctx, req.(*CreateTransportLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_GetTransportLogsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransportLogsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetTransportLogsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_GetTransportLogsInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetTransportLogsInfo(ctx, req.(*GetTransportLogsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_GetTransportType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetTransportType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_GetTransportType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetTransportType(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_GetServiceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetServiceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_GetServiceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetServiceType(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportService_ServiceDesc is the grpc.ServiceDesc for TransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TransportService",
	HandlerType: (*TransportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransport",
			Handler:    _TransportService_CreateTransport_Handler,
		},
		{
			MethodName: "UpdateTransport",
			Handler:    _TransportService_UpdateTransport_Handler,
		},
		{
			MethodName: "GetTransportInfo",
			Handler:    _TransportService_GetTransportInfo_Handler,
		},
		{
			MethodName: "CreateTransportLog",
			Handler:    _TransportService_CreateTransportLog_Handler,
		},
		{
			MethodName: "GetTransportLogsInfo",
			Handler:    _TransportService_GetTransportLogsInfo_Handler,
		},
		{
			MethodName: "GetTransportType",
			Handler:    _TransportService_GetTransportType_Handler,
		},
		{
			MethodName: "GetServiceType",
			Handler:    _TransportService_GetServiceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transport.proto",
}
