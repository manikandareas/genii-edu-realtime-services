package entity

import "time"

type EmailVerification struct {
	ID     int       `gorm:"primaryKey;column:email_verification_id"`
	UserID string    `gorm:"column:user_id;type:text;not null"`
	Code   string    `gorm:"column:code;type:text;not null"`
	SentAt time.Time `gorm:"column:sent_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null"`
	User   User      `gorm:"foreignKey:user_id;references:user_id"`
	Timestamps
}

func (e *EmailVerification) TableName() string {
	return "email_verifications"
}
