package controller

import (
	"echo-mangosteen/internal/model"
	"echo-mangosteen/internal/service"
	"echo-mangosteen/pkg/errors"
	"echo-mangosteen/pkg/response"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type userController struct {
	userService service.UserService
}

// Login implements UserController.
func (ctrl *userController) Login(ctx echo.Context) error {
	reqBody := new(model.UserLoginRequest)
	if err := ctx.Bind(reqBody); err != nil {
		return response.Build(ctx, err, nil)
	}
	v := validate.Struct(reqBody)
	if v.Validate() {
		return ctx.JSON(200, "login!")
	}
	return response.Build(ctx, errors.BadRequest(), v.Errors)
}

type UserController interface {
	Login(ctx echo.Context) error
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}
