package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/pkg/jwt"
)

func (m *Middleware) RequireAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header := ctx.Get("Authorization")
		if header == "" {
			return errx.ErrNoBearerToken
		}

		headerSlice := strings.Split(header, " ")
		if len(headerSlice) != 2 && headerSlice[0] != "Bearer" {
			return errx.ErrInvalidBearerToken
		}

		token := headerSlice[1]
		var claims jwt.Claims
		err := m.jwt.Decode(token, &claims)
		if err != nil {
			return errx.ErrInvalidBearerToken
		}

		notBefore, err := claims.GetNotBefore()
		if err != nil {
			return errx.ErrInvalidBearerToken
		}

		if notBefore.After(time.Now()) {
			return errx.ErrBearerTokenNotActive
		}

		expirationTime, err := claims.GetExpirationTime()
		if err != nil {
			return errx.ErrInvalidBearerToken
		}

		if expirationTime.Before(time.Now()) {
			return errx.ErrExpiredBearerToken
		}

		ctx.Locals("claims", claims)

		return ctx.Next()
	}
}
