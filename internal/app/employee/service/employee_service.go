package service

import (
	"context"

	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/enums"
	"github.com/mqqff/absensi-app/pkg/bcrypt"
	"github.com/mqqff/absensi-app/pkg/uuid"
	"github.com/mqqff/absensi-app/pkg/validator"
)

type employeeService struct {
	userRepo  contracts.EmployeeRepository
	validator validator.ValidatorInterface
	uuid      uuid.UUIDInterface
	bcrypt    bcrypt.BcryptInterface
}

func NewEmployeeService(
	userRepo contracts.EmployeeRepository,
	validator validator.ValidatorInterface,
	uuid uuid.UUIDInterface,
	bcrypt bcrypt.BcryptInterface,
) contracts.EmployeeService {
	return &employeeService{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
		bcrypt:    bcrypt,
	}
}

func (s *employeeService) CreateEmployee(ctx context.Context, data dto.CreateEmployeeRequest) error {
	valErr := s.validator.Validate(data)
	if valErr != nil {
		return valErr
	}

	hashedPassword, err := s.bcrypt.Hash(data.Password)

	if err != nil {
		return err
	}

	userID, _ := s.uuid.NewV7()

	employee := entity.Employee{
		ID:    userID,
		Name:  data.Name,
		Email: data.Email,
		Phone: data.Phone,
		Position: enums.NullEmployeePositionIdx{
			EmployeePositionIdx: data.Position,
			Valid:               true,
		},
		Department: enums.NullEmployeeDepartmentIdx{
			EmployeeDepartmentIdx: data.Department,
			Valid:                 true,
		},
		Salary:   data.Salary,
		Address:  data.Address,
		Password: hashedPassword,
		Status:   data.Status,
	}

	return s.userRepo.CreateEmployee(ctx, employee)
}
