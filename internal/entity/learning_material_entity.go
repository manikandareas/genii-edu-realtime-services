package entity

import "time"

type LearningMaterial struct {
	ID          string    `gorm:"primaryKey;column:material_id;type:text;default:uuid_generate_v7()"`
	Title       string    `gorm:"column:title;type:text;not null"`
	Content     string    `gorm:"column:content;type:text;not null"`
	AuthorID    string    `gorm:"column:author_id;type:text;not null"`
	ClassID     string    `gorm:"column:class_id;type:text;not null"`
	PublishedAt time.Time `gorm:"column:published_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null"`
	Author      User      `gorm:"foreignKey:author_id;references:user_id"`
	Class       Class     `gorm:"foreignKey:class_id;references:class_id"`
	Timestamps
}

func (l *LearningMaterial) TableName() string {
	return "learning_materials"
}
