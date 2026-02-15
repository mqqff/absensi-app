package errx

import "github.com/gofiber/fiber/v2"

var (
	ErrAttendanceOpenNotFound = NewError(
		fiber.StatusNotFound,
		"attendance_open_not_found",
		"sudah check-out atau belum check-in",
	)
	ErrAlreadyCheckedInToday = NewError(
		fiber.StatusForbidden,
		"already_checked_in",
		"kamu sudah check-in hari ini",
	)
	ErrPendingCheckout = NewError(
		fiber.StatusForbidden,
		"pending_checkout",
		"ada absensi yang belum check-out, hubungi HRD untuk menyelesaikan absensi tersebut",
	)
)
