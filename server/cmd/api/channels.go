package main

import (
	"fmt"
	"forum/internal/data"
	"net/http"
	"time"
)

func (app *application) createChannelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new channel")
}

func (app *application) showChannelHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	channel := data.Channel{
		ID:        int64(id),
		UserID:    1,
		CreatedAt: time.Now(),
		Title:     "Programming",
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"channel": channel}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
