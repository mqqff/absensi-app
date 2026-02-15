package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/errx"
)

type authRepository struct {
	conn *sqlx.DB
}

func NewAuthRepository(conn *sqlx.DB) contracts.AuthRepository {
	return &authRepository{
		conn: conn,
	}
}

func (r *authRepository) GetEmployeeByEmail(ctx context.Context, email string) (entity.Employee, error) {
	user := entity.Employee{}
	query := `SELECT * FROM employees WHERE email = $1`
	err := r.conn.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Employee{}, errx.ErrEmployeeNotFound.
				WithDetails(map[string]interface{}{
					"email": email,
				}).
				WithLocation("authRepository.GetEmployeeByEmail")
		}

		return entity.Employee{}, err
	}

	return user, nil
}
