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
}

// Bootstrap sets up the application with the provided configuration.
// It sets up the repositories, use cases, controllers, middlewares, and routes.
func Bootstrap(config *BootstrapConfig) {
	// hub for sse
	hub := &model.Hub{
		NotificationChannel: map[string]chan model.NotificationResponse{},
	}

	// setup repositories
	sessionRepository := repository.NewSessionRepository(config.Log)
	notificationRepository := repository.NewNotificationRepository(config.Log)
	learningMaterialRepository := repository.NewLearningMaterialRepository(config.Log)
	classMemberRepository := repository.NewClassMemberRepository(config.Log)

	// setup use cases
	sessionUsecase := usecase.NewSessionUsecase(config.DB, config.Log, config.Validate, sessionRepository)
	notificationUsecase := usecase.NewNotificationUsecase(config.DB, config.Log, config.Validate, notificationRepository, hub)
	learningMaterialUsecase := usecase.NewLearningMaterialUsecase(config.DB, config.Log, config.Validate, learningMaterialRepository, classMemberRepository, notificationUsecase)

	// setup controllers
	learningMaterialController := http.NewLearningMaterialController(config.Log, learningMaterialUsecase)

	// setup middlewares
	authMiddleware := middleware.NewAuth(sessionUsecase)
	teacherMiddleware := middleware.NewTeacher(sessionUsecase)

	// setup sse
	notificationSSE := sse.NewNotificationSSE(hub)

	// setup routes
	routeConfig := &route.RouteConfig{
		App:                        config.App,
		AuthMiddleware:             authMiddleware,
		LearningMaterialController: learningMaterialController,
		TeacherMiddleware:          teacherMiddleware,
	}

	sseConfig := &sseconfig.SSERoute{
		App:             config.App,
		Hub:             hub,
		AuthMiddleware:  authMiddleware,
		NotificationSSE: notificationSSE,
	}

	routeConfig.Setup()
	sseConfig.Setup()
}
