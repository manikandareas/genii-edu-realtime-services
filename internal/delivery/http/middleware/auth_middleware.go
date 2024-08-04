package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
	"github.com/manikandareas/genii-edu-realtime-services/internal/usecase"
)

func NewAuth(sessionUsecase *usecase.SessionUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Get("x-user-id")

		if userId == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		session, err := sessionUsecase.Verify(ctx.Context(), userId)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		sessionUsecase.Log.Debugf("User : %+v", userId)

		ctx.Locals("session", session)

		return ctx.Next()
	}
}

func GetSession(ctx *fiber.Ctx) *model.SessionResponse {
	return ctx.Locals("session").(*model.SessionResponse)
}
