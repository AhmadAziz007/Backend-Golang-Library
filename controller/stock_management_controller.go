package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type StockManagementController interface {
	CreateStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByStockId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
