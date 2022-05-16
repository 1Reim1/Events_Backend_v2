package controllers

import (
	"Events_Backend_v2/internal/domain/event"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(int(id))
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) FindByCoords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		latitude, err := strconv.ParseFloat(chi.URLParam(r, "latitude"), 64)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s", err)
			}
			return
		}
		longitude, err := strconv.ParseFloat(chi.URLParam(r, "longitude"), 64)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s", err)
			}
			return
		}
		radius, err := strconv.ParseFloat(chi.URLParam(r, "radius"), 64)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s", err)
			}
			return
		}
		events, err := (*c.service).FindByCoords(latitude, longitude, radius)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}
