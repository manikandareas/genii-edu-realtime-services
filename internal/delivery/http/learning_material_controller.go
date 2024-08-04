package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http/middleware"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
	"github.com/manikandareas/genii-edu-realtime-services/internal/usecase"
	"github.com/sirupsen/logrus"
)

type LearningMaterialController struct {
	Log     *logrus.Logger
	Usecase *usecase.LearningMaterialUsecase
}

func NewLearningMaterialController(log *logrus.Logger, usecase *usecase.LearningMaterialUsecase) *LearningMaterialController {
	return &LearningMaterialController{
		Log:     log,
		Usecase: usecase,
	}
}

func (c *LearningMaterialController) SendNotificationAfterCreate(ctx *fiber.Ctx) error {
	request := new(model.SendNotificationRequest)
	session := middleware.GetSession(ctx)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	err := c.Usecase.SendNotification(ctx.Context(), request, session)

	if err != nil {
		c.Log.WithError(err).Error("Failed to send notification after create")
		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Notification sent",
	})
}
