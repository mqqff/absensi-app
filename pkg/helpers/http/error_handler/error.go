package error_handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/pkg/helpers/http/response"
	"github.com/mqqff/absensi-app/pkg/log"
	"github.com/mqqff/absensi-app/pkg/validator"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var valErr validator.ValidationErrors
	if errors.As(err, &valErr) {
		return response.SendResponse(c, fiber.StatusUnprocessableEntity, map[string]any{
			"message":    "Validation error",
			"error":      valErr,
			"error_code": "validation_error",
		})
	}

	var reqErr *errx.RequestError
	if errors.As(err, &reqErr) {
		log.Error(log.LogInfo{
			"error_code": reqErr.ErrorCode,
			"location":   reqErr.Location,
			"details":    reqErr.Details,
			"error":      reqErr.Err,
		}, "[ErrorHandler] Request error")

		return response.SendResponse(c, reqErr.StatusCode, reqErr)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return response.SendResponse(c, fiberErr.Code, fiber.Map{})
	}

	return response.SendResponse(c, fiber.StatusInternalServerError, errx.ErrInternalServer.WithError(err))
}
