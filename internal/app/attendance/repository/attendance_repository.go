package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/entity"
)

type attendanceRepository struct {
	conn *sqlx.DB
}

func NewAttendanceRepository(conn *sqlx.DB) contracts.AttendanceRepository {
	return &attendanceRepository{
		conn: conn,
	}
}

func (s *attendanceRepository) CreateAttendance(ctx context.Context, attendance entity.Attendance) error {
	query := `INSERT INTO attendances (id, employee_id, check_in) VALUES ($1, $2, $3)`
	_, err := s.conn.Exec(query, attendance.ID, attendance.EmployeeID, time.Now())
	if err != nil {
		return err
	}

	return nil
}
