package http

import (
	errorCommon "github.com/aryahmph/skripsi-common/error"
	jwtCommon "github.com/aryahmph/skripsi-common/jwt"
	uModel "github.com/aryahmph/skripsi-common/model/user"
	"github.com/labstack/echo/v4"
)

func Auth(jwtManager *jwtCommon.JWTManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			a := c.Request().Header.Get("Authorization")
			if len(a) <= BEARER {
				return errorCommon.NewInvariantError("Authorization header not valid")
			}
			idToken := a[BEARER:]

			// verify token is real from user
			uid, role, err := jwtManager.Verify(idToken)
			if err != nil {
				return errorCommon.NewUnauthorizedError("Token not valid")
			}

			c.Set(AUTH_USER, uModel.AuthUser{
				ID:   uid,
				Role: role,
			})
			return next(c)
		}
	}
}
