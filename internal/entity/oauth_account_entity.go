package entity

import "time"

type OAuthAccount struct {
	ID             string    `gorm:"primaryKey;column:id;type:text;default:uuid_generate_v7()"`
	UserID         string    `gorm:"column:user_id;type:text;not null"`
	Provider       string    `gorm:"column:provider;type:text;not null"`
	ProviderUserID string    `gorm:"column:provider_user_id;type:text;not null"`
	AccessToken    string    `gorm:"column:access_token;type:text"`
	RefreshToken   string    `gorm:"column:refresh_token;type:text"`
	ExpiresAt      time.Time `gorm:"column:expires_at;type:timestamp with time zone"`
	User           User      `gorm:"foreignKey:user_id;references:user_id"`
	Timestamps
}

func (o *OAuthAccount) TableName() string {
	return "oauth_accounts"
}
