package dto

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/enums"
)

type EmployeeResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Position   string    `json:"position"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	Address    string    `json:"address"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetEmployeeResponse struct {
	Employees []EmployeeResponse `json:"employees"`
}

type GetEmployee struct {
	ID string `param:"id" validate:"required,uuid"`
}

type CreateEmployeeRequest struct {
	Name       string                      `json:"name" validate:"required"`
	Email      string                      `json:"email" validate:"required,email"`
	Phone      string                      `json:"phone" validate:"required"`
	Salary     float64                     `json:"salary" validate:"required,gt=0"`
	Password   string                      `json:"password" validate:"required,min=6"`
	Address    string                      `json:"address" validate:"required"`
	Position   enums.EmployeePositionIdx   `json:"position" validate:"required,oneof=1 2 3 4"`
	Department enums.EmployeeDepartmentIdx `json:"department" validate:"required,oneof=1 2 3 4 5"`
	Status     enums.EmployeeStatus        `json:"status"`
}

type UpdateEmployeeRequest struct {
	ID         uuid.UUID                   `param:"id" validate:"required,uuid"`
	Name       string                      `json:"name" validate:"required"`
	Email      string                      `json:"email" validate:"required,email"`
	Phone      string                      `json:"phone" validate:"required"`
	Salary     float64                     `json:"salary" validate:"required,gt=0"`
	Password   string                      `json:"password" validate:"omitempty,min=6"`
	Address    string                      `json:"address" validate:"required"`
	Position   enums.EmployeePositionIdx   `json:"position" validate:"required,oneof=1 2 3 4"`
	Department enums.EmployeeDepartmentIdx `json:"department" validate:"required,oneof=1 2 3 4 5"`
	Status     enums.EmployeeStatus        `json:"status"`
}

type DeleteEmployeeParam struct {
	ID string `param:"id" validate:"required,uuid"`
}

func FormatToEmployeeResponse(employee entity.Employee) EmployeeResponse {
	fmt.Println(employee)
	var position string
	if employee.Position.Valid {
		position = enums.EmployeePositionMapIdx[employee.Position.EmployeePositionIdx].LongLabel["id"]
	}

	var department string
	if employee.Department.Valid {
		department = enums.EmployeeDepartmentMapIdx[employee.Department.EmployeeDepartmentIdx].LongLabel["id"]
	}

	return EmployeeResponse{
		ID:         employee.ID.String(),
		Name:       employee.Name,
		Email:      employee.Email,
		Phone:      employee.Phone,
		Position:   position,
		Department: department,
		Salary:     employee.Salary,
		Address:    employee.Address,
		Status:     enums.EmployeeStatusMap[employee.Status],
		CreatedAt:  employee.CreatedAt,
		UpdatedAt:  employee.UpdatedAt,
	}
}
