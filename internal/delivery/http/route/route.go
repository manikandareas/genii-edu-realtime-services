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
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "API is running",
		})
	})
	c.App.Get("/api/metrics", monitor.New(monitor.Config{Title: "Genii Edu Realtime Services Metrics"}))
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)

	c.App.Get("/api/auth/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "API with Auth is running",
		})
	})

	c.App.Post("/api/learning-material/new", c.TeacherMiddleware, c.LearningMaterialController.SendNotificationAfterCreate)

}
