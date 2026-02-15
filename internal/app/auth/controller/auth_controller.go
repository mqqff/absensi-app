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

type authController struct {
	authService contracts.AuthService
}

func InitAuthController(
	router fiber.Router,
	authService contracts.AuthService,
	middleware *middlewares.Middleware,
) {
	authCtr := authController{
		authService: authService,
	}

	authGroup := router.Group("/auth")
	authGroup.Get("/session", middleware.RequireAuth(), authCtr.GetSession)
	authGroup.Post("/login", authCtr.Login)
}

func (authCtr *authController) GetSession(ctx *fiber.Ctx) error {
	claims, ok := ctx.Locals("claims").(jwt.Claims)
	if !ok {
		return errx.ErrClaimsNotFound
	}

	res, err := authCtr.authService.GetSession(ctx.Context(), claims.Email)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (authCtr *authController) Login(ctx *fiber.Ctx) error {
	var loginReq dto.LoginWithCredentialsRequest
	if err := ctx.BodyParser(&loginReq); err != nil {
		return err
	}

	res, err := authCtr.authService.LoginWithCredentials(ctx.Context(), loginReq)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}
