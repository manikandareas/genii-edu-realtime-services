package sse

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/manikandareas/genii-edu-realtime-services/internal/delivery/http/middleware"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
)

type NotificationSSE struct {
	Hub *model.Hub
}

func NewNotificationSSE(hub *model.Hub) *NotificationSSE {
	return &NotificationSSE{
		Hub: hub,
	}
}

func (s *NotificationSSE) StreamNotification(ctx *fiber.Ctx) error {
	ctx.Set(("Content-Type"), "text/event-stream")
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")

	session := middleware.GetSession(ctx)
	s.Hub.Mutex.Lock()
	s.Hub.NotificationChannel[session.UserID] = make(chan model.Event)
	s.Hub.Mutex.Unlock()
	ctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		event := fmt.Sprintf("event: %s\n"+
			"data: \n\n", "initial")
		fmt.Fprint(w, event)
		w.Flush()

		for e := range s.Hub.NotificationChannel[session.UserID] {
			data, _ := json.Marshal(e)

			event = fmt.Sprintf("event: %s\n"+
				"data: %s\n\n", data, data)

			_, _ = fmt.Fprint(w, event)
			_ = w.Flush()
		}
	})
	return nil
}
