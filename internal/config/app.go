package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http/middleware"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http/route"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/sse"
	sseconfig "github.com/manikandareas/genii-edu-realtime-services/internal/delivery/sse/sse_route"

	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
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
	Hub      *model.Hub
}

// Bootstrap sets up the application with the provided configuration.
// It sets up the repositories, use cases, controllers, middlewares, and routes.
func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	sessionRepository := repository.NewSessionRepository(config.Log)
	notificationRepository := repository.NewNotificationRepository(config.Log)
	learningMaterialRepository := repository.NewLearningMaterialRepository(config.Log)
	classMemberRepository := repository.NewClassMemberRepository(config.Log)

	// setup use cases
	sessionUsecase := usecase.NewSessionUsecase(config.DB, config.Log, config.Validate, sessionRepository)
	notificationUsecase := usecase.NewNotificationUsecase(config.DB, config.Log, config.Validate, notificationRepository, config.Hub)
	learningMaterialUsecase := usecase.NewLearningMaterialUsecase(config.DB, config.Log, config.Validate, learningMaterialRepository, classMemberRepository)

	// setup controllers
	factoryController := http.NewFactoryController(config.Log)
	learningMaterialController := http.NewLearningMaterialController(config.Log, learningMaterialUsecase)
	notificationController := http.NewNotificationController(config.Log, notificationUsecase)
	// setup middlewares
	authMiddleware := middleware.NewAuth(sessionUsecase)
	teacherMiddleware := middleware.NewTeacher(sessionUsecase)

	// setup sse
	notificationSSE := sse.NewNotificationSSE(config.Hub)

	// setup routes
	routeConfig := &route.RouteConfig{
		App:                        config.App,
		AuthMiddleware:             authMiddleware,
		LearningMaterialController: learningMaterialController,
		TeacherMiddleware:          teacherMiddleware,
		FactoryController:          factoryController,
		NotificationController:     notificationController,
	}

	sseConfig := &sseconfig.SSERoute{
		App:             config.App,
		Hub:             config.Hub,
		AuthMiddleware:  authMiddleware,
		NotificationSSE: notificationSSE,
	}

	routeConfig.Setup()
	sseConfig.Setup()
}
