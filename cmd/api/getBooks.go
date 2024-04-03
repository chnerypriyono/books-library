package main

import (
	"net/http"
)

func (app *application) getBooksHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"books": "will place books data here"
	}

	// No need to add acess control origin headers. On other routes, that may be necessary
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
