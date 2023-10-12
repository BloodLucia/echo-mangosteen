package controller

import (
	"echo-mangosteen/internal/service"

	"github.com/labstack/echo/v4"
)

type itemController struct {
	service service.ItemService
}

// CreateItem implements ItemController.
func (*itemController) CreateItem(c echo.Context) error {
	panic("unimplemented")
}

type ItemController interface {
	CreateItem(c echo.Context) error
}

func NewItemController(service service.ItemService) ItemController {
	return &itemController{service}
}
