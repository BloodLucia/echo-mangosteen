//go:build wireinject
// +build wireinject

package wire

import (
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/internal/data"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/internal/router"
	"echo-mangosteen/internal/service"
	"echo-mangosteen/pkg/cache"
	"echo-mangosteen/pkg/config"
	"echo-mangosteen/pkg/jwt"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var controllerProvider = wire.NewSet(
	controller.NewPingController,
	controller.NewUserController,
	controller.NewCodeController,
	controller.NewTagController,
)

var repoProvider = wire.NewSet(
	repo.NewUserRepo,
	repo.NewTagRepo,
)

var serviceProvider = wire.NewSet(
	service.NewUserService,
	service.NewTagService,
)

func NewApp(*config.Config) (*echo.Echo, func(), error) {
	panic(wire.Build(
		data.NewData,
		cache.NewCahce,
		repoProvider,
		serviceProvider,
		controllerProvider,
		router.NewRouter,
		jwt.New,
	))
}
