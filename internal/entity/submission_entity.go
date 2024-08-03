package entity

import "time"

type Submission struct {
	ID           string     `gorm:"primaryKey;column:submission_id;type:text;default:uuid_generate_v7()"`
	AssignmentID string     `gorm:"column:assignment_id;type:text;not null"`
	StudentID    string     `gorm:"column:student_id;type:text;not null"`
	IsGraded     bool       `gorm:"column:is_graded;default:false;not null"`
	Grade        float64    `gorm:"column:grade;type:numeric"`
	SubmittedAt  time.Time  `gorm:"column:submitted_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	Assignment   Assignment `gorm:"foreignKey:assignment_id;references:assignment_id"`
	Student      User       `gorm:"foreignKey:student_id;references:user_id"`
}

func (s *Submission) TableName() string {
	return "submissions"
}
