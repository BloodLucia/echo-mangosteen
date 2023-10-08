package main

import (
	"echo-mangosteen/internal/common/data"
	"echo-mangosteen/internal/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	pingCtrl := controller.NewPingController()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ping", pingCtrl.Ping)

	_, f, err := data.NewData()

	if err != nil {
		log.Errorf("database error: %s", err)
	}

	e.Logger.Fatal(e.Start(":1323"))

	defer f()
}
