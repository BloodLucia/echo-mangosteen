package controller

import (
	"echo-mangosteen/pkg/response"

	"github.com/labstack/echo/v4"
)

type pingController struct {
}

// Pingc implements PingController.
func (ctrl *pingController) Ping(c echo.Context) error {
	return response.Build(c, nil, "pong!")
}

type PingController interface {
	Ping(c echo.Context) error
}

func NewPingController() PingController {
	return &pingController{}
}
