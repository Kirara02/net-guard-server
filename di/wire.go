//go:build wireinject
// +build wireinject

package di

import (
	"NetGuardServer/controllers"
	"NetGuardServer/repository"
	"NetGuardServer/services"

	"github.com/google/wire"
)

// Provider set for repositories
var repositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewServerRepository,
	repository.NewHistoryRepository,
)

// Provider set for services
var serviceSet = wire.NewSet(
	services.NewAuthService,
	services.NewServerService,
	services.NewHistoryService,
	services.NewNotificationService,
)

// Provider set for controllers
var controllerSet = wire.NewSet(
	controllers.NewAuthController,
	controllers.NewServerController,
	controllers.NewHistoryController,
)

// App holds all application dependencies
type App struct {
	AuthController    *controllers.AuthController
	ServerController  *controllers.ServerController
	HistoryController *controllers.HistoryController
}

// InitializeApp initializes the entire application with dependency injection
func InitializeApp() (*App, error) {
	wire.Build(
		repositorySet,
		serviceSet,
		controllerSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}