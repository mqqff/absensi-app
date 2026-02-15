package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
)

type EmployeeRepository interface {
	GetEmployees(ctx context.Context, query dto.EmployeesQuery, pagination dto.PaginationRequest) ([]entity.Employee, error)
	CountEmployees(ctx context.Context, query dto.EmployeesQuery) (int, error)
	GetEmployeeByID(ctx context.Context, employeeID string) (entity.Employee, error)
	CreateEmployee(ctx context.Context, employee entity.Employee) error
	UpdateEmployee(ctx context.Context, employee entity.Employee) error
	DeleteEmployee(ctx context.Context, employeeID string) error
}

type EmployeeService interface {
	GetEmployees(ctx context.Context, query dto.GetEmployeesQuery) (dto.GetEmployeesResponse, error)
	GetEmployee(ctx context.Context, param dto.GetEmployee) (dto.EmployeeResponse, error)
	CreateEmployee(ctx context.Context, data dto.CreateEmployeeRequest) error
	UpdateEmployee(ctx context.Context, data dto.UpdateEmployeeRequest) error
	DeleteEmployee(ctx context.Context, param dto.DeleteEmployeeParam) error
}
