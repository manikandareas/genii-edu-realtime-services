package repository

import (
	"github.com/manikandareas/genii-edu-realtime-services/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ClassMemberRepository struct {
	Repository[entity.ClassMember]
	Log *logrus.Logger
}

func NewClassMemberRepository(log *logrus.Logger) *ClassMemberRepository {
	return &ClassMemberRepository{
		Log: log,
	}
}

func (r *ClassMemberRepository) FindByClassIdAndRole(tx *gorm.DB, classId string, role string) ([]entity.ClassMember, error) {
	var classMembers []entity.ClassMember

	err := tx.Where("class_id = ?", classId).Where("status_role = ?", role).Find(&classMembers).Error
	if err != nil {
		return nil, err
	}

	return classMembers, nil
}
