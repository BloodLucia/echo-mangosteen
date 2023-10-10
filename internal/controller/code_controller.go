package controller

import (
	"echo-mangosteen/internal/model"
	"echo-mangosteen/pkg/errors"
	"echo-mangosteen/pkg/response"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type codeController struct{}

// SendValidationCode implements CodeController.
func (*codeController) SendValidationCode(ctx echo.Context) error {
	reqBody := new(model.UserSendValidationCodeRequest)
	if err := ctx.Bind(reqBody); err != nil {
		return response.Build(ctx, errors.BadRequest(), nil)
	}

	isValid := validate.Struct(reqBody)

	if isValid.Validate() {
		return response.Build(ctx, nil, nil)
	}

	return response.Build(ctx, errors.InvalidRequestBody(), isValid.Errors)
}

type CodeController interface {
	SendValidationCode(ctx echo.Context) error
}

func NewCodeController() CodeController {
	return &codeController{}
}
