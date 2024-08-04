package model

import "time"

type NotificationResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NotificationRequest struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}
