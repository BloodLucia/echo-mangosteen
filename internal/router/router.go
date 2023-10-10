package router

import (
	"echo-mangosteen/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	pingCtrl controller.PingController,
	userCtrl controller.UserController,
	codeCtrl controller.CodeController,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/ping", pingCtrl.Ping)

	e.POST("/login", userCtrl.Login)

	e.POST("/send_validation_code", codeCtrl.SendValidationCode)

	return e
}
