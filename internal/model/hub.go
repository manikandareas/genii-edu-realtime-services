package model

type Hub struct {
	NotificationChannel map[string]chan NotificationResponse
}
