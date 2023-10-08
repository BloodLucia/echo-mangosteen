package controller

import "github.com/labstack/echo/v4"

type controller struct {
}

// Pingc implements PingController.
func (ctrl *controller) Ping(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"msg": "pong!",
	})
}

type PingController interface {
	Ping(c echo.Context) error
}

func NewPingController() PingController {
	return &controller{}
}
