package main

import (
	"fmt"
	"net/http"
)

func (app *application) createChannelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new channel")
}

func (app *application) showChannelHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of channel %d\n", id)
}
