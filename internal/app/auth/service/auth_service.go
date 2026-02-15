package service

import (
	"context"
	"errors"

	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/enums"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/pkg/bcrypt"
	"github.com/mqqff/absensi-app/pkg/jwt"
	"github.com/mqqff/absensi-app/pkg/uuid"
	"github.com/mqqff/absensi-app/pkg/validator"
)

type authService struct {
	authRepo  contracts.AuthRepository
	userRepo  contracts.EmployeeRepository
	validator validator.ValidatorInterface
	uuid      uuid.UUIDInterface
	jwt       jwt.CustomJwtInterface
	bcrypt    bcrypt.BcryptInterface
}

func NewAuthService(
	authRepo contracts.AuthRepository,
	userRepo contracts.EmployeeRepository,
	validator validator.ValidatorInterface,
	uuid uuid.UUIDInterface,
	jwt jwt.CustomJwtInterface,
	bcrypt bcrypt.BcryptInterface,
) contracts.AuthService {
	return &authService{
		authRepo:  authRepo,
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
		jwt:       jwt,
		bcrypt:    bcrypt,
	}
}

func (s *authService) LoginWithCredentials(
	ctx context.Context,
	req dto.LoginWithCredentialsRequest,
) (dto.LoginResponse, error) {
	valErr := s.validator.Validate(req)
	if valErr != nil {
		return dto.LoginResponse{}, valErr
	}

	employee, err := s.authRepo.GetEmployeeByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, errx.ErrEmployeeNotFound) {
			return dto.LoginResponse{}, errx.ErrCredentialsNotMatch
		}

		return dto.LoginResponse{}, err
	}

	if !s.bcrypt.Compare(req.Password, employee.Password) {
		return dto.LoginResponse{}, errx.ErrCredentialsNotMatch
	}

	position := enums.PositionUnknownIdx
	if employee.Position.Valid {
		position = employee.Position.EmployeePositionIdx
	}

	department := enums.DepartmentUnknownIdx
	if employee.Department.Valid {
		department = employee.Department.EmployeeDepartmentIdx
	}

	accessToken, err := s.jwt.Create(employee.ID, employee.Email, employee.Name, position, department)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	res := dto.LoginResponse{
		AccessToken: accessToken,
	}

	return res, nil
}

func (s *authService) GetSession(ctx context.Context, email string) (dto.GetSessionResponse, error) {
	user, err := s.authRepo.GetEmployeeByEmail(ctx, email)
	if err != nil {
		return dto.GetSessionResponse{}, err
	}

	res := dto.GetSessionResponse{
		ID:         user.ID.String(),
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Position:   user.Position.EmployeePositionIdx.String(),
		Department: user.Department.EmployeeDepartmentIdx.String(),
		Salary:     user.Salary,
		Address:    user.Address,
		Status:     user.Status.String(),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return res, nil
}
