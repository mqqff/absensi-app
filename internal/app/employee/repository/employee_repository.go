package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/pkg/helpers/pgerror"
)

type employeeRepository struct {
	conn *sqlx.DB
}

func NewEmployeeRepository(conn *sqlx.DB) contracts.EmployeeRepository {
	return &employeeRepository{
		conn: conn,
	}
}

func (r *employeeRepository) buildEmployeeFilter(query dto.EmployeesQuery) (string, []interface{}) {
	baseQuery := `
		FROM employees
		WHERE deleted_at IS NULL
	`

	var (
		conditions []string
		args       []interface{}
	)

	if query.Name != "" {
		args = append(args, "%"+query.Name+"%")
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", len(args)))
	}

	if query.Email != "" {
		args = append(args, "%"+query.Email+"%")
		conditions = append(conditions, fmt.Sprintf("email ILIKE $%d", len(args)))
	}

	if query.Position != 0 {
		args = append(args, query.Position)
		conditions = append(conditions, fmt.Sprintf("position = $%d", len(args)))
	}

	if query.Department != 0 {
		args = append(args, query.Department)
		conditions = append(conditions, fmt.Sprintf("department = $%d", len(args)))
	}

	if query.Status != 0 {
		args = append(args, query.Status)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	return baseQuery, args
}

func (r *employeeRepository) GetEmployees(
	ctx context.Context,
	query dto.EmployeesQuery,
	pagination dto.PaginationRequest,
) ([]entity.Employee, error) {

	var employees []entity.Employee

	baseQuery, args := r.buildEmployeeFilter(query)

	dataQuery := `
		SELECT id, name, email, phone, position, department, salary, address, status
	` + baseQuery + fmt.Sprintf(`
		ORDER BY name ASC
		LIMIT $%d OFFSET $%d
	`, len(args)+1, len(args)+2)

	args = append(args, pagination.Limit, pagination.Offset)

	if err := r.conn.SelectContext(ctx, &employees, dataQuery, args...); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *employeeRepository) CountEmployees(ctx context.Context, query dto.EmployeesQuery) (int, error) {
	baseQuery, args := r.buildEmployeeFilter(query)

	countQuery := `SELECT COUNT(*) ` + baseQuery

	var total int
	if err := r.conn.GetContext(ctx, &total, countQuery, args...); err != nil {
		return 0, err
	}

	return total, nil
}

func (r *employeeRepository) GetEmployeeByID(ctx context.Context, employeeID string) (entity.Employee, error) {
	var employee entity.Employee
	query := `SELECT id, name, email, phone, position, department, salary, address, status FROM employees WHERE id = $1`
	err := r.conn.GetContext(ctx, &employee, query, employeeID)
	if err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee entity.Employee) error {
	query := `
		INSERT INTO employees (id, name, email, phone, position, salary, address, status, password)
		VALUES (:id, :name, :email, :phone, :position, :salary, :address, :status, :password)
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

func (r *employeeRepository) UpdateEmployee(ctx context.Context, employee entity.Employee) error {
	query := `
		UPDATE employees
		SET name = :name, email = :email, phone = :phone, position = :position, salary = :salary, address = :address, status = :status
		WHERE id = :id
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
						WithLocation("employeeRepository.UpdateEmployee"),
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
