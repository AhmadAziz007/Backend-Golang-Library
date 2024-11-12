package implement3

import (
	"github.com/julienschmidt/httprouter"
	"library-synapsis/helper"
	"library-synapsis/model/web"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/update"
	"library-synapsis/service"
	"net/http"
	"strconv"
)

type RentBookControllerImpl struct {
	RentBookService service.RentBookService
}

func NewRentBookController(rentBookService service.RentBookService) *RentBookControllerImpl {
	return &RentBookControllerImpl{
		RentBookService: rentBookService,
	}
}

func (controller *RentBookControllerImpl) CreateRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rentBookCreateRequest := create.RentBookCreateRequest{}
	helper.ReadFromRequestBody(request, &rentBookCreateRequest)

	rentBookResponse := controller.RentBookService.CreateRent(request.Context(), rentBookCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rentBookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RentBookControllerImpl) UpdateRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rentBookUpdateRequest := update.RentBookUpdateRequest{}
	helper.ReadFromRequestBody(request, &rentBookUpdateRequest)

	rentId := params.ByName("rentId")
	id, err := strconv.Atoi(rentId)
	helper.PanicIfError(err)

	rentBookUpdateRequest.RentId = id

	rentBookResponse := controller.RentBookService.UpdateRent(request.Context(), rentBookUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rentBookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RentBookControllerImpl) FindByRentId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rentId := params.ByName("rentId")
	id, err := strconv.Atoi(rentId)
	helper.PanicIfError(err)

	rentBookResponse := controller.RentBookService.FindByRentId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rentBookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RentBookControllerImpl) DeleteRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rentId := params.ByName("rentId")
	id, err := strconv.Atoi(rentId)
	helper.PanicIfError(err)

	controller.RentBookService.DeleteRent(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Rent deleted successfully",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RentBookControllerImpl) FindAllRent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rentBookResponse := controller.RentBookService.FindAllRent(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rentBookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
	writer.Write([]byte("List of Rent"))
}
