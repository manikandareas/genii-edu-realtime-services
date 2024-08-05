package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http"
)

type RouteConfig struct {
	App                        *fiber.App
	AuthMiddleware             fiber.Handler
	TeacherMiddleware          fiber.Handler
	LearningMaterialController *http.LearningMaterialController
	NotificationController     *http.NotificationController
	FactoryController          *http.FactoryController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/ping", c.FactoryController.HandlePing)
	c.App.Get("/api/metrics", monitor.New(monitor.Config{Title: "Genii Edu Realtime Services Metrics"}))
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/auth/ping", c.FactoryController.HandleAuthPing)
	c.App.Post("/api/notifications/broadcast", c.NotificationController.HandleBroadcast)
}
