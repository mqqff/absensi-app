package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/mqqff/absensi-app/domain/enums"
)

type Attendance struct {
	ID         uuid.UUID              `db:"id" json:"id"`
	EmployeeID uuid.UUID              `db:"employee_id" json:"employee_id"`
	Employee   Employee               `db:"employee" json:"employee"`
	CheckIn    time.Time              `db:"check_in" json:"check_in"`
	CheckOut   *time.Time             `db:"check_out" json:"check_out"`
	Status     enums.AttendanceStatus `db:"status" json:"status"`
}
