package model

type Hub struct {
	NotificationChannel map[string]chan Event
}

type Event struct {
	Event string `json:"event"`
}
