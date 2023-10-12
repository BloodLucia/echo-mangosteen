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

	v1Auth := e.Group("/api/v1")
	v1Auth.Use(middleware.JWTAuth(jwt))
	{
		v1Auth.POST("/tags/add", tagCtrl.AddTag)
		v1Auth.DELETE("/tags/delete/:tagId", tagCtrl.DeleteTag)

		// 不需要鉴权
		v1Auth.GET("/ping", pingCtrl.Ping)
		v1Auth.POST("/login", userCtrl.Login)
	}

	return e
}
