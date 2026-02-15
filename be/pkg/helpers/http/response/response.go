package response

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/internal/infra/env"
)

type Response struct {
	Payload interface{} `json:"payload"`
}

func SendResponse(
	ctx *fiber.Ctx,
	code int,
	payload any,
) error {
	if code >= 400 {
		if err, isError := payload.(error); isError {
			var reqErr *errx.RequestError
			if errors.As(err, &reqErr) {
				if env.AppEnv.AppEnv != "production" {
					payload = map[string]any{
						"error": map[string]any{
							"message":    reqErr.Message,
							"error":      reqErr.Err,
							"error_code": reqErr.ErrorCode,
							"location":   reqErr.Location,
							"details":    reqErr.Details,
						},
					}
				} else {
					payload = map[string]any{
						"error": reqErr,
					}
				}
			} else {
				payload = map[string]any{
					"error": map[string]any{
						"message": "unknown error",
						"error":   err.Error(),
					},
				}
			}
		}
	}

	return ctx.Status(code).JSON(
		Response{
			Payload: payload,
		},
	)
}
