package jwt

import (
	"time"

	"github.com/mqqff/absensi-app/domain/enums"
	"github.com/mqqff/absensi-app/internal/infra/env"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomJwtInterface interface {
	Create(employeeID uuid.UUID, name string, email string, position enums.EmployeePositionIdx, department enums.EmployeeDepartmentIdx) (string, error)
	Decode(tokenString string, claims *Claims) error
}

type Claims struct {
	jwt.RegisteredClaims
	EmployeeID uuid.UUID                   `json:"employee_id"`
	Name       string                      `json:"name"`
	Email      string                      `json:"email"`
	Position   enums.EmployeePositionIdx   `json:"position"`
	Department enums.EmployeeDepartmentIdx `json:"department"`
}

type CustomJwtStruct struct {
	SecretKey   string
	ExpiredTime time.Duration
}

var Jwt = getJwt()

func getJwt() CustomJwtInterface {
	return &CustomJwtStruct{
		SecretKey:   env.AppEnv.JwtSecretKey,
		ExpiredTime: env.AppEnv.JwtExpTime,
	}
}

func (j *CustomJwtStruct) Create(employeeID uuid.UUID, email, name string, position enums.EmployeePositionIdx, department enums.EmployeeDepartmentIdx) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "absensi-app",
			Audience:  jwt.ClaimStrings{"absensi-app"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
		EmployeeID: employeeID,
		Name:       name,
		Email:      email,
		Position:   position,
		Department: department,
	}

	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJWT, err := unsignedJWT.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func (j *CustomJwtStruct) Decode(tokenString string, claims *Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (any, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	return nil
}
