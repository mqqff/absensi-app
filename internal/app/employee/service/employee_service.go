package service

import (
	"context"

	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/entity"
	"github.com/mqqff/absensi-app/domain/enums"
	"github.com/mqqff/absensi-app/pkg/bcrypt"
	customUUID "github.com/mqqff/absensi-app/pkg/uuid"
	"github.com/mqqff/absensi-app/pkg/validator"
)

type employeeService struct {
	userRepo  contracts.EmployeeRepository
	validator validator.ValidatorInterface
	uuid      customUUID.UUIDInterface
	bcrypt    bcrypt.BcryptInterface
}

func NewEmployeeService(
	userRepo contracts.EmployeeRepository,
	validator validator.ValidatorInterface,
	uuid customUUID.UUIDInterface,
	bcrypt bcrypt.BcryptInterface,
) contracts.EmployeeService {
	return &employeeService{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
		bcrypt:    bcrypt,
	}
}

func (s *employeeService) GetEmployee(ctx context.Context, param dto.GetEmployee) (dto.EmployeeResponse, error) {
	valErr := s.validator.Validate(param)
	if valErr != nil {
		return dto.EmployeeResponse{}, valErr
	}

	employee, err := s.userRepo.GetEmployeeByID(ctx, param.ID)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	response := dto.FormatToEmployeeResponse(employee)

	return response, nil
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

func (s *employeeService) UpdateEmployee(ctx context.Context, data dto.UpdateEmployeeRequest) error {
	valErr := s.validator.Validate(data)
	if valErr != nil {
		return valErr
	}

	employee := entity.Employee{
		ID:    data.ID,
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
		Salary:  data.Salary,
		Address: data.Address,
		Status:  data.Status,
	}

	return s.userRepo.UpdateEmployee(ctx, employee)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, param dto.DeleteEmployeeParam) error {
	return s.userRepo.DeleteEmployee(ctx, param.ID)
}
