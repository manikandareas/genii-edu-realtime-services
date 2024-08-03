package entity

import "time"

type ClassMember struct {
	ID               int                       `gorm:"primaryKey;column:class_member_id"`
	UserID           string                    `gorm:"column:user_id;type:text;not null"`
	ClassID          string                    `gorm:"column:class_id;type:text;not null"`
	StatusCompletion ClassCompletionStatusEnum `gorm:"column:status_completion;type:class_completion_status;default:'ongoing';not null"`
	Role             RoleEnum                  `gorm:"column:status_role;type:role;default:'student';not null"`
	JoinedAt         time.Time                 `gorm:"column:joined_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt        time.Time                 `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	Class            Class                     `gorm:"foreignKey:class_id;references:class_id"`
	User             User                      `gorm:"foreignKey:user_id;references:user_id"`
}

func (c *ClassMember) TableName() string {
	return "class_members"
}
