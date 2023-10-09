package router

import (
	"echo-mangosteen/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	pingC controller.PingController,
	userC controller.UserController,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/ping", pingC.Ping)

	e.POST("/login", userC.Login)

	return e
}
