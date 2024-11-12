package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RentBookController interface {
	CreateRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByRentId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
