package entity

type PersonalComment struct {
	ID           string     `gorm:"primaryKey;column:assignment_personal_comment_id;type:text;default:uuid_generate_v7()"`
	AssignmentID string     `gorm:"column:assignment_id;type:text;not null"`
	StudentID    string     `gorm:"column:student_id;type:text;not null"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentID;references:ID"`
	Student      User       `gorm:"foreignKey:StudentID;references:ID"`
	Comments     []Comment  `gorm:"foreignKey:ID;references:AssignmentPersonalChatID"`
	Timestamps
}

func (p *PersonalComment) TableName() string {
	return "assignment_personal_comments"
}
