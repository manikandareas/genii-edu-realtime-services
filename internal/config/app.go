package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http/middleware"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http/route"
	"github.com/manikandareas/genii-edu-realtime-services/internal/repository"
	"github.com/manikandareas/genii-edu-realtime-services/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

// Bootstrap sets up the application with the provided configuration.
// It sets up the repositories, use cases, controllers, middlewares, and routes.
func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	sessionRepository := repository.NewSessionRepository(config.Log)
	// setup use cases
	sessionUsecase := usecase.NewSessionUsecase(config.DB, config.Log, config.Validate, sessionRepository)
	// setup controllers

	// setup middlewares
	authMiddleware := middleware.NewAuth(sessionUsecase)
	// setup routes

	routeConfig := &route.RouteConfig{
		App:            config.App,
		AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}
