package contracts

import (
	"context"

	"github.com/google/uuid"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
)

type AttendanceRepository interface {
	GetOpenAttendanceByEmployeeID(ctx context.Context, employeeID uuid.UUID) (entity.Attendance, error)
	CheckIn(ctx context.Context, attendance entity.Attendance) error
	CheckOut(ctx context.Context, attendance entity.Attendance) error
}

type AttendanceService interface {
	GetOpenAttendance(ctx context.Context, employeeID uuid.UUID) (dto.AttendanceResponse, error)
	CheckIn(ctx context.Context, data dto.CheckInRequest) (dto.AttendanceResponse, error)
	CheckOut(ctx context.Context, data dto.CheckOutRequest) (dto.AttendanceResponse, error)
}
