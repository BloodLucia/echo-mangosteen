package controller

import "github.com/labstack/echo/v4"

type userController struct {
}

// Login implements UserController.
func (ctrl *userController) Login(c echo.Context) error {
	return c.JSON(200, "login!")
}

type UserController interface {
	Login(c echo.Context) error
}

func NewUserController() UserController {
	return &userController{}
}
