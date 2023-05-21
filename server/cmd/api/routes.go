package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Get("/v1/channels", app.listChannelsHandler)
	router.Post("/v1/channels", app.createChannelHandler)
	router.Get("/v1/channels/{id}", app.showChannelHandler)
	router.Patch("/v1/channels/{id}", app.updateChannelHandler)
	router.Delete("/v1/channels/{id}", app.deleteChannelHandler)

	router.Get("/v1/threads", app.listThreadsHandler)
	router.Post("/v1/threads", app.createThreadHandler)
	router.Get("/v1/threads/{id}", app.showThreadHandler)
	router.Patch("/v1/threads/{id}", app.updateThreadHandler)
	router.Delete("/v1/threads/{id}", app.deleteThreadHandler)

	router.Post("/v1/users", app.registerUserHandler)

	return app.recoverPanic(app.rateLimit(router))
}
