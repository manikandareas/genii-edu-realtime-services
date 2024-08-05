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

type BroadcastRequest struct {
	Event      string   `json:"event" validate:"required,min=3"`
	Recipients []string `json:"recipients" validate:"required"`
}
