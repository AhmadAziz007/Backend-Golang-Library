package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthorManagementController interface {
	CreateAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByAuthorId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
