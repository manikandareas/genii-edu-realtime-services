package http

import (
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
