package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:           config.GetString("app.name"),
		ErrorHandler:      NewErrorHandler(),
		Prefork:           config.GetBool("web.prefork"),
		EnablePrintRoutes: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
}
