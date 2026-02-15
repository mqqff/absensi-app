package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee entity.Employee) error
	UpdateEmployee(ctx context.Context, employee entity.Employee) error
	DeleteEmployee(ctx context.Context, employeeID string) error
}

type EmployeeService interface {
	CreateEmployee(ctx context.Context, data dto.CreateEmployeeRequest) error
	UpdateEmployee(ctx context.Context, data dto.UpdateEmployeeRequest) error
	DeleteEmployee(ctx context.Context, param dto.DeleteEmployeeParam) error
}
