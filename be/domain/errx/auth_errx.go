package errx

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrEmployeeNotFound = NewError(
		fiber.StatusNotFound,
		"employee_not_found",
		"akun tidak ditemukan",
	)
	ErrCredentialsNotMatch = NewError(
		fiber.StatusUnauthorized,
		"credentials_not_match",
		"email atau password yang Anda masukkan salah",
	)
	ErrEmailAlreadyUsed = NewError(
		fiber.StatusConflict,
		"email_already_used",
		"email sudah digunakan oleh akun lain",
	)
	ErrPhoneAlreadyUsed = NewError(
		fiber.StatusConflict,
		"phone_already_used",
		"nomor telepon sudah digunakan oleh akun lain",
	)
)
