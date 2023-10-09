package controller

import (
	"echo-mangosteen/internal/service"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService service.UserService
}

// Login implements UserController.
func (ctrl *userController) Login(c echo.Context) error {
	return c.JSON(200, "login!")
}

type UserController interface {
	Login(c echo.Context) error
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}
