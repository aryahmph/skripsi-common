package jwt

import "github.com/golang-jwt/jwt/v5"

type (
	CustomClaims struct {
		ID   string `json:"id"`
		Role string `json:"role"`
		jwt.RegisteredClaims
	}
)
