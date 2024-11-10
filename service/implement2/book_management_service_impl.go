package implement2

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"library-synapsis/exception"
	"library-synapsis/helper"
	"library-synapsis/helper/model"
	"library-synapsis/model/domain"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
	"library-synapsis/repository"
	"library-synapsis/service"
)

type BookManagementServiceImpl struct {
	BookRepository repository.BookManagementRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookManagementService(bookRepository repository.BookManagementRepository, db *sql.DB, validate *validator.Validate) service.BookManagementService {
	return &BookManagementServiceImpl{
		BookRepository: bookRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *BookManagementServiceImpl) CreateBook(ctx context.Context, request create.BookManagementCreateRequest) response.BookManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book := domain.BookManagement{
		CategoryId:   request.CategoryId,
		AuthorId:     request.AuthorId,
		Judul:        request.Judul,
		CodeBook:     request.CodeBook,
		DateofPublic: request.DateofPublic,
	}

	book = service.BookRepository.CreateBook(ctx, tx, book)
	return model.ToBookResponse(book)
}

func (service *BookManagementServiceImpl) FindByBookId(ctx context.Context, bookId int) response.BookManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindByBookId(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToBookResponse(book)
}

func (service *BookManagementServiceImpl) UpdateBook(ctx context.Context, request update.BookManagementUpdateRequest) response.BookManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindByBookId(ctx, tx, request.BookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	book.CategoryId = request.CategoryId
	book.AuthorId = request.AuthorId
	book.Judul = request.Judul
	book.CodeBook = request.CodeBook
	book.DateofPublic = request.DateofPublic

	book = service.BookRepository.UpdateBook(ctx, tx, book)
	return model.ToBookResponse(book)
}

func (service *BookManagementServiceImpl) DeleteBook(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindByBookId(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.DeleteBook(ctx, tx, book)
}

func (service *BookManagementServiceImpl) FindByBookLikeCriteria(ctx context.Context, judul, authorName string) []response.BookManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	books, err := service.BookRepository.FindByBookLikeCriteria(ctx, tx, judul, authorName)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToBookResponses(books)
}

func (service *BookManagementServiceImpl) FindAllBook(ctx context.Context) []response.BookManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.FindAllBook(ctx, tx)
	return model.ToBookResponses(books)
}
