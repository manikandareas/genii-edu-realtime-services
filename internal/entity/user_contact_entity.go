package entity

type UserContact struct {
	ID        int    `gorm:"primaryKey;column:user_contact_id;autoIncrement"`
	UserID    string `gorm:"column:user_id;type:text;not null"`
	Instagram string `gorm:"column:instagram;type:text"`
	Twitter   string `gorm:"column:twitter;type:text"`
	Github    string `gorm:"column:github;type:text"`
	Linkedin  string `gorm:"column:linkedin;type:text"`
	Website   string `gorm:"column:website;type:text"`
	Whatsapp  string `gorm:"column:whatsapp;type:text"`
	Facebook  string `gorm:"column:facebook;type:text"`
	// User      User   `gorm:"foreignKey:user_id;references:user_id"`
	Timestamps
}

func (u *UserContact) TableName() string {
	return "user_contacts"
}
