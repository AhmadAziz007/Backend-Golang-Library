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

type StockManagementControllerImpl struct {
	StockManagementService service.StockManagementService
}

func NewStockManagementController(stockManagementService service.StockManagementService) *StockManagementControllerImpl {
	return &StockManagementControllerImpl{
		StockManagementService: stockManagementService,
	}
}

func (controller *StockManagementControllerImpl) CreateStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stockManagementCreateRequest := create.StockManagementCreateRequest{}
	helper.ReadFromRequestBody(request, &stockManagementCreateRequest)

	stockManagementResponse := controller.StockManagementService.CreateStock(request.Context(), stockManagementCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stockManagementResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockManagementControllerImpl) UpdateStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stockManagementUpdateRequest := update.StockManagementUpdateRequest{}
	helper.ReadFromRequestBody(request, &stockManagementUpdateRequest)

	stockId := params.ByName("stockId")
	id, err := strconv.Atoi(stockId)
	helper.PanicIfError(err)

	stockManagementUpdateRequest.StockId = id
	stockManagementResponse := controller.StockManagementService.UpdateStock(request.Context(), stockManagementUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stockManagementResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockManagementControllerImpl) DeleteStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stockId := params.ByName("stockId")
	id, err := strconv.Atoi(stockId)
	helper.PanicIfError(err)

	controller.StockManagementService.DeleteStock(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Stock deleted successfully",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockManagementControllerImpl) FindByStockId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stockId := params.ByName("stockId")
	id, err := strconv.Atoi(stockId)
	helper.PanicIfError(err)

	stockResponse := controller.StockManagementService.FindByStockId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stockResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StockManagementControllerImpl) FindAllStock(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stocks := controller.StockManagementService.FindAllStock(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   stocks,
	}
	helper.WriteToResponseBody(writer, webResponse)
	writer.Write([]byte("List of Stock"))
}
