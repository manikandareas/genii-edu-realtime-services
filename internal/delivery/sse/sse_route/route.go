package sseconfig

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/sse"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
)

type SSERoute struct {
	App             *fiber.App
	AuthMiddleware  fiber.Handler
	Hub             *model.Hub
	NotificationSSE *sse.NotificationSSE
}

func (c *SSERoute) Setup() {
	c.SetupNotificationRoute()
}

func (c *SSERoute) SetupNotificationRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/users/:id/notification-stream", c.NotificationSSE.StreamNotification)
}
