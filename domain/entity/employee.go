package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/mqqff/absensi-app/domain/enums"
)

type Employee struct {
	ID         uuid.UUID                   `db:"id" json:"id"`
	Name       string                      `db:"name" json:"name"`
	Email      string                      `db:"email" json:"email"`
	Phone      string                      `db:"phone" json:"phone"`
	Position   enums.EmployeePositionIdx   `db:"position" json:"position"`
	Department enums.EmployeeDepartmentIdx `db:"department" json:"department"`
	Salary     float64                     `db:"salary" json:"salary"`
	Address    string                      `db:"address" json:"address"`
	Password   string                      `db:"password" json:"-"`
	Status     enums.EmployeeStatus        `db:"status" json:"status"`
	CreatedAt  time.Time                   `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time                   `db:"updated_at" json:"updated_at"`
	DeletedAt  sql.NullTime                `db:"deleted_at" json:"deleted_at"`
}
