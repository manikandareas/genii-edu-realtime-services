package entity

type Class struct {
	ID          string              `gorm:"primaryKey;column:class_id;type:text;default:uuid_generate_v7()"`
	ClassName   string              `gorm:"column:class_name;type:text;not null"`
	Slug        string              `gorm:"column:slug;type:text;unique;not null"`
	Description string              `gorm:"column:description;type:text"`
	ClassCode   string              `gorm:"column:class_code;type:text;not null"`
	TeacherID   string              `gorm:"column:teacher_id;type:text;not null"`
	AccessType  ClassAccessTypeEnum `gorm:"column:access_type;type:class_access_type;default:'private';not null"`
	Teacher     User                `gorm:"foreignKey:teacher_id;references:user_id"`
	Timestamps
}

func (c *Class) TableName() string {
	return "classes"
}
