package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
)

type AttendanceRepository interface {
	CheckIn(ctx context.Context, attendance entity.Attendance) error
}

type AttendanceService interface {
	CheckIn(ctx context.Context, data dto.CheckInRequest) (dto.AttendanceResponse, error)
}
