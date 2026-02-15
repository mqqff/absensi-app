package contracts

import (
	"context"

	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
)

type AuthRepository interface {
	GetEmployeeByEmail(ctx context.Context, email string) (entity.Employee, error)
}

type AuthService interface {
	LoginWithCredentials(ctx context.Context, req dto.LoginWithCredentialsRequest) (dto.LoginResponse, error)
	GetSession(ctx context.Context, nim string) (dto.GetSessionResponse, error)
}
