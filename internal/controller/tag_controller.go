package controller

import (
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/service"
	"echo-mangosteen/pkg/errors"
	"echo-mangosteen/pkg/response"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type tagController struct {
	service service.TagService
}

// AddTag implements TagController.
func (tc *tagController) AddTag(ctx echo.Context) error {
	reqBody := &model.TagAddRequest{}
	userId := CurrentUser(ctx)
	if err := ctx.Bind(reqBody); err != nil {
		return response.Build(ctx, errors.InvalidRequestBody(), nil)
	}
	v := validate.Struct(reqBody)
	if v.Validate() {
		if err := tc.service.AddTag(ctx.Request().Context(), userId, reqBody); err != nil {
			return response.Build(ctx, err, nil)
		}

		return response.Build(ctx, nil, nil)
	}

	return response.Build(ctx, errors.BadRequest(), v.Errors)
}

type TagController interface {
	AddTag(ctx echo.Context) error
}

func NewTagController(service service.TagService) TagController {
	return &tagController{service}
}
