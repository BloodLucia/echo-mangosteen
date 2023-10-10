package controller

import "github.com/labstack/echo/v4"

type codeController struct{}

// SendValidationCode implements CodeController.
func (*codeController) SendValidationCode(ctx echo.Context) error {
	panic("unimplemented")
}

type CodeController interface {
	SendValidationCode(ctx echo.Context) error
}

func NewCodeController() CodeController {
	return &codeController{}
}
