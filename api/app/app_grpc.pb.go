// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: app/app.proto

package app

import (
	context "context"
	kiae "github.com/kiaedev/kiae/api/kiae"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *Application, opts ...grpc.CallOption) (*Application, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Application, error)
	Delete(ctx context.Context, in *kiae.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Read(ctx context.Context, in *kiae.IdRequest, opts ...grpc.CallOption) (*Application, error)
	DoAction(ctx context.Context, in *ActionPayload, opts ...grpc.CallOption) (*Application, error)
	CfgCreate(ctx context.Context, in *AppCfg, opts ...grpc.CallOption) (*Configuration, error)
	CfgUpdate(ctx context.Context, in *AppCfg, opts ...grpc.CallOption) (*Configuration, error)
	CfgDelete(ctx context.Context, in *AppCfg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EnvCreate(ctx context.Context, in *AppEnv, opts ...grpc.CallOption) (*Environment, error)
	EnvUpdate(ctx context.Context, in *AppEnv, opts ...grpc.CallOption) (*Environment, error)
	EnvDelete(ctx context.Context, in *AppEnv, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/app.AppService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Create(ctx context.Context, in *Application, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/app.AppService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/app.AppService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Delete(ctx context.Context, in *kiae.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/app.AppService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Read(ctx context.Context, in *kiae.IdRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/app.AppService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DoAction(ctx context.Context, in *ActionPayload, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/app.AppService/DoAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) CfgCreate(ctx context.Context, in *AppCfg, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/app.AppService/CfgCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) CfgUpdate(ctx context.Context, in *AppCfg, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/app.AppService/CfgUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) CfgDelete(ctx context.Context, in *AppCfg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/app.AppService/CfgDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) EnvCreate(ctx context.Context, in *AppEnv, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := c.cc.Invoke(ctx, "/app.AppService/EnvCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) EnvUpdate(ctx context.Context, in *AppEnv, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := c.cc.Invoke(ctx, "/app.AppService/EnvUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) EnvDelete(ctx context.Context, in *AppEnv, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/app.AppService/EnvDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *Application) (*Application, error)
	Update(context.Context, *UpdateRequest) (*Application, error)
	Delete(context.Context, *kiae.IdRequest) (*emptypb.Empty, error)
	Read(context.Context, *kiae.IdRequest) (*Application, error)
	DoAction(context.Context, *ActionPayload) (*Application, error)
	CfgCreate(context.Context, *AppCfg) (*Configuration, error)
	CfgUpdate(context.Context, *AppCfg) (*Configuration, error)
	CfgDelete(context.Context, *AppCfg) (*emptypb.Empty, error)
	EnvCreate(context.Context, *AppEnv) (*Environment, error)
	EnvUpdate(context.Context, *AppEnv) (*Environment, error)
	EnvDelete(context.Context, *AppEnv) (*emptypb.Empty, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAppServiceServer) Create(context.Context, *Application) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAppServiceServer) Update(context.Context, *UpdateRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAppServiceServer) Delete(context.Context, *kiae.IdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAppServiceServer) Read(context.Context, *kiae.IdRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedAppServiceServer) DoAction(context.Context, *ActionPayload) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoAction not implemented")
}
func (UnimplementedAppServiceServer) CfgCreate(context.Context, *AppCfg) (*Configuration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CfgCreate not implemented")
}
func (UnimplementedAppServiceServer) CfgUpdate(context.Context, *AppCfg) (*Configuration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CfgUpdate not implemented")
}
func (UnimplementedAppServiceServer) CfgDelete(context.Context, *AppCfg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CfgDelete not implemented")
}
func (UnimplementedAppServiceServer) EnvCreate(context.Context, *AppEnv) (*Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnvCreate not implemented")
}
func (UnimplementedAppServiceServer) EnvUpdate(context.Context, *AppEnv) (*Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnvUpdate not implemented")
}
func (UnimplementedAppServiceServer) EnvDelete(context.Context, *AppEnv) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnvDelete not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Create(ctx, req.(*Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kiae.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Delete(ctx, req.(*kiae.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kiae.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Read(ctx, req.(*kiae.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DoAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DoAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/DoAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DoAction(ctx, req.(*ActionPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_CfgCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppCfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CfgCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/CfgCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CfgCreate(ctx, req.(*AppCfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_CfgUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppCfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CfgUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/CfgUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CfgUpdate(ctx, req.(*AppCfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_CfgDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppCfg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CfgDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/CfgDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CfgDelete(ctx, req.(*AppCfg))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_EnvCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppEnv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).EnvCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/EnvCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).EnvCreate(ctx, req.(*AppEnv))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_EnvUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppEnv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).EnvUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/EnvUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).EnvUpdate(ctx, req.(*AppEnv))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_EnvDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppEnv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).EnvDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/EnvDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).EnvDelete(ctx, req.(*AppEnv))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AppService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AppService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AppService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AppService_Delete_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _AppService_Read_Handler,
		},
		{
			MethodName: "DoAction",
			Handler:    _AppService_DoAction_Handler,
		},
		{
			MethodName: "CfgCreate",
			Handler:    _AppService_CfgCreate_Handler,
		},
		{
			MethodName: "CfgUpdate",
			Handler:    _AppService_CfgUpdate_Handler,
		},
		{
			MethodName: "CfgDelete",
			Handler:    _AppService_CfgDelete_Handler,
		},
		{
			MethodName: "EnvCreate",
			Handler:    _AppService_EnvCreate_Handler,
		},
		{
			MethodName: "EnvUpdate",
			Handler:    _AppService_EnvUpdate_Handler,
		},
		{
			MethodName: "EnvDelete",
			Handler:    _AppService_EnvDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/app.proto",
}
