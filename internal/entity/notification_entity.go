package entity

type Notification struct {
	ID      int    `gorm:"column:notification_id;primaryKey"`
	Message string `gorm:"column:message;not null"`
	IsRead  bool   `gorm:"column:is_read;not null;default:false"`
	UserId  string `gorm:"column:user_id"`
	User    User   `gorm:"foreignKey:user_id;references:user_id"`
	Timestamps
}

func (n *Notification) TableName() string {
	return "notifications"
}
