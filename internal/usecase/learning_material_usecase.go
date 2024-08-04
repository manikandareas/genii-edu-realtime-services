package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/manikandareas/genii-edu-realtime-services/internal/entity"
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
	NotificationRepository     *repository.NotificationRepository
}

func NewLearningMaterialUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, learningMaterialRepository *repository.LearningMaterialRepository, classMemberRepository *repository.ClassMemberRepository, notificationRepository *repository.NotificationRepository) *LearningMaterialUsecase {
	return &LearningMaterialUsecase{
		DB:                         db,
		Log:                        log,
		Validate:                   validate,
		LearningMaterialRepository: learningMaterialRepository,
		ClassMemberRepository:      classMemberRepository,
		NotificationRepository:     notificationRepository,
	}
}

func (u *LearningMaterialUsecase) SendNotificationAfterCreate(ctx context.Context, classId string) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	recipients, err := u.ClassMemberRepository.FindByClassIdAndRole(tx, classId, "student")
	if err != nil {
		return err
	}

	for _, recipient := range recipients {
		notification := &entity.Notification{
			UserID:  recipient.UserID,
			Message: "New learning material has been added",
			IsRead:  false,
		}

		err := u.NotificationRepository.Create(tx, notification)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}
	return nil
}
