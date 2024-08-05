package usecase

import (
	"github.com/go-playground/validator/v10"
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
}

func NewLearningMaterialUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, learningMaterialRepository *repository.LearningMaterialRepository, classMemberRepository *repository.ClassMemberRepository) *LearningMaterialUsecase {
	return &LearningMaterialUsecase{
		DB:                         db,
		Log:                        log,
		Validate:                   validate,
		LearningMaterialRepository: learningMaterialRepository,
		ClassMemberRepository:      classMemberRepository,
	}
}
