package entity

type PersonalComment struct {
	ID           string     `gorm:"primaryKey;column:assignment_personal_comment_id;type:text;default:uuid_generate_v7()"`
	AssignmentID string     `gorm:"column:assignment_id;type:text;not null"`
	StudentID    string     `gorm:"column:student_id;type:text;not null"`
	Assignment   Assignment `gorm:"foreignKey:assignment_id;references:assignment_id"`
	Student      User       `gorm:"foreignKey:student_id;references:user_id"`
	Comments     []Comment  `gorm:"foreignKey:assignment_personal_comment_id;references:room_id"`
	Timestamps
}

func (p *PersonalComment) TableName() string {
	return "assignment_personal_comments"
}
