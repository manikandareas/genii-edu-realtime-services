package repository

import (
	"github.com/manikandareas/genii-edu-realtime-services/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SessionRepository struct {
	Repository[entity.Session]
	Log *logrus.Logger
}

func NewSessionRepository(log *logrus.Logger) *SessionRepository {
	return &SessionRepository{
		Log: log,
	}
}

func (r *SessionRepository) FindByUserId(db *gorm.DB, session *entity.Session, userId string) error {
	return db.Where("user_id = ?", userId).First(session).Error
}
