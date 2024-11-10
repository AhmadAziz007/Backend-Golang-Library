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

type UserManagementControllerImpl struct {
	UserManagementService service.UserManagementService
}

func NewUserManegementController(UserManagementService service.UserManagementService) *UserManagementControllerImpl {
	return &UserManagementControllerImpl{
		UserManagementService: UserManagementService,
	}
}

func (controller *UserManagementControllerImpl) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userManagementCreateRequest := create.UserManagementCreateRequest{}
	helper.ReadFromRequestBody(request, &userManagementCreateRequest)

	userManagementResponse := controller.UserManagementService.CreateUser(request.Context(), userManagementCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userManagementResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserManagementControllerImpl) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userManagementUpdateRequest := update.UserManagementUpdateRequest{}
	helper.ReadFromRequestBody(request, &userManagementUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userManagementUpdateRequest.UserId = id

	userManagementResponse := controller.UserManagementService.UpdateUser(request.Context(), userManagementUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userManagementResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserManagementControllerImpl) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserManagementService.DeleteUser(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Book deleted successfully",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserManagementControllerImpl) FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserManagementService.FindByUserId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserManagementControllerImpl) FindAllUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	user := controller.UserManagementService.FindAllUser(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
