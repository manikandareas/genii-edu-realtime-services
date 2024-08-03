package entity

type File struct {
	ID                 string           `gorm:"primaryKey;column:file_id;type:text;default:uuid_generate_v7()"`
	Url                string           `gorm:"column:url;type:text;not null"`
	Key                string           `gorm:"column:key;type:text;not null"`
	Name               string           `gorm:"column:name;type:text;not null"`
	Type               FilesTypeEnum    `gorm:"column:type;type:files_type;not null"`
	UserId             string           `gorm:"column:user_id;type:text;not null"`
	LearningMaterialId string           `gorm:"column:learning_material_id;type:text"`
	AssignmentId       string           `gorm:"column:assignment_id;type:text"`
	SubmissionId       string           `gorm:"column:submission_id;type:text"`
	ClassID            string           `gorm:"column:class_id;type:text"`
	IsProfilePicture   bool             `gorm:"column:is_profile_picture;default:false;not null"`
	User               User             `gorm:"foreignKey:user_id;references:user_id"`
	LearningMaterial   LearningMaterial `gorm:"foreignKey:learning_material_id;references:material_id"`
	Assignment         Assignment       `gorm:"foreignKey:assignment_id;references:assignment_id"`
	Submission         Submission       `gorm:"foreignKey:submission_id;references:submission_id"`
	Class              Class            `gorm:"foreignKey:class_id;references:class_id"`
	Timestamps
}

func (f *File) TableName() string {
	return "files"
}
