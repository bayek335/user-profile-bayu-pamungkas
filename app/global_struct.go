package app

import "github.com/golang-jwt/jwt/v5"

type SendEmail struct {
	Email string
	Code  string
}

type ValidateCode struct {
	Code string `json:"verification_code" binding:"required,min=6"`
}

// Other response
type ResponseWithOutData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Jwt claims
type JwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
