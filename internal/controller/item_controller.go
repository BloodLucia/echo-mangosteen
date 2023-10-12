package controller

import "echo-mangosteen/internal/service"

type itemController struct {
	service service.ItemService
}

type ItemController interface {
}

func NewItemController(service service.ItemService) ItemController {
	return &itemController{service}
}
