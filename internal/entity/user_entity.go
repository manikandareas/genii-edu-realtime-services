package entity

type User struct {
	ID                 string      `gorm:"column:user_id;primaryKey"`
	Name               string      `gorm:"column:name"`
	Username           string      `gorm:"column:username;unique"`
	Email              string      `gorm:"column:email"`
	PasswordHash       string      `gorm:"column:password_hash"`
	Role               RoleEnum    `gorm:"column:role;type:role"`
	ProfilePicture     string      `gorm:"column:profile_picture"`
	Bio                string      `gorm:"column:bio"`
	IsEmailVerified    bool        `gorm:"column:is_email_verified;default:false;not null"`
	OnBoardingComplete bool        `gorm:"column:on_boarding_complete;default:false;not null"`
	UserContact        UserContact `gorm:"foreignKey:ID;references:UserID"`
	Classes            []ClassMember
	Timestamps
}

func (u *User) TableName() string {
	return "users"
}

// type role string

// const (
// 	teacher role = "teacher"
// 	student role = "student"
// )

// func (self *role) Scan(value interface{}) error {
// 	*self = role(value.([]byte))
// 	return nil
// }

// func (self role) Value() (driver.Value, error) {
// 	return string(self), nil
// }
