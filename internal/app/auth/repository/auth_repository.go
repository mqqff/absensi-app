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

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}
	query := `SELECT * FROM users WHERE email = $1`
	err := r.conn.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errx.ErrCredentialsNotMatch.
				WithDetails(map[string]interface{}{
					"email": email,
				}).
				WithLocation("authRepository.GetUserByEmail")
		}

		return entity.User{}, err
	}

	return user, nil
}
