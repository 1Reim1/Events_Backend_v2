package main

import (
	"Events_Backend_v2/cmd/config"
	"Events_Backend_v2/internal/domain/event"
	"Events_Backend_v2/internal/infra/http"
	"Events_Backend_v2/internal/infra/http/controllers"
	"fmt"
)

func main() {
	// Config
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Printf("config.NewConfig() error: %s", err)
	}
	// Event
	eventRepository, err := event.NewMySQLRepository(conf)
	if err != nil {
		fmt.Printf("event.NewMySQLRepository() error: %s", err)
	}
	eventService := event.NewSimpleService(eventRepository)
	eventController := controllers.NewEventController(&eventService)
	// Server
	err = http.Server(
		http.Router(eventController),
		conf,
	)
	if err != nil {
		fmt.Printf("http.Server error: %s", err)
		return
	}
}
