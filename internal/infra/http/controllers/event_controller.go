package controllers

import (
	"Events_Backend_v2/internal/domain/event"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
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
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s\n", err)
			}
			return
		}
		event, err := (*c.service).FindOne(int(id))
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

func (c *EventController) FindByCoords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		latitude, err := strconv.ParseFloat(chi.URLParam(r, "latitude"), 64)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s\n", err)
			}
			return
		}
		longitude, err := strconv.ParseFloat(chi.URLParam(r, "longitude"), 64)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s\n", err)
			}
			return
		}
		radius, err := strconv.ParseFloat(chi.URLParam(r, "radius"), 64)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s\n", err)
			}
			return
		}
		events, err := (*c.service).FindByCoords(latitude, longitude, radius)
		if err != nil {
			fmt.Printf("EventController.FindByCoords(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindByCoords(): %s\n", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s\n", err)
		}
	}
}

func (c *EventController) PostOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("EventController.PostOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.PostOne(): %s\n", err)
			}
			return
		}
		err = (*c.service).PostOne(body)
		if err != nil {
			fmt.Printf("EventController.PostOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.PostOne(): %s\n", err)
			}
			return
		}

		err = success(w, "posted")
		if err != nil {
			fmt.Printf("EventController.PostOne(): %s\n", err)
		}
	}
}

func (c *EventController) UpdateOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.UpdateOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.UpdateOne(): %s\n", err)
			}
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("EventController.UpdateOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.UpdateOne(): %s\n", err)
			}
			return
		}

		err = (*c.service).UpdateOne(int(id), body)
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
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.DeleteOne(): %s\n", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DeleteOne(): %s\n", err)
			}
			return
		}

		err = (*c.service).DeleteOne(int(id))
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
