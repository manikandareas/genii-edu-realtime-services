package model

type SendNotificationRequest struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"author_id"`
	ClassID  string `json:"class_id"`
}

// ID          string    `gorm:"primaryKey;column:material_id;type:text;default:uuid_generate_v7()"`
// Title       string    `gorm:"column:title;type:text;not null"`
// Content     string    `gorm:"column:content;type:text;not null"`
// AuthorID    string    `gorm:"column:author_id;type:text;not null"`
// ClassID     string    `gorm:"column:class_id;type:text;not null"`
// PublishedAt time.Time `gorm:"column:published_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null"`
// Author      User      `gorm:"foreignKey:AuthorID;references:ID"`
// Class       Class     `gorm:"foreignKey:ClassID;references:ID"`
// Timestamps
