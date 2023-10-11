package controller

import (
	"echo-mangosteen/pkg/response"

	"github.com/labstack/echo/v4"
)

type tagController struct {
}

// AddTag implements TagController.
func (*tagController) AddTag(ctx echo.Context) error {
	return response.Build(ctx, nil, "add tag")
}

type TagController interface {
	AddTag(ctx echo.Context) error
}

func NewTagController() TagController {
	return &tagController{}
}
