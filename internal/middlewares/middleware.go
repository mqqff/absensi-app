package middlewares

import "github.com/mqqff/absensi-app/pkg/jwt"

type Middleware struct {
	jwt jwt.CustomJwtInterface
}

func NewMiddleware(jwt jwt.CustomJwtInterface) *Middleware {
	return &Middleware{jwt}
}
