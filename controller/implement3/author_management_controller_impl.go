package implement3

import (
	"library-synapsis/helper"
	"library-synapsis/model/web"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/update"
	"library-synapsis/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AuthorManagementControllerImpl struct {
	AuthorManagementService service.AuthorManagementService
}

func NewAuthorManagementController(authorManagementService service.AuthorManagementService) *AuthorManagementControllerImpl {
	return &AuthorManagementControllerImpl{
		AuthorManagementService: authorManagementService,
	}
}

func (controller *AuthorManagementControllerImpl) CreateAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authorManagementCreateRequest := create.AuthorManagementCreateRequest{}
	helper.ReadFromRequestBody(request, &authorManagementCreateRequest)

	authorManagementResponse := controller.AuthorManagementService.CreateAuthor(request.Context(), authorManagementCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authorManagementResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthorManagementControllerImpl) UpdateAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authorManagementUpdateRequest := update.AuthorManagementUpdateRequest{}
	helper.ReadFromRequestBody(request, &authorManagementUpdateRequest)

	authorId := params.ByName("authorId")
	id, err := strconv.Atoi(authorId)
	helper.PanicIfError(err)

	authorManagementUpdateRequest.AuthorId = id

	authorManagementResponse := controller.AuthorManagementService.UpdateAuthor(request.Context(), authorManagementUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authorManagementResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthorManagementControllerImpl) DeleteAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authorId := params.ByName("authorId")
	id, err := strconv.Atoi(authorId)
	helper.PanicIfError(err)

	controller.AuthorManagementService.DeleteAuthor(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthorManagementControllerImpl) FindByAuthorId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authorId := params.ByName("authorId")
	id, err := strconv.Atoi(authorId)
	helper.PanicIfError(err)

	authorResponse := controller.AuthorManagementService.FindByAuthorId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthorManagementControllerImpl) FindAllAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authorResponse := controller.AuthorManagementService.FindAllAuthor(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
