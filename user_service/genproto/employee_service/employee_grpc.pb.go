// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: employee.proto

package employee_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	Create(ctx context.Context, in *CreateEmployee, opts ...grpc.CallOption) (*GetEmployee, error)
	GetByID(ctx context.Context, in *EmployeePrimaryKey, opts ...grpc.CallOption) (*GetEmployee, error)
	GetList(ctx context.Context, in *GetListEmployeeRequest, opts ...grpc.CallOption) (*GetListEmployeeResponse, error)
	Update(ctx context.Context, in *UpdateEmployee, opts ...grpc.CallOption) (*GetEmployee, error)
	Delete(ctx context.Context, in *EmployeePrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
	Check(ctx context.Context, in *EmployeePrimaryKey, opts ...grpc.CallOption) (*CheckEmployeeResp, error)
	Login(ctx context.Context, in *EmployeeLoginRequest, opts ...grpc.CallOption) (*EmployeeLoginResponse, error)
	Register(ctx context.Context, in *EmployeeRegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RegisterConfirm(ctx context.Context, in *EmployeeRegisterConfRequest, opts ...grpc.CallOption) (*EmployeeLoginResponse, error)
	ChangePassword(ctx context.Context, in *EmployeeChangePassword, opts ...grpc.CallOption) (*EmployeeChangePasswordResp, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) Create(ctx context.Context, in *CreateEmployee, opts ...grpc.CallOption) (*GetEmployee, error) {
	out := new(GetEmployee)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetByID(ctx context.Context, in *EmployeePrimaryKey, opts ...grpc.CallOption) (*GetEmployee, error) {
	out := new(GetEmployee)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetList(ctx context.Context, in *GetListEmployeeRequest, opts ...grpc.CallOption) (*GetListEmployeeResponse, error) {
	out := new(GetListEmployeeResponse)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) Update(ctx context.Context, in *UpdateEmployee, opts ...grpc.CallOption) (*GetEmployee, error) {
	out := new(GetEmployee)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) Delete(ctx context.Context, in *EmployeePrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) Check(ctx context.Context, in *EmployeePrimaryKey, opts ...grpc.CallOption) (*CheckEmployeeResp, error) {
	out := new(CheckEmployeeResp)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) Login(ctx context.Context, in *EmployeeLoginRequest, opts ...grpc.CallOption) (*EmployeeLoginResponse, error) {
	out := new(EmployeeLoginResponse)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) Register(ctx context.Context, in *EmployeeRegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) RegisterConfirm(ctx context.Context, in *EmployeeRegisterConfRequest, opts ...grpc.CallOption) (*EmployeeLoginResponse, error) {
	out := new(EmployeeLoginResponse)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/RegisterConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ChangePassword(ctx context.Context, in *EmployeeChangePassword, opts ...grpc.CallOption) (*EmployeeChangePasswordResp, error) {
	out := new(EmployeeChangePasswordResp)
	err := c.cc.Invoke(ctx, "/employee_service_go.EmployeeService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations should embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	Create(context.Context, *CreateEmployee) (*GetEmployee, error)
	GetByID(context.Context, *EmployeePrimaryKey) (*GetEmployee, error)
	GetList(context.Context, *GetListEmployeeRequest) (*GetListEmployeeResponse, error)
	Update(context.Context, *UpdateEmployee) (*GetEmployee, error)
	Delete(context.Context, *EmployeePrimaryKey) (*empty.Empty, error)
	Check(context.Context, *EmployeePrimaryKey) (*CheckEmployeeResp, error)
	Login(context.Context, *EmployeeLoginRequest) (*EmployeeLoginResponse, error)
	Register(context.Context, *EmployeeRegisterRequest) (*empty.Empty, error)
	RegisterConfirm(context.Context, *EmployeeRegisterConfRequest) (*EmployeeLoginResponse, error)
	ChangePassword(context.Context, *EmployeeChangePassword) (*EmployeeChangePasswordResp, error)
}

// UnimplementedEmployeeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) Create(context.Context, *CreateEmployee) (*GetEmployee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEmployeeServiceServer) GetByID(context.Context, *EmployeePrimaryKey) (*GetEmployee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedEmployeeServiceServer) GetList(context.Context, *GetListEmployeeRequest) (*GetListEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedEmployeeServiceServer) Update(context.Context, *UpdateEmployee) (*GetEmployee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEmployeeServiceServer) Delete(context.Context, *EmployeePrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEmployeeServiceServer) Check(context.Context, *EmployeePrimaryKey) (*CheckEmployeeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedEmployeeServiceServer) Login(context.Context, *EmployeeLoginRequest) (*EmployeeLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedEmployeeServiceServer) Register(context.Context, *EmployeeRegisterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedEmployeeServiceServer) RegisterConfirm(context.Context, *EmployeeRegisterConfRequest) (*EmployeeLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterConfirm not implemented")
}
func (UnimplementedEmployeeServiceServer) ChangePassword(context.Context, *EmployeeChangePassword) (*EmployeeChangePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmployee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Create(ctx, req.(*CreateEmployee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetByID(ctx, req.(*EmployeePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetList(ctx, req.(*GetListEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Update(ctx, req.(*UpdateEmployee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Delete(ctx, req.(*EmployeePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Check(ctx, req.(*EmployeePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Login(ctx, req.(*EmployeeLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Register(ctx, req.(*EmployeeRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_RegisterConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRegisterConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).RegisterConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/RegisterConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).RegisterConfirm(ctx, req.(*EmployeeRegisterConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeChangePassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee_service_go.EmployeeService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ChangePassword(ctx, req.(*EmployeeChangePassword))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee_service_go.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EmployeeService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _EmployeeService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _EmployeeService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EmployeeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EmployeeService_Delete_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _EmployeeService_Check_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _EmployeeService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _EmployeeService_Register_Handler,
		},
		{
			MethodName: "RegisterConfirm",
			Handler:    _EmployeeService_RegisterConfirm_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _EmployeeService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee.proto",
}
