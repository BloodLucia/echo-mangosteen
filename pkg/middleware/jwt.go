package middleware

import (
	"echo-mangosteen/pkg/errors"
	"echo-mangosteen/pkg/jwt"
	"echo-mangosteen/pkg/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTAuth(j *jwt.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			whiteList := []string{"/login", "/ping"}
			requestPATH := c.Request().URL.Path
			if contains(whiteList, strings.TrimPrefix(requestPATH, "/api/v1")) {
				return next(c)
			}
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

func contains(slice []string, element interface{}) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
