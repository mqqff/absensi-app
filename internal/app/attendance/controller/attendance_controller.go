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
	attendanceGroup.Post("/", attendanceCtr.CreateAttendance)
}

func (c *attendanceController) CreateAttendance(ctx *fiber.Ctx) error {
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
