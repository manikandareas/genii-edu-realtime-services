package main

import (
	"fmt"
	"sync"

	"github.com/manikandareas/genii-edu-realtime-services/internal/config"
	"github.com/manikandareas/genii-edu-realtime-services/internal/model"
)

func main() {
	var hub = &model.Hub{
		NotificationChannel: make(map[string]chan model.Event),
		Mutex:               sync.Mutex{},
	}

	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	validator := config.NewValidator(viperConfig)
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Validate: validator,
		Log:      log,
		Config:   viperConfig,
		Hub:      hub,
	})

	go config.GracefulShutdown(&config.GracefulShutdownConfig{
		App: app,
		Hub: hub,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
