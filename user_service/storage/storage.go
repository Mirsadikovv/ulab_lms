package storage

import (
	"context"
	"go_user_service/genproto/admin_service"
	"go_user_service/genproto/employee_service"

	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	Admin() AdminRepoI
	Employee() EmployeeRepoI
	Redis() IRedisStorage
}

type AdminRepoI interface {
	Create(context.Context, *admin_service.CreateAdmin) (*admin_service.GetAdmin, error)
	Update(context.Context, *admin_service.UpdateAdmin) (*admin_service.GetAdmin, error)
	GetAll(context.Context, *admin_service.GetListAdminRequest) (*admin_service.GetListAdminResponse, error)
	GetById(context.Context, *admin_service.AdminPrimaryKey) (*admin_service.GetAdmin, error)
	Delete(context.Context, *admin_service.AdminPrimaryKey) (emptypb.Empty, error)
	ChangePassword(context.Context, *admin_service.AdminChangePassword) (*admin_service.AdminChangePasswordResp, error)
	GetByLogin(context.Context, string) (*admin_service.GetAdminByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type EmployeeRepoI interface {
	Create(context.Context, *employee_service.CreateEmployee) (*employee_service.GetEmployee, error)
	Update(context.Context, *employee_service.UpdateEmployee) (*employee_service.GetEmployee, error)
	GetAll(context.Context, *employee_service.GetListEmployeeRequest) (*employee_service.GetListEmployeeResponse, error)
	GetById(context.Context, *employee_service.EmployeePrimaryKey) (*employee_service.GetEmployee, error)
	Delete(context.Context, *employee_service.EmployeePrimaryKey) (emptypb.Empty, error)
	Check(context.Context, *employee_service.EmployeePrimaryKey) (*employee_service.CheckEmployeeResp, error)
	ChangePassword(context.Context, *employee_service.EmployeeChangePassword) (*employee_service.EmployeeChangePasswordResp, error)
	GetByLogin(context.Context, string) (*employee_service.GetEmployeeByLogin, error)
	GetPassword(context.Context, string) (string, error)
}

type IRedisStorage interface {
	SetX(context.Context, string, interface{}, time.Duration) error
	Get(context.Context, string) (interface{}, error)
	Del(context.Context, string) error
}
