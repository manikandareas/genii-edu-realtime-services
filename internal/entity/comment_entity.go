package entity

type Comment struct {
	ID                       int    `gorm:"primaryKey;column:comment_id"`
	Content                  string `gorm:"column:content;not null"`
	AssignmentPersonalChatID string `gorm:"column:room_id;type:text;not null"`
	SenderID                 string `gorm:"column:sender_id;type:text;not null"`
	Sender                   User   `gorm:"foreignKey:sender_id;references:user_id"`
	Timestamps
}

func (c *Comment) TableName() string {
	return "comments"
}
