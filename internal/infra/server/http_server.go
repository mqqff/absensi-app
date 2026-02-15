package server

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mqqff/absensi-app/domain/errx"
	employeeRepo "github.com/mqqff/absensi-app/internal/app/employee/repository"
	"github.com/mqqff/absensi-app/pkg/uuid"
	"github.com/mqqff/absensi-app/pkg/validator"

	"github.com/mqqff/absensi-app/internal/infra/env"
	"github.com/mqqff/absensi-app/internal/middlewares"
	"github.com/mqqff/absensi-app/pkg/bcrypt"
	errorHandler "github.com/mqqff/absensi-app/pkg/helpers/http/error_handler"
	"github.com/mqqff/absensi-app/pkg/jwt"
	// s3 "github.com/mqqff/absensi-app/pkg/s3"

	authCtr "github.com/mqqff/absensi-app/internal/app/auth/controller"
	authRepo "github.com/mqqff/absensi-app/internal/app/auth/repository"
	authSvc "github.com/mqqff/absensi-app/internal/app/auth/service"
)

type HTTPServer interface {
	MountMiddlewares()
	MountRoutes(db *sqlx.DB)
	Start(part string)
	GetApp() *fiber.App
}

type httpServer struct {
	app *fiber.App
}

func NewServer() HTTPServer {
	config := fiber.Config{
		BodyLimit:     50 * 1024 * 1024, // 50 MiB
		CaseSensitive: true,
		AppName:       env.AppEnv.AppName,
		ServerHeader:  env.AppEnv.AppName,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		ErrorHandler:  errorHandler.ErrorHandler,
	}

	app := fiber.New(config)

	return &httpServer{
		app: app,
	}
}

func (s *httpServer) GetApp() *fiber.App {
	return s.app
}

func (s *httpServer) Start(part string) {
	if part[0] != ':' {
		part = ":" + part
	}

	err := s.app.Listen(part)

	if err != nil {
		panic(err)
	}
}

func (s *httpServer) MountMiddlewares() {
	s.app.Use(middlewares.Cors())
	s.app.Use(middlewares.Recover())
	s.app.Use(middlewares.Logger())
	s.app.Use(middlewares.Helmet())
	s.app.Use(middlewares.Compress())
}

func (s *httpServer) MountRoutes(db *sqlx.DB) {
	bcrypt := bcrypt.Bcrypt
	uuid := uuid.UUID
	validator := validator.Validator
	jwt := jwt.Jwt

	middleware := middlewares.NewMiddleware(jwt)

	s.app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Halo bang!")
	})

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	s.app.Use(middlewares.APIKey())

	// Repositories
	authRepository := authRepo.NewAuthRepository(db)
	employeeRepository := employeeRepo.NewEmployeeRepository(db)

	// Services
	authService := authSvc.NewAuthService(authRepository, employeeRepository, validator, uuid, jwt, bcrypt)

	// Controllers
	authCtr.InitAuthController(v1, authService, middleware)

	s.app.Use(func(ctx *fiber.Ctx) error {
		return errx.ErrNotFound
	})
}
