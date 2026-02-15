package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
}
