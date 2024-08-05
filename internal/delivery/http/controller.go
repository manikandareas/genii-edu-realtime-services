package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FactoryController struct {
	Log *logrus.Logger
}

func NewFactoryController(log *logrus.Logger) *FactoryController {
	return &FactoryController{
		Log: log,
	}
}

func (c *FactoryController) HandlePing(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "API is running",
	})
}

func (c *FactoryController) HandleAuthPing(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "API with Auth is running",
	})
}
