package entity

import "time"

type Session struct {
	ID        string    `gorm:"column:id;primaryKey"`
	UserId    string    `gorm:"column:user_id"`
	User      User      `gorm:"foreignKey:user_id;references:user_id"`
	ExpiresAt time.Time `gorm:"column:expires_at;type:timestamp with time zone;not null"`
}

func (s *Session) TableName() string {
	return "sessions"
}
