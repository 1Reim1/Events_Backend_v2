package http

import (
	"Events_Backend_v2/internal/infra/http/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router(eventController *controllers.EventController) http.Handler {
	router := chi.NewRouter()

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/", func(apiRouter chi.Router) {
			AddEventRoutes(&apiRouter, eventController)

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})
	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			eventController.FindAll(),
		)
		apiRouter.Get(
			"/{id}",
			eventController.FindOne(),
		)
		apiRouter.Get(
			"/{latitude}/{longitude}/{radius}",
			eventController.FindByCoords(),
		)
	})
}
