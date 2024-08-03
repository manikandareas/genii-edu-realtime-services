package entity

import "time"

type Assignment struct {
	ID          string    `gorm:"primaryKey;column:assignment_id;type:text;default:uuid_generate_v7()"`
	Title       string    `gorm:"column:title;type:text;not null"`
	Description string    `gorm:"column:description;type:text;not null"`
	AuthorID    string    `gorm:"column:author_id;type:text;not null"`
	DueDate     time.Time `gorm:"column:due_date;type:timestamp with time zone;not null"`
	ClassID     string    `gorm:"column:class_id;type:text;not null"`
	IsOpen      bool      `gorm:"column:is_open;default:true;not null"`
	PublishedAt time.Time `gorm:"column:published_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null"`
	Author      User      `gorm:"foreignKey:author_id;references:user_id"`
	Class       Class     `gorm:"foreignKey:class_id;references:class_id"`
	Timestamps
}

func (a *Assignment) TableName() string {
	return "assignments"
}
