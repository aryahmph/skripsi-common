package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func New(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) Verify(tokenString string) (id, role string, err error) {
	claims := &CustomClaims{}
	if _, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	}); err != nil {
		return "", "", err
	}
	return claims.ID, claims.Role, err
}

func (j JWTManager) Generate(id, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		id,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
	})
	return token.SignedString(j.AccessTokenKey)
}
