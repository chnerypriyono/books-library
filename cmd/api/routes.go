package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/getBooks", app.getBooksHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/deleteBook", app.deleteBookHandler)
	router.HandlerFunc(http.MethodPut, "/v1/updateBook", app.updateBookHandler)
	router.HandlerFunc(http.MethodPost, "/v1/createBook", app.createBookHandler)

	return router
}
