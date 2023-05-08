package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Post("/v1/channels", app.createChannelHandler)
	router.Get("/v1/channels/{id}", app.showChannelHandler)

	return router
}
