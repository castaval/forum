package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Get("/v1/channels", app.listChannelsHandler)
	router.Post("/v1/channels", app.createChannelHandler)
	router.Get("/v1/channels/{id}", app.showChannelHandler)
	router.Patch("/v1/channels/{id}", app.updateChannelHandler)
	router.Delete("/v1/channels/{id}", app.deleteChannelHandler)
	return router
}
