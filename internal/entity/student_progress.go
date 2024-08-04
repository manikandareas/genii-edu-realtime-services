package entity

type StudentProgress struct {
	ID           string                  `gorm:"primaryKey;column:progress_id;type:text;default:uuid_generate_v7()"`
	StudentID    string                  `gorm:"column:student_id;type:text;not null"`
	ClassID      string                  `gorm:"column:class_id;type:text;not null"`
	ProgressType StudentProgressTypeEnum `gorm:"column:progress_type;type:student_progress_type;not null"`
	Detail       string                  `gorm:"column:detail;type:text;not null"`
	Student      User                    `gorm:"foreignKey:StudentID;references:ID"`
	Class        Class                   `gorm:"foreignKey:ClassID;references:ID"`
	Timestamps
}

func (s *StudentProgress) TableName() string {
	return "student_progresses"
}
