package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookManagementController interface {
	CreateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByBookId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByBookLikeCriteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
