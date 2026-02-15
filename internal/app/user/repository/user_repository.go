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

type userRepository struct {
	conn *sqlx.DB
	tx   *sqlx.Tx
}

func NewUserRepository(conn *sqlx.DB) contracts.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (r *userRepository) Begin(ctx context.Context) error {
	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	r.tx = tx

	return nil
}

func (r *userRepository) GetTx() *sqlx.Tx {
	return r.tx
}

func (r *userRepository) UseTx(tx *sqlx.Tx) {
	r.tx = tx
}

func (r *userRepository) Commit() error {
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

func (r *userRepository) Rollback() error {
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

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) error {
	query := `
		INSERT INTO users (id, name, email, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.conn.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErrors := []pgerror.PgError{
				{
					Code:           pgerror.UniqueViolation,
					ConstraintName: "users_email_key",
					Err: errx.ErrEmailAlreadyUsed.
						WithDetails(map[string]interface{}{
							"email": user.Email,
						}).
						WithError(err).
						WithLocation("userRepository.CreateUser"),
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
