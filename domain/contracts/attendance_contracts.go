package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
)

type AttendanceRepository interface {
	CreateAttendance(ctx context.Context, attendance entity.Attendance) error
}

type AttendanceService interface {
	CreateAttendance(ctx context.Context, data dto.CreateAttendanceRequest) (dto.AttendanceResponse, error)
}
