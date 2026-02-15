package bootstrap

import (
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/mqqff/absensi-app/internal/infra/database"
	"github.com/mqqff/absensi-app/internal/infra/env"
	"github.com/mqqff/absensi-app/internal/infra/server"
)

func Init() {
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	server := server.NewServer()

	app := server.GetApp()

	app.Get("/metrics", monitor.New())

	server.MountMiddlewares()
	server.MountRoutes(psqlDB)
	server.Start(env.AppEnv.AppPort)
}
