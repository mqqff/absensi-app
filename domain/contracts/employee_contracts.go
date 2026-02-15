package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/entity"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee entity.Employee) error
}
