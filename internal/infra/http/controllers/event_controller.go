package controllers

import (
	"Events_Backend_v2/internal/domain/event"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
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
		var latitude, longitude, radius float64
		var err error
		if r.URL.Query().Has("latitude") && r.URL.Query().Has("longitude") && r.URL.Query().Has("radius") {
			latitude, err = strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s\n", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.FindAll(): %s\n", err)
				}
				return
			}
			longitude, err = strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s\n", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.FindAll(): %s\n", err)
				}
				return
			}
			radius, err = strconv.ParseFloat(r.URL.Query().Get("radius"), 64)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s\n", err)
				err = internalServerError(w, err)
				if err != nil {
					fmt.Printf("EventController.FindAll(): %s\n", err)
				}
				return
			}
		}

		events, err := (*c.service).FindAll(latitude, longitude, radius)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s\n", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s\n", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s\n", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s\n", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s\n", err)
		}
	}
}

func (c *EventController) CreateOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event event.Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			fmt.Printf("EventController.CreateOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateOne(): %s\n", err)
			}
			return
		}
		err = (*c.service).CreateOne(&event)
		if err != nil {
			fmt.Printf("EventController.CreateOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateOne(): %s\n", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.CreateOne(): %s\n", err)
		}
	}
}

func (c *EventController) UpdateOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event event.Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			fmt.Printf("EventController.UpdateOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.UpdateOne(): %s\n", err)
			}
			return
		}

		err = (*c.service).UpdateOne(&event)
		if err != nil {
			fmt.Printf("EventController.UpdateOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.UpdateOne(): %s\n", err)
			}
			return
		}

		err = success(w, "updated")
		if err != nil {
			fmt.Printf("EventController.UpdateOne(): %s\n", err)
		}
	}
}

func (c *EventController) DeleteOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.DeleteOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DeleteOne(): %s\n", err)
			}
			return
		}

		err = (*c.service).DeleteOne(id)
		if err != nil {
			fmt.Printf("EventController.DeleteOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DeleteOne(): %s\n", err)
			}
			return
		}

		err = success(w, "deleted")
		if err != nil {
			fmt.Printf("EventController.DeleteOne(): %s\n", err)
		}
	}
}
