package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserManagementController interface {
	CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
