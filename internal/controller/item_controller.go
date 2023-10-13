package controller

import (
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/service"
	"echo-mangosteen/pkg/errors"
	"echo-mangosteen/pkg/response"
	"net/http"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type itemController struct {
	service service.ItemService
}

// CreateItem implements ItemController.
func (ic *itemController) CreateItem(c echo.Context) error {
	var reqBody model.CreateItemRequest
	userId := CurrentUser(c)
	if err := c.Bind(&reqBody); err != nil {
		// TODO 自定义错误
		return response.Build(c, errors.New(http.StatusUnprocessableEntity, "无法解析请求体"), nil)
	}
	v := validate.Struct(&reqBody)
	if v.Validate() {
		if err := ic.service.CreateItem(c.Request().Context(), userId, &reqBody); err != nil {
			return response.Build(c, err, nil)
		}

		return response.Build(c, nil, nil)
	}

	return response.Build(c, errors.InvalidRequestBody(), v.Errors)
}

type ItemController interface {
	CreateItem(c echo.Context) error
}

func NewItemController(service service.ItemService) ItemController {
	return &itemController{service}
}
