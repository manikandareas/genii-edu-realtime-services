package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/manikandareas/genii-edu-realtime-services/internal/entity"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
	"github.com/manikandareas/genii-edu-realtime-services/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SessionUsecase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	SessionRepository *repository.SessionRepository
}

func NewSessionUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, sessionRepository *repository.SessionRepository) *SessionUsecase {
	return &SessionUsecase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		SessionRepository: sessionRepository,
	}
}

func (u *SessionUsecase) Verify(ctx context.Context, userId string) (*model.SessionResponse, error) {
	session := new(entity.Session)

	err := u.SessionRepository.FindByUserIdWithUser(u.DB, session, userId)
	if err != nil {
		return nil, err
	}

	return &model.SessionResponse{
		ID:        session.ID,
		UserID:    session.UserID,
		ExpiresAt: session.ExpiresAt,
		User: model.UserResponse{
			Name:     session.User.Name,
			Username: session.User.Username,
			Email:    session.User.Email,
			Role:     string(session.User.Role),
		},
	}, nil
}
