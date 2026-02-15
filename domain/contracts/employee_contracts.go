package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee entity.Employee) error
}

type EmployeeService interface {
	CreateEmployee(ctx context.Context, data dto.CreateEmployeeRequest) error
}
