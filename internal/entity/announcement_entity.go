package entity

import "time"

type Announcement struct {
	ID        int       `gorm:"primaryKey;column:announcement_id"`
	ClassID   string    `gorm:"column:class_id;type:text;not null"`
	Title     string    `gorm:"column:title;type:text;not null"`
	Content   string    `gorm:"column:content;type:text;not null"`
	PostedAt  time.Time `gorm:"column:posted_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null"`
	AuthorID  string    `gorm:"column:author_id;type:text;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	Class     Class     `gorm:"foreignKey:ClassID;references:ID"`
	Author    User      `gorm:"foreignKey:AuthorID;references:ID"`
}

func (a *Announcement) TableName() string {
	return "announcements"
}
