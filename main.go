package main

import (
	"echo-mangosteen/internal/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	pingCtrl := controller.NewPingController()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ping", pingCtrl.Ping)
	e.Logger.Fatal(e.Start(":1323"))
}
