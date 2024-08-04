package model

import "github.com/manikandareas/genii-edu-realtime-services/internal/entity"

type Hub struct {
	NotificationChannel map[string]chan entity.Notification
}
