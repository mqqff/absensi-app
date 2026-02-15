package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/pkg/log"
)

func (m *Middleware) Limit(maxCount int, duration string) fiber.Handler {
	d, _ := time.ParseDuration(duration)

	return limiter.New(limiter.Config{
		Max:        maxCount,
		Expiration: d,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.Method() + "-" + ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			log.Info(log.LogInfo{
				"max":      maxCount,
				"duration": duration,
				"ip":       ctx.IP(),
				"method":   ctx.Method(),
				"path":     ctx.Path(),
			}, "[LIMITER][Limit] Limit reached")
			return errx.ErrTooManyRequests
		},
	})
}
