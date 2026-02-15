package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mqqff/absensi-app/domain/contracts"
	"github.com/mqqff/absensi-app/domain/dto"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/internal/middlewares"
	"github.com/mqqff/absensi-app/pkg/helpers/http/response"
	"github.com/mqqff/absensi-app/pkg/jwt"
)

type attendanceController struct {
	attendanceService contracts.AttendanceService
}

func InitAttendanceController(
	router fiber.Router,
	attendanceService contracts.AttendanceService,
	middleware *middlewares.Middleware,
) {
	attendanceCtr := attendanceController{
		attendanceService: attendanceService,
	}

	attendanceGroup := router.Group("/attendances", middleware.RequireAuth())
	attendanceGroup.Get("/open", attendanceCtr.GetOpenAttendance)
	attendanceGroup.Post("/checkin", attendanceCtr.CheckIn)
	attendanceGroup.Patch("/checkout", attendanceCtr.CheckOut)
}

func (c *attendanceController) GetOpenAttendance(ctx *fiber.Ctx) error {
	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrNoBearerToken
	}

	res, err := c.attendanceService.GetOpenAttendance(ctx.Context(), claims.EmployeeID)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (c *attendanceController) CheckIn(ctx *fiber.Ctx) error {
	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrNoBearerToken
	}

	req := dto.CheckInRequest{
		EmployeeID: claims.EmployeeID,
	}

	res, err := c.attendanceService.CheckIn(ctx.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusCreated, res)
}

func (c *attendanceController) CheckOut(ctx *fiber.Ctx) error {
	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrNoBearerToken
	}

	req := dto.CheckOutRequest{
		EmployeeID: claims.EmployeeID,
	}

	res, err := c.attendanceService.CheckOut(ctx.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}
