//go:build wireinject
// +build wireinject

package wire

import (
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/internal/data"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/internal/router"
	"echo-mangosteen/internal/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var controllerProvider = wire.NewSet(
	controller.NewPingController,
	controller.NewUserController,
)

var repoProvider = wire.NewSet(
	repo.NewUserRepo,
)

var serviceProvider = wire.NewSet(
	service.NewUserService,
)

func NewApp() (*echo.Echo, func(), error) {
	panic(wire.Build(
		data.NewData,
		repoProvider,
		serviceProvider,
		controllerProvider,
		router.NewRouter,
	))
}
