package service

import (
	"context"
	"errors"
	"fmt"
	"go_user_service/config"
	"go_user_service/genproto/employee_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/pkg/jwt"
	"go_user_service/pkg/smtp"
	"time"

	"go_user_service/grpc/client"
	"go_user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EmployeeService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	redis    storage.IRedisStorage
}

func NewEmployeeService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI, redis storage.IRedisStorage) *EmployeeService {
	return &EmployeeService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
		redis:    redis,
	}
}

func (f *EmployeeService) Create(ctx context.Context, req *employee_service.CreateEmployee) (*employee_service.GetEmployee, error) {

	f.log.Info("---CreateEmployee--->>>", logger.Any("req", req))

	resp, err := f.strg.Employee().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateEmployee--->>>", logger.Error(err))
		return &employee_service.GetEmployee{}, err
	}

	return resp, nil
}
func (f *EmployeeService) Update(ctx context.Context, req *employee_service.UpdateEmployee) (*employee_service.GetEmployee, error) {

	f.log.Info("---UpdateEmployee--->>>", logger.Any("req", req))

	resp, err := f.strg.Employee().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateEmployee--->>>", logger.Error(err))
		return &employee_service.GetEmployee{}, err
	}

	return resp, nil
}

func (f *EmployeeService) GetList(ctx context.Context, req *employee_service.GetListEmployeeRequest) (*employee_service.GetListEmployeeResponse, error) {
	f.log.Info("---GetListEmployee--->>>", logger.Any("req", req))

	resp, err := f.strg.Employee().GetAll(ctx, req)
	if err != nil {
		f.log.Error("---GetListEmployee--->>>", logger.Error(err))
		return &employee_service.GetListEmployeeResponse{}, err
	}

	return resp, nil
}

func (f *EmployeeService) GetByID(ctx context.Context, id *employee_service.EmployeePrimaryKey) (*employee_service.GetEmployee, error) {
	f.log.Info("---GetEmployee--->>>", logger.Any("req", id))

	resp, err := f.strg.Employee().GetById(ctx, id)
	if err != nil {
		f.log.Error("---GetEmployee--->>>", logger.Error(err))
		return &employee_service.GetEmployee{}, err
	}

	return resp, nil
}

func (f *EmployeeService) Delete(ctx context.Context, req *employee_service.EmployeePrimaryKey) (*emptypb.Empty, error) {

	f.log.Info("---DeleteEmployee--->>>", logger.Any("req", req))

	_, err := f.strg.Employee().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteEmployee--->>>", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (f *EmployeeService) Check(ctx context.Context, id *employee_service.EmployeePrimaryKey) (*employee_service.CheckEmployeeResp, error) {
	f.log.Info("---GetEmployee--->>>", logger.Any("req", id))

	resp, err := f.strg.Employee().Check(ctx, id)
	if err != nil {
		f.log.Error("---GetEmployee--->>>", logger.Error(err))
		return &employee_service.CheckEmployeeResp{}, err
	}

	return resp, nil
}

func (a *EmployeeService) Login(ctx context.Context, loginRequest *employee_service.EmployeeLoginRequest) (*employee_service.EmployeeLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.EmployeeLogin)
	employee, err := a.strg.Employee().GetByLogin(ctx, loginRequest.EmployeeLogin)
	if err != nil {
		a.log.Error("error while getting employee credentials by login", logger.Error(err))
		return &employee_service.EmployeeLoginResponse{}, err
	}

	if err = hash.CompareHashAndPassword(employee.EmployeePassword, loginRequest.EmployeePassword); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return &employee_service.EmployeeLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["employee_id"] = employee.Id
	m["employee_role"] = config.USER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for employee login", logger.Error(err))
		return &employee_service.EmployeeLoginResponse{}, err
	}

	return &employee_service.EmployeeLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *EmployeeService) Register(ctx context.Context, loginRequest *employee_service.EmployeeRegisterRequest) (*emptypb.Empty, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Mail)

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering CRM system. Don't give it to anyone", otpCode)

	err := a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis employee register", logger.Error(err))
		return &emptypb.Empty{}, err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to employee register", logger.Error(err))
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (a *EmployeeService) RegisterConfirm(ctx context.Context, req *employee_service.EmployeeRegisterConfRequest) (*employee_service.EmployeeLoginResponse, error) {
	resp := &employee_service.EmployeeLoginResponse{}

	otp, err := a.redis.Get(ctx, req.Mail)
	if err != nil {
		a.log.Error("error while getting otp code for employee register confirm", logger.Error(err))
		return resp, err
	}
	if req.Otp != otp {
		a.log.Error("incorrect otp code for employee register confirm", logger.Error(err))
		return resp, errors.New("incorrect otp code")
	}
	req.Employee[0].Email = req.Mail

	id, err := a.strg.Employee().Create(ctx, req.Employee[0])
	if err != nil {
		a.log.Error("error while creating employee", logger.Error(err))
		return resp, err
	}
	var m = make(map[interface{}]interface{})

	m["employee_id"] = id
	m["employee_role"] = config.USER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for employee register confirm", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (f *EmployeeService) ChangePassword(ctx context.Context, pass *employee_service.EmployeeChangePassword) (*employee_service.EmployeeChangePasswordResp, error) {
	f.log.Info("---ChangePassword--->>>", logger.Any("req", pass))

	resp, err := f.strg.Employee().ChangePassword(ctx, pass)
	if err != nil {
		f.log.Error("---ChangePassword--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
