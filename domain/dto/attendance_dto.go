package dto

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceResponse struct {
	ID         uuid.UUID  `json:"id"`
	EmployeeID uuid.UUID  `json:"employee_id"`
	CheckIn    time.Time  `json:"check_in"`
	CheckOut   *time.Time `json:"check_out,omitempty"`
}

type CreateAttendanceRequest struct {
	EmployeeID uuid.UUID `json:"employee_id"`
}
