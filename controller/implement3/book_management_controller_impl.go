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

type BookManagementControllerImpl struct {
	BookManagementService service.BookManagementService
}

func NewBookManagementController(BookManagementService service.BookManagementService) *BookManagementControllerImpl {
	return &BookManagementControllerImpl{
		BookManagementService: BookManagementService,
	}
}

func (controller *BookManagementControllerImpl) CreateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookManagementCreateRequest := create.BookManagementCreateRequest{}
	helper.ReadFromRequestBody(request, &bookManagementCreateRequest)

	bookManagementResponse := controller.BookManagementService.CreateBook(request.Context(), bookManagementCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookManagementResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookManagementControllerImpl) FindByBookId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookManagementService.FindByBookId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookManagementControllerImpl) UpdateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookManagementUpdateRequest := update.BookManagementUpdateRequest{}
	helper.ReadFromRequestBody(request, &bookManagementUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookManagementUpdateRequest.BookId = id

	bookManagementResponse := controller.BookManagementService.UpdateBook(request.Context(), bookManagementUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookManagementResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookManagementControllerImpl) DeleteBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookManagementService.DeleteBook(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Book deleted successfully",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookManagementControllerImpl) FindByBookLikeCriteria(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var criteriaRequest struct {
		Judul      string `json:"judul"`
		AuthorName string `json:"authorName"`
	}

	helper.ReadFromRequestBody(request, &criteriaRequest)
	bookResponses := controller.BookManagementService.FindByBookLikeCriteria(request.Context(), criteriaRequest.Judul, criteriaRequest.AuthorName)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookManagementControllerImpl) FindAllBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	books := controller.BookManagementService.FindAllBook(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   books,
	}
	helper.WriteToResponseBody(writer, webResponse)
	writer.Write([]byte("List of books"))
}
