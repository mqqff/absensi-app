package service

import (
	"context"
	"time"

	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/pkg/uuid"
	"github.com/mqqff/absensi-app/pkg/validator"
)

type attendanceService struct {
	attendanceRepo contracts.AttendanceRepository
	validator      validator.ValidatorInterface
	uuid           uuid.UUIDInterface
}

func NewAttendanceService(
	attendanceRepo contracts.AttendanceRepository,
	validator validator.ValidatorInterface,
	uuid uuid.UUIDInterface,
) contracts.AttendanceService {
	return &attendanceService{
		attendanceRepo, validator, uuid,
	}
}

func (s *attendanceService) CheckIn(ctx context.Context, data dto.CheckInRequest) (dto.AttendanceResponse, error) {
	attendanceID, _ := s.uuid.NewV7()

	attendance := entity.Attendance{
		ID:         attendanceID,
		EmployeeID: data.EmployeeID,
		CheckIn:    time.Now(),
	}

	err := s.attendanceRepo.CheckIn(ctx, attendance)
	if err != nil {
		return dto.AttendanceResponse{}, err
	}

	return dto.AttendanceResponse{
		ID:         attendance.ID,
		EmployeeID: attendance.EmployeeID,
		CheckIn:    attendance.CheckIn,
	}, nil
}
