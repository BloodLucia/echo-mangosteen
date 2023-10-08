package router

import (
	"echo-mangosteen/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(
	pingC controller.PingController,
) *echo.Echo {
	e := echo.New()


	e.GET("/ping", pingC.Ping)


	return e
}
