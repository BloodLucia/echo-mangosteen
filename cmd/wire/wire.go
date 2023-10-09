//go:build wireinject
// +build wireinject

package wire

import (
	"echo-mangosteen/internal/common/data"
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/internal/router"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var controllerProvider = wire.NewSet(
	controller.NewPingController,
)

func NewApp() (*echo.Echo, func(), error) {
	panic(wire.Build(
		data.NewData,
		controllerProvider,
		router.NewRouter,
	))
}
