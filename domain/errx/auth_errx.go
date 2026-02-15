package errx

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrEmployeeNotFound = NewError(
		fiber.StatusNotFound,
		"employee_not_found",
		"Akun tidak ditemukan",
	)
	ErrCredentialsNotMatch = NewError(
		fiber.StatusUnauthorized,
		"credentials_not_match",
		"Email atau password yang Anda masukkan salah",
	)
	ErrEmailAlreadyUsed = NewError(
		fiber.StatusConflict,
		"email_already_used",
		"Email sudah digunakan oleh akun lain",
	)
)
