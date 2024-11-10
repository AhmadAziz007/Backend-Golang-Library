package service

import (
	"context"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
)

type BookManagementService interface {
	CreateBook(ctx context.Context, request create.BookManagementCreateRequest) response.BookManagementResponse
	UpdateBook(ctx context.Context, request update.BookManagementUpdateRequest) response.BookManagementResponse
	DeleteBook(ctx context.Context, bookId int)
	FindByBookId(ctx context.Context, bookId int) response.BookManagementResponse
	FindByBookLikeCriteria(ctx context.Context, judul, authorName string) []response.BookManagementResponse
	FindAllBook(ctx context.Context) []response.BookManagementResponse
}
