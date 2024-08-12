package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_user_service/genproto/employee_service"
	"go_user_service/pkg"
	"go_user_service/pkg/hash"
	"go_user_service/pkg/logger"
	"go_user_service/storage"
	"log"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type employeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) storage.EmployeeRepoI {
	return &employeeRepo{
		db: db,
	}
}

func generateEmployeeLogin(db *pgxpool.Pool, ctx context.Context) (string, error) {
	var nextVal int
	err := db.QueryRow(ctx, "SELECT nextval('employee_external_id_seq')").Scan(&nextVal)
	if err != nil {
		return "", err
	}
	employeeLogin := "S" + fmt.Sprintf("%05d", nextVal)
	return employeeLogin, nil
}

func (c *employeeRepo) Create(ctx context.Context, req *employee_service.CreateEmployee) (*employee_service.GetEmployee, error) {
	var birthday sql.NullString
	id := uuid.NewString()
	pasword, err := hash.HashPassword(req.EmployeePassword)
	if err != nil {
		log.Println("error while hashing password", err)
	}

	employeeLogin, err := generateEmployeeLogin(c.db, ctx)
	if err != nil {
		log.Println("error while generating login", err)
	}

	comtag, err := c.db.Exec(ctx, `
		INSERT INTO employees (
			id,
			employee_login,
			birthday,
			gender,
			fullname,
			email,
			phone,
			employee_password
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8
		)`,
		id,
		employeeLogin,
		birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		pasword)
	if err != nil {
		log.Println("error while creating employee", comtag)
		return nil, err
	}
	req.Birthday = pkg.NullStringToString(birthday)

	employee, err := c.GetById(ctx, &employee_service.EmployeePrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting employee by id")
		return nil, err
	}
	return employee, nil
}

func (c *employeeRepo) Update(ctx context.Context, req *employee_service.UpdateEmployee) (*employee_service.GetEmployee, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE employees SET
		birthday = $1,
		gender = $2,
		fullname = $3,
		email = $4,
		phone = $5,
		updated_at = NOW()
		WHERE id = $6
		`,
		req.Birthday,
		req.Gender,
		req.Fullname,
		req.Email,
		req.Phone,
		req.Id)
	if err != nil {
		log.Println("error while updating employee")
		return nil, err
	}

	employee, err := c.GetById(ctx, &employee_service.EmployeePrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting employee by id")
		return nil, err
	}
	return employee, nil
}

func (c *employeeRepo) GetAll(ctx context.Context, req *employee_service.GetListEmployeeRequest) (*employee_service.GetListEmployeeResponse, error) {
	employees := employee_service.GetListEmployeeResponse{}

	var (
		birthday   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)
	filter_by_name := ""
	offest := (req.Offset - 1) * req.Limit
	if req.Search != "" {
		filter_by_name = fmt.Sprintf(`AND fullname ILIKE '%%%v%%'`, req.Search)
	}
	query := `SELECT
				id,
				employee_login,
				birthday,
				gender,
				fullname,
				email,
				phone,
				created_at,
				updated_at
			FROM employees
			WHERE TRUE AND deleted_at is null ` + filter_by_name + `
			OFFSET $1 LIMIT $2
`
	rows, err := c.db.Query(ctx, query, offest, req.Limit)

	if err != nil {
		log.Println("error while getting all employees")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			employee employee_service.GetEmployee
		)
		if err = rows.Scan(
			&employee.Id,
			&employee.EmployeeLogin,
			&birthday,
			&employee.Gender,
			&employee.Fullname,
			&employee.Email,
			&employee.Phone,
			&created_at,
			&updated_at,
		); err != nil {
			return &employees, err
		}
		employee.Birthday = pkg.NullStringToString(birthday)
		employee.CreatedAt = pkg.NullStringToString(created_at)
		employee.UpdatedAt = pkg.NullStringToString(updated_at)

		employees.Employees = append(employees.Employees, &employee)
	}

	err = c.db.QueryRow(ctx, `SELECT count(*) from employees WHERE TRUE AND deleted_at is null `+filter_by_name+``).Scan(&employees.Count)
	if err != nil {
		return &employees, err
	}

	return &employees, nil
}

func (c *employeeRepo) GetById(ctx context.Context, id *employee_service.EmployeePrimaryKey) (*employee_service.GetEmployee, error) {
	var (
		employee   employee_service.GetEmployee
		birthday   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	query := `SELECT
				id,
				employee_login,
				birthday,
				gender,
				fullname,
				email,
				phone,
				created_at,
				updated_at
			FROM employees
			WHERE id = $1 AND deleted_at IS NULL`

	rows := c.db.QueryRow(ctx, query, id.Id)

	if err := rows.Scan(
		&employee.Id,
		&employee.EmployeeLogin,
		&birthday,
		&employee.Gender,
		&employee.Fullname,
		&employee.Email,
		&employee.Phone,
		&created_at,
		&updated_at); err != nil {
		return &employee, err
	}
	employee.Birthday = pkg.NullStringToString(birthday)
	employee.CreatedAt = pkg.NullStringToString(created_at)
	employee.UpdatedAt = pkg.NullStringToString(updated_at)

	return &employee, nil
}

func (c *employeeRepo) Delete(ctx context.Context, id *employee_service.EmployeePrimaryKey) (emptypb.Empty, error) {

	_, err := c.db.Exec(ctx, `
		UPDATE employees SET
		deleted_at = NOW()
		WHERE id = $1
		`,
		id.Id)

	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}

func (c *employeeRepo) Check(ctx context.Context, id *employee_service.EmployeePrimaryKey) (*employee_service.CheckEmployeeResp, error) {
	query := `SELECT EXISTS (
                SELECT 1
                FROM employees
                WHERE id = $1 AND deleted_at IS NULL
            )`

	var exists bool
	err := c.db.QueryRow(ctx, query, id.Id).Scan(&exists)
	if err != nil {
		return nil, err
	}

	resp := &employee_service.CheckEmployeeResp{
		Check: exists,
	}

	return resp, nil
}

func (c *employeeRepo) ChangePassword(ctx context.Context, pass *employee_service.EmployeeChangePassword) (*employee_service.EmployeeChangePasswordResp, error) {
	var hashedPass string
	var resp employee_service.EmployeeChangePasswordResp
	query := `SELECT employee_password
	FROM employees
	WHERE employee_login = $1 AND deleted_at is null`

	err := c.db.QueryRow(ctx, query,
		pass.EmployeeLogin,
	).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("incorrect login")
		}
		log.Println("failed to get employee password from database", logger.Error(err))
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass.OldPassword))
	if err != nil {
		return nil, errors.New("password mismatch")
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed to generate employee new password", logger.Error(err))
		return nil, err
	}

	query = `UPDATE employees SET 
		employee_password = $1, 
		updated_at = NOW() 
	WHERE employee_login = $2 AND deleted_at is null`

	_, err = c.db.Exec(ctx, query, newHashedPassword, pass.EmployeeLogin)
	if err != nil {
		log.Println("failed to change employee password in database", logger.Error(err))
		return nil, err
	}
	resp.Comment = "Password changed successfully"

	return &resp, nil
}

func (c *employeeRepo) GetByLogin(ctx context.Context, login string) (*employee_service.GetEmployeeByLogin, error) {
	var (
		employee   employee_service.GetEmployeeByLogin
		birthday   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	query := `SELECT 
		id, 
		employee_login,
		birthday, 
		gender,
		fullname,
		email,
		phone,
		employee_password,
		created_at, 
		updated_at
		FROM employees WHERE employee_login = $1 AND deleted_at is null`

	row := c.db.QueryRow(ctx, query, login)

	err := row.Scan(
		&employee.Id,
		&employee.EmployeeLogin,
		&birthday,
		&employee.Gender,
		&employee.Fullname,
		&employee.Email,
		&employee.Phone,
		&employee.EmployeePassword,
		&created_at,
		&updated_at,
	)

	if err != nil {
		log.Println("failed to scan employee by LOGIN from database", logger.Error(err))
		return &employee_service.GetEmployeeByLogin{}, err
	}

	employee.Birthday = pkg.NullStringToString(birthday)
	employee.CreatedAt = pkg.NullStringToString(created_at)
	employee.UpdatedAt = pkg.NullStringToString(updated_at)

	return &employee, nil
}

func (c *employeeRepo) GetPassword(ctx context.Context, login string) (string, error) {
	var hashedPass string

	query := `SELECT employee_password
	FROM employees
	WHERE employee_login = $1 AND deleted_at is null`

	err := c.db.QueryRow(ctx, query, login).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("incorrect login")
		} else {
			log.Println("failed to get employee password from database", logger.Error(err))
			return "", err
		}
	}

	return hashedPass, nil
}
