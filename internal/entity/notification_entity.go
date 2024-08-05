package entity

type Notification struct {
	ID      int    `gorm:"column:notification_id;primaryKey"`
	Title   string `gorm:"column:title;not null"`
	Message string `gorm:"column:message;not null"`
	Url     string `gorm:"column:url"`
	IsRead  bool   `gorm:"column:is_read;not null;default:false"`
	UserID  string `gorm:"column:user_id"`

	User User `gorm:"foreignKey:user_id;references:ID"`
	Timestamps
}

func (n *Notification) TableName() string {
	return "notifications"
}
