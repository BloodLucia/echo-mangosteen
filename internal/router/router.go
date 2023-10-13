package router

import (
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/pkg/jwt"
	"echo-mangosteen/pkg/middleware"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	jwt *jwt.JWT,
	userCtrl controller.UserController,
	codeCtrl controller.CodeController,
	tagCtrl controller.TagController,
	itemCtrl controller.ItemController,
) *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.Logger())

	v1Auth := e.Group("/api/v1")
	v1Auth.Use(middleware.JWTAuth(jwt, "/login"))
	{

		// tags
		v1Auth.POST("/tags/add", tagCtrl.AddTag)
		v1Auth.DELETE("/tags/delete/:tagId", tagCtrl.DeleteTag)

		// items
		v1Auth.POST("/items/add", itemCtrl.CreateItem)

		// no auth
		v1Auth.POST("/login", userCtrl.Login)
	}

	return e
}
