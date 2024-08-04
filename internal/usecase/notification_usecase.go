package usecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/manikandareas/genii-edu-realtime-services/internal/entity"
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

func (u *NotificationUsecase) Insert(ctx context.Context, request *model.NotificationRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	notification := &entity.Notification{
		Message: request.Message,
		UserID:  request.UserID,
		IsRead:  false,
		Timestamps: entity.Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	err := u.NotificationRepository.Create(tx, notification)
	if err != nil {
		return err
	}

	if channel, ok := u.Hub.NotificationChannel[request.UserID]; ok {
		channel <- model.NotificationResponse{
			ID:        notification.ID,
			Message:   notification.Message,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
			UserID:    notification.UserID,
			UpdatedAt: notification.UpdatedAt,
		}
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error inserting notification")
		return err
	}

	return nil
}
