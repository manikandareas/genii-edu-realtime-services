package usecase

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
	"github.com/manikandareas/genii-edu-realtime-services/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LearningMaterialUsecase struct {
	DB                         *gorm.DB
	Log                        *logrus.Logger
	Validate                   *validator.Validate
	LearningMaterialRepository *repository.LearningMaterialRepository
	ClassMemberRepository      *repository.ClassMemberRepository
	NotificationUsecase        *NotificationUsecase
}

func NewLearningMaterialUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, learningMaterialRepository *repository.LearningMaterialRepository, classMemberRepository *repository.ClassMemberRepository, notificationUsecase *NotificationUsecase) *LearningMaterialUsecase {
	return &LearningMaterialUsecase{
		DB:                         db,
		Log:                        log,
		Validate:                   validate,
		LearningMaterialRepository: learningMaterialRepository,
		ClassMemberRepository:      classMemberRepository,
		NotificationUsecase:        notificationUsecase,
	}
}

func (u *LearningMaterialUsecase) SendNotification(ctx context.Context, request *model.SendNotificationRequest, session *model.SessionResponse) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	recipients, err := u.ClassMemberRepository.FindByClassIdAndRole(tx, request.ClassID, "student")
	if err != nil {
		return err
	}

	for _, recipient := range recipients {
		notification := &model.NotificationRequest{
			UserID:  recipient.UserID,
			Message: fmt.Sprintf("%s Just uploaded by %s", request.Title, session.User.Name),
		}

		go func(data *model.NotificationRequest) {
			u.NotificationUsecase.Insert(context.Background(), data)
		}(notification)
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}

	return nil
}
