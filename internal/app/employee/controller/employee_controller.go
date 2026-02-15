package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/internal/middlewares"
	"github.com/mqqff/absensi-app/pkg/helpers/http/response"
)

type employeeController struct {
	employeeService contracts.EmployeeService
}

func InitEmployeeController(
	router fiber.Router,
	employeeService contracts.EmployeeService,
	middleware *middlewares.Middleware,
) {
	employeeCtr := employeeController{
		employeeService: employeeService,
	}

	employeeGroup := router.Group("/employees", middleware.RequireAuth())
	employeeGroup.Post("/", employeeCtr.CreateEmployee)
	employeeGroup.Delete("/:id", employeeCtr.DeleteEmployee)
}

func (e *employeeController) CreateEmployee(ctx *fiber.Ctx) error {
	var req dto.CreateEmployeeRequest

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := e.employeeService.CreateEmployee(ctx.Context(), req); err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusCreated, nil)
}

func (e *employeeController) DeleteEmployee(ctx *fiber.Ctx) error {
	var req dto.DeleteEmployeeParam
	if err := ctx.ParamsParser(&req); err != nil {
		return err
	}

	if err := e.employeeService.DeleteEmployee(ctx.Context(), req); err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, nil)
}
