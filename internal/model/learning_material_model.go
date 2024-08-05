package model

type SendNotificationRequest struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Url     string `json:"url"`
	ClassID string `json:"class_id"`
}
