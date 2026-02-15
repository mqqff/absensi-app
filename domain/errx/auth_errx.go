package errx

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrUserNotFound = NewError(
		fiber.StatusNotFound,
		"user_not_found",
		"We couldn't find that account",
	)
	ErrCredentialsNotMatch = NewError(
		fiber.StatusUnauthorized,
		"credentials_not_match",
		"Incorrect email or password. Please try again.",
	)
	ErrEmailAlreadyUsed = NewError(
		fiber.StatusConflict,
		"email_already_used",
		"The email address is already in use by another account",
	)
)
