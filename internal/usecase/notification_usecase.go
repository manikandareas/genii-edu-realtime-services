package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
	"github.com/manikandareas/genii-edu-realtime-services/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type NotificationUsecase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	NotificationRepository *repository.NotificationRepository
	Hub                    *model.Hub
}

func NewNotificationUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, notificationRepository *repository.NotificationRepository, hub *model.Hub) *NotificationUsecase {
	return &NotificationUsecase{
		DB:                     db,
		Log:                    log,
		Validate:               validate,
		NotificationRepository: notificationRepository,
		Hub:                    hub,
	}
}

func (u *NotificationUsecase) HandleBroadcast(ctx context.Context, request *model.BroadcastRequest) error {
	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}
	for _, recipient := range request.Recipients {
		go func() {
			if channel, ok := u.Hub.NotificationChannel[recipient]; ok {
				channel <- model.Event{
					Event: request.Event,
				}
			}
		}()
	}
	return nil
}
