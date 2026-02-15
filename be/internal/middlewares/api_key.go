package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/internal/infra/env"
)

func APIKey() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		apiKey := ctx.Get("x-api-key")
		if apiKey == "" {
			return errx.ErrNoAPIKey
		}

		keySlice := strings.Split(apiKey, " ")
		if len(keySlice) != 2 {
			return errx.ErrInvalidAPIKey
		}

		key := keySlice[1]
		if key != env.AppEnv.ApiKey {
			return errx.ErrInvalidAPIKey
		}

		return ctx.Next()
	}
}
