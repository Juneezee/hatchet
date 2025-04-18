// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: v1/workflows.proto

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	PutWorkflow(ctx context.Context, in *CreateWorkflowVersionRequest, opts ...grpc.CallOption) (*CreateWorkflowVersionResponse, error)
	CancelTasks(ctx context.Context, in *CancelTasksRequest, opts ...grpc.CallOption) (*CancelTasksResponse, error)
	ReplayTasks(ctx context.Context, in *ReplayTasksRequest, opts ...grpc.CallOption) (*ReplayTasksResponse, error)
	TriggerWorkflowRun(ctx context.Context, in *TriggerWorkflowRunRequest, opts ...grpc.CallOption) (*TriggerWorkflowRunResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) PutWorkflow(ctx context.Context, in *CreateWorkflowVersionRequest, opts ...grpc.CallOption) (*CreateWorkflowVersionResponse, error) {
	out := new(CreateWorkflowVersionResponse)
	err := c.cc.Invoke(ctx, "/v1.AdminService/PutWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CancelTasks(ctx context.Context, in *CancelTasksRequest, opts ...grpc.CallOption) (*CancelTasksResponse, error) {
	out := new(CancelTasksResponse)
	err := c.cc.Invoke(ctx, "/v1.AdminService/CancelTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ReplayTasks(ctx context.Context, in *ReplayTasksRequest, opts ...grpc.CallOption) (*ReplayTasksResponse, error) {
	out := new(ReplayTasksResponse)
	err := c.cc.Invoke(ctx, "/v1.AdminService/ReplayTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) TriggerWorkflowRun(ctx context.Context, in *TriggerWorkflowRunRequest, opts ...grpc.CallOption) (*TriggerWorkflowRunResponse, error) {
	out := new(TriggerWorkflowRunResponse)
	err := c.cc.Invoke(ctx, "/v1.AdminService/TriggerWorkflowRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	PutWorkflow(context.Context, *CreateWorkflowVersionRequest) (*CreateWorkflowVersionResponse, error)
	CancelTasks(context.Context, *CancelTasksRequest) (*CancelTasksResponse, error)
	ReplayTasks(context.Context, *ReplayTasksRequest) (*ReplayTasksResponse, error)
	TriggerWorkflowRun(context.Context, *TriggerWorkflowRunRequest) (*TriggerWorkflowRunResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) PutWorkflow(context.Context, *CreateWorkflowVersionRequest) (*CreateWorkflowVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutWorkflow not implemented")
}
func (UnimplementedAdminServiceServer) CancelTasks(context.Context, *CancelTasksRequest) (*CancelTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTasks not implemented")
}
func (UnimplementedAdminServiceServer) ReplayTasks(context.Context, *ReplayTasksRequest) (*ReplayTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplayTasks not implemented")
}
func (UnimplementedAdminServiceServer) TriggerWorkflowRun(context.Context, *TriggerWorkflowRunRequest) (*TriggerWorkflowRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerWorkflowRun not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_PutWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PutWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AdminService/PutWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PutWorkflow(ctx, req.(*CreateWorkflowVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CancelTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CancelTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AdminService/CancelTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CancelTasks(ctx, req.(*CancelTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ReplayTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplayTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ReplayTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AdminService/ReplayTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ReplayTasks(ctx, req.(*ReplayTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_TriggerWorkflowRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerWorkflowRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).TriggerWorkflowRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AdminService/TriggerWorkflowRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).TriggerWorkflowRun(ctx, req.(*TriggerWorkflowRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutWorkflow",
			Handler:    _AdminService_PutWorkflow_Handler,
		},
		{
			MethodName: "CancelTasks",
			Handler:    _AdminService_CancelTasks_Handler,
		},
		{
			MethodName: "ReplayTasks",
			Handler:    _AdminService_ReplayTasks_Handler,
		},
		{
			MethodName: "TriggerWorkflowRun",
			Handler:    _AdminService_TriggerWorkflowRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/workflows.proto",
}
