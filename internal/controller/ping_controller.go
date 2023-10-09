package controller

import (
	"echo-mangosteen/internal/common/data"

	"github.com/labstack/echo/v4"
)

type pingController struct {
	*data.Data
}

// Pingc implements PingController.
func (ctrl *pingController) Ping(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"msg": "pong!",
	})
}

type PingController interface {
	Ping(c echo.Context) error
}

func NewPingController(data *data.Data) PingController {
	return &pingController{Data: data}
}
