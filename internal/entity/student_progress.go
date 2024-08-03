package entity

type StudentProgress struct {
	ID           string                  `gorm:"primaryKey;column:progress_id;type:text;default:uuid_generate_v7()"`
	StudentID    string                  `gorm:"column:student_id;type:text;not null"`
	ClassID      string                  `gorm:"column:class_id;type:text;not null"`
	ProgressType StudentProgressTypeEnum `gorm:"column:progress_type;type:student_progress_type;not null"`
	Detail       string                  `gorm:"column:detail;type:text;not null"`
	Student      User                    `gorm:"foreignKey:student_id;references:user_id"`
	Class        Class                   `gorm:"foreignKey:class_id;references:class_id"`
	Timestamps
}

func (s *StudentProgress) TableName() string {
	return "student_progresses"
}
