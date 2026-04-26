//go:build wireinject

package wire

import (
	"simplebank/internal/controller"
	"simplebank/internal/repository"
	"simplebank/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}