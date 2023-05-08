package main

import (
	"fmt"
	"forum/internal/data"
	"forum/internal/validator"
	"net/http"
	"time"
)

func (app *application) createChannelHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int64  `json:"user_id"`
		Title  string `json:"title"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	channel := &data.Channel{
		Title:     input.Title,
		UserID:    input.UserID,
		CreatedAt: time.Now(),
		Version:   1,
	}

	v := validator.New()

	if data.ValidateChannel(v, channel); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
