package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/enums"
	"github.com/mqqff/absensi-app/domain/errx"
	customUUID "github.com/mqqff/absensi-app/pkg/uuid"
	"github.com/mqqff/absensi-app/pkg/validator"
)

type attendanceService struct {
	attendanceRepo contracts.AttendanceRepository
	validator      validator.ValidatorInterface
	uuid           customUUID.UUIDInterface
}

func NewAttendanceService(
	attendanceRepo contracts.AttendanceRepository,
	validator validator.ValidatorInterface,
	uuid customUUID.UUIDInterface,
) contracts.AttendanceService {
	return &attendanceService{
		attendanceRepo, validator, uuid,
	}
}

func (s *attendanceService) CheckIn(ctx context.Context, data dto.CheckInRequest) (dto.AttendanceResponse, error) {
	attendance, err := s.attendanceRepo.GetOpenAttendanceByEmployeeID(ctx, data.EmployeeID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return dto.AttendanceResponse{}, err
		}
	}

	if attendance.ID != uuid.Nil {
		if time.Now().Day() != attendance.CheckIn.Day() {
			return dto.AttendanceResponse{}, errx.ErrPendingCheckout
		}
	}

	attendanceID, _ := s.uuid.NewV7()

	attendance = entity.Attendance{
		ID:         attendanceID,
		EmployeeID: data.EmployeeID,
		CheckIn:    time.Now(),
		Status:     enums.OnTime,
	}

	lateCheckInTime := time.Date(attendance.CheckIn.Year(), attendance.CheckIn.Month(), attendance.CheckIn.Day(), 8, 0, 0, 0, attendance.CheckIn.Location())
	if attendance.CheckIn.After(lateCheckInTime) {
		attendance.Status = enums.Late
	}

	err = s.attendanceRepo.CheckIn(ctx, attendance)
	if err != nil {
		return dto.AttendanceResponse{}, err
	}

	return dto.AttendanceResponse{
		ID:      attendance.ID,
		CheckIn: attendance.CheckIn,
		Status:  attendance.Status.String(),
	}, nil
}

func (s *attendanceService) CheckOut(ctx context.Context, data dto.CheckOutRequest) (dto.AttendanceResponse, error) {
	attendance, err := s.attendanceRepo.GetOpenAttendanceByEmployeeID(ctx, data.EmployeeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.AttendanceResponse{}, errx.ErrAttendanceOpenNotFound
		}

		return dto.AttendanceResponse{}, err
	}

	checkoutTime := time.Now()
	attendance.CheckOut = &checkoutTime

	err = s.attendanceRepo.CheckOut(ctx, attendance)
	if err != nil {
		return dto.AttendanceResponse{}, err
	}

	return dto.AttendanceResponse{
		ID:       attendance.ID,
		CheckIn:  attendance.CheckIn,
		CheckOut: attendance.CheckOut,
		Status:   attendance.Status.String(),
	}, nil
}
