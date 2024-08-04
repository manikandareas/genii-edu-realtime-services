package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/usecase"
)

func NewTeacher(sessionUsecase *usecase.SessionUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session := GetSession(ctx)

		if session.User.Role != "teacher" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return ctx.Next()
	}
}
