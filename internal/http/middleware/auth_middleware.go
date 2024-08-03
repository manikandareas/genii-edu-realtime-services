package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
	"github.com/manikandareas/genii-edu-realtime-services/internal/usecase"
)

func NewAuth(sessionUsecase *usecase.SessionUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := new(model.Auth)

		if err := ctx.BodyParser(auth); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		session, err := sessionUsecase.Verify(ctx.Context(), auth.ID)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		sessionUsecase.Log.Debugf("User : %+v", auth.ID)

		ctx.Locals("session", session)

		return ctx.Next()
	}
}

func GetSession(ctx *fiber.Ctx) *model.SessionResponse {
	return ctx.Locals("session").(*model.SessionResponse)
}
