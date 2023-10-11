package middleware

import (
	"echo-mangosteen/pkg/errors"
	"echo-mangosteen/pkg/jwt"
	"echo-mangosteen/pkg/response"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(j *jwt.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return response.Build(c, errors.Unauthorized(), nil)
			}
			claims, err := j.ParseToken(tokenString)
			if err != nil {
				return response.Build(c, errors.Unauthorized(), nil)
			}
			c.Set("claims", claims)

			return next(c)
		}
	}
}
