package repository

import (
	"github.com/manikandareas/genii-edu-realtime-services/internal/entity"
	"github.com/sirupsen/logrus"
)

type LearningMaterialRepository struct {
	Repository[entity.LearningMaterial]
	Log *logrus.Logger
}

func NewLearningMaterialRepository(log *logrus.Logger) *LearningMaterialRepository {
	return &LearningMaterialRepository{
		Log: log,
	}
}

// func (r *LearningMaterialRepository) SendNotificationAfterCreate(tx *gorm.DB, classId string) error {
// 	recipients := new([]entity.User)

// 	// err := tx.Model(&entity.User{}).Where("class_id = ?", classId).Find(recipients).Error

// 	return nil
// }
