package entity

import "time"

type Session struct {
	ID        string    `gorm:"column:id;primaryKey"`
	UserID    string    `gorm:"column:user_id"`
	ExpiresAt time.Time `gorm:"column:expires_at;type:timestamp with time zone;not null"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
}

func (Session) TableName() string {
	return "sessions"
}
