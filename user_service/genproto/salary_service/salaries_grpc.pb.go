// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: salaries.proto

package salary_service

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

// SalaryServiceClient is the client API for SalaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SalaryServiceClient interface {
	Create(ctx context.Context, in *CreateSalary, opts ...grpc.CallOption) (*GetSalary, error)
	GetByID(ctx context.Context, in *SalaryPrimaryKey, opts ...grpc.CallOption) (*GetSalary, error)
	GetList(ctx context.Context, in *GetListSalaryRequest, opts ...grpc.CallOption) (*GetListSalaryResponse, error)
	Update(ctx context.Context, in *UpdateSalary, opts ...grpc.CallOption) (*GetSalary, error)
	Delete(ctx context.Context, in *SalaryPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
	Check(ctx context.Context, in *SalaryPrimaryKey, opts ...grpc.CallOption) (*CheckSalaryResp, error)
}

type salaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSalaryServiceClient(cc grpc.ClientConnInterface) SalaryServiceClient {
	return &salaryServiceClient{cc}
}

func (c *salaryServiceClient) Create(ctx context.Context, in *CreateSalary, opts ...grpc.CallOption) (*GetSalary, error) {
	out := new(GetSalary)
	err := c.cc.Invoke(ctx, "/salary_service_go.SalaryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salaryServiceClient) GetByID(ctx context.Context, in *SalaryPrimaryKey, opts ...grpc.CallOption) (*GetSalary, error) {
	out := new(GetSalary)
	err := c.cc.Invoke(ctx, "/salary_service_go.SalaryService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salaryServiceClient) GetList(ctx context.Context, in *GetListSalaryRequest, opts ...grpc.CallOption) (*GetListSalaryResponse, error) {
	out := new(GetListSalaryResponse)
	err := c.cc.Invoke(ctx, "/salary_service_go.SalaryService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salaryServiceClient) Update(ctx context.Context, in *UpdateSalary, opts ...grpc.CallOption) (*GetSalary, error) {
	out := new(GetSalary)
	err := c.cc.Invoke(ctx, "/salary_service_go.SalaryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salaryServiceClient) Delete(ctx context.Context, in *SalaryPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/salary_service_go.SalaryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salaryServiceClient) Check(ctx context.Context, in *SalaryPrimaryKey, opts ...grpc.CallOption) (*CheckSalaryResp, error) {
	out := new(CheckSalaryResp)
	err := c.cc.Invoke(ctx, "/salary_service_go.SalaryService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SalaryServiceServer is the server API for SalaryService service.
// All implementations should embed UnimplementedSalaryServiceServer
// for forward compatibility
type SalaryServiceServer interface {
	Create(context.Context, *CreateSalary) (*GetSalary, error)
	GetByID(context.Context, *SalaryPrimaryKey) (*GetSalary, error)
	GetList(context.Context, *GetListSalaryRequest) (*GetListSalaryResponse, error)
	Update(context.Context, *UpdateSalary) (*GetSalary, error)
	Delete(context.Context, *SalaryPrimaryKey) (*empty.Empty, error)
	Check(context.Context, *SalaryPrimaryKey) (*CheckSalaryResp, error)
}

// UnimplementedSalaryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSalaryServiceServer struct {
}

func (UnimplementedSalaryServiceServer) Create(context.Context, *CreateSalary) (*GetSalary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSalaryServiceServer) GetByID(context.Context, *SalaryPrimaryKey) (*GetSalary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSalaryServiceServer) GetList(context.Context, *GetListSalaryRequest) (*GetListSalaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSalaryServiceServer) Update(context.Context, *UpdateSalary) (*GetSalary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSalaryServiceServer) Delete(context.Context, *SalaryPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSalaryServiceServer) Check(context.Context, *SalaryPrimaryKey) (*CheckSalaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

// UnsafeSalaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SalaryServiceServer will
// result in compilation errors.
type UnsafeSalaryServiceServer interface {
	mustEmbedUnimplementedSalaryServiceServer()
}

func RegisterSalaryServiceServer(s grpc.ServiceRegistrar, srv SalaryServiceServer) {
	s.RegisterService(&SalaryService_ServiceDesc, srv)
}

func _SalaryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSalary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalaryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salary_service_go.SalaryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalaryServiceServer).Create(ctx, req.(*CreateSalary))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalaryService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalaryPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalaryServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salary_service_go.SalaryService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalaryServiceServer).GetByID(ctx, req.(*SalaryPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalaryService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListSalaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalaryServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salary_service_go.SalaryService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalaryServiceServer).GetList(ctx, req.(*GetListSalaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalaryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSalary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalaryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salary_service_go.SalaryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalaryServiceServer).Update(ctx, req.(*UpdateSalary))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalaryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalaryPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalaryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salary_service_go.SalaryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalaryServiceServer).Delete(ctx, req.(*SalaryPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalaryService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalaryPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalaryServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/salary_service_go.SalaryService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalaryServiceServer).Check(ctx, req.(*SalaryPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// SalaryService_ServiceDesc is the grpc.ServiceDesc for SalaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SalaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "salary_service_go.SalaryService",
	HandlerType: (*SalaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SalaryService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SalaryService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SalaryService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SalaryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SalaryService_Delete_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _SalaryService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "salaries.proto",
}
