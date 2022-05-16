package http

import (
	"Events_Backend_v2/cmd/config"
	"net/http"
)

func Server(router http.Handler, conf *config.Config) error {
	srv := &http.Server{
		Addr:    conf.BindAddr,
		Handler: router,
	}

	return srv.ListenAndServe()
}
