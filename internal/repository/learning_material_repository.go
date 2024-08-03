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
