package entity

type File struct {
	ID                 string           `gorm:"primaryKey;column:file_id;type:text;default:uuid_generate_v7()"`
	Url                string           `gorm:"column:url;type:text;not null"`
	Key                string           `gorm:"column:key;type:text;not null"`
	Name               string           `gorm:"column:name;type:text;not null"`
	Type               FilesTypeEnum    `gorm:"column:type;type:files_type;not null"`
	UserID             string           `gorm:"column:user_id;type:text;not null"`
	LearningMaterialID string           `gorm:"column:learning_material_id;type:text"`
	AssignmentID       string           `gorm:"column:assignment_id;type:text"`
	SubmissionID       string           `gorm:"column:submission_id;type:text"`
	ClassID            string           `gorm:"column:class_id;type:text"`
	IsProfilePicture   bool             `gorm:"column:is_profile_picture;default:false;not null"`
	User               User             `gorm:"foreignKey:UserID;references:ID"`
	LearningMaterial   LearningMaterial `gorm:"foreignKey:LearningMaterialID;references:ID"`
	Assignment         Assignment       `gorm:"foreignKey:AssignmentID;references:ID"`
	Submission         Submission       `gorm:"foreignKey:SubmissionID;references:ID"`
	Class              Class            `gorm:"foreignKey:ClassID;references:ID"`
	Timestamps
}

func (f *File) TableName() string {
	return "files"
}
