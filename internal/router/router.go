package router

import (
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/pkg/jwt"
	"echo-mangosteen/pkg/middleware"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	jwt *jwt.JWT,
	pingCtrl controller.PingController,
	userCtrl controller.UserController,
	codeCtrl controller.CodeController,
	tagCtrl controller.TagController,
) *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.Logger())

	v1NoAuth := e.Group("/api/v1")
	{
		v1NoAuth.GET("/ping", pingCtrl.Ping)
		v1NoAuth.POST("/login", userCtrl.Login)
	}

	v1Auth := e.Group("/api/v1")
	v1Auth.Use(middleware.JWTMiddleware(jwt))
	{
		v1Auth.POST("/tags/add", tagCtrl.AddTag)
		v1Auth.DELETE("/tags/delete/:tagId", tagCtrl.DeleteTag)
	}

	return e
}
