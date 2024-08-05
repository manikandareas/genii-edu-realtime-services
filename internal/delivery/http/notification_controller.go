package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
	"github.com/manikandareas/genii-edu-realtime-services/internal/usecase"
	"github.com/sirupsen/logrus"
)

type NotificationController struct {
	Log     *logrus.Logger
	Usecase *usecase.NotificationUsecase
}

func NewNotificationController(log *logrus.Logger, usecase *usecase.NotificationUsecase) *NotificationController {
	return &NotificationController{
		Log:     log,
		Usecase: usecase,
	}
}

func (c *NotificationController) HandleBroadcast(ctx *fiber.Ctx) error {
	request := new(model.BroadcastRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	err := c.Usecase.HandleBroadcast(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("error broadcasting")
		return fiber.ErrBadRequest
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Broadcasted",
	})
}
