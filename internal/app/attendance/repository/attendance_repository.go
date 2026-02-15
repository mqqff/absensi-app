package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/pkg/helpers/pgerror"
)

type attendanceRepository struct {
	conn *sqlx.DB
}

func NewAttendanceRepository(conn *sqlx.DB) contracts.AttendanceRepository {
	return &attendanceRepository{
		conn: conn,
	}
}

func (s *attendanceRepository) CheckIn(ctx context.Context, attendance entity.Attendance) error {
	query := `INSERT INTO attendances (id, employee_id, check_in, status) VALUES ($1, $2, $3, $4)`
	_, err := s.conn.Exec(query, attendance.ID, attendance.EmployeeID, attendance.CheckIn, attendance.Status)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "unique_employee_daily",
					Err: errx.ErrAlreadyCheckedInToday.
						WithError(err).
						WithLocation("attendanceRepository.CheckIn"),
				},
			}

			if customPgErr := pgerror.HandlePgError(*pgErr, pgErrors); customPgErr != nil {
				return customPgErr
			}
		}

		return err
	}

	return nil
}

func (s *attendanceRepository) GetOpenAttendanceByEmployeeID(ctx context.Context, employeeID uuid.UUID) (entity.Attendance, error) {
	var attendance entity.Attendance

	query := `SELECT id, check_in, status FROM attendances WHERE employee_id = $1 AND check_out IS NULL`
	err := s.conn.GetContext(ctx, &attendance, query, employeeID)
	if err != nil {
		return entity.Attendance{}, err
	}

	return attendance, nil
}

func (s *attendanceRepository) CheckOut(ctx context.Context, attendance entity.Attendance) error {
	query := `UPDATE attendances SET check_out = $1, status = $2 WHERE id = $3`
	_, err := s.conn.Exec(query, attendance.CheckOut, attendance.Status, attendance.ID)
	if err != nil {
		return err
	}

	return nil
}
