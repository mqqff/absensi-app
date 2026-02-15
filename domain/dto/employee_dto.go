package dto

import "github.com/mqqff/absensi-app/domain/enums"

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

type DeleteEmployeeParam struct {
	ID string `param:"id" validate:"required,uuid"`
}
