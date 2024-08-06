package model

import "sync"

type Hub struct {
	NotificationChannel map[string]chan Event
	Mutex               sync.Mutex
}

type Event struct {
	Event string `json:"event"`
}
