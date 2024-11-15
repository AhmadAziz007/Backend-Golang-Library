package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByCategoryId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
