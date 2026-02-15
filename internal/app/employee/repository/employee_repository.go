package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/pkg/helpers/pgerror"
)

type employeeRepository struct {
	conn *sqlx.DB
	tx   *sqlx.Tx
}

func NewEmployeeRepository(conn *sqlx.DB) contracts.EmployeeRepository {
	return &employeeRepository{
		conn: conn,
	}
}

func (r *employeeRepository) Begin(ctx context.Context) error {
	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	r.tx = tx

	return nil
}

func (r *employeeRepository) GetTx() *sqlx.Tx {
	return r.tx
}

func (r *employeeRepository) UseTx(tx *sqlx.Tx) {
	r.tx = tx
}

func (r *employeeRepository) Commit() error {
	if r.tx == nil {
		return errors.New("transaction is nil")
	}

	err := r.tx.Commit()
	if err != nil {
		return err
	}

	r.tx = nil

	return nil
}

func (r *employeeRepository) Rollback() error {
	if r.tx == nil {
		return errors.New("transaction is nil")
	}

	err := r.tx.Rollback()
	if err != nil {
		return err
	}

	r.tx = nil

	return nil
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee entity.Employee) error {
	query := `
		INSERT INTO employees (id, name, email, phone, position, salary, address, password, status)
		VALUES (:id, :name, :email, :phone, :position, :salary, :address, :password, :status)
	`
	_, err := r.conn.NamedExecContext(ctx, query, employee)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "employees_email_key",
					Err: errx.ErrEmailAlreadyUsed.
						WithDetails(map[string]interface{}{
							"email": employee.Email,
						}).
						WithError(err).
						WithLocation("employeeRepository.CreateEmployee"),
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

func (r *employeeRepository) DeleteEmployee(ctx context.Context, employeeID string) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := r.conn.ExecContext(ctx, query, employeeID)
	if err != nil {
		return err
	}

	return nil
}
