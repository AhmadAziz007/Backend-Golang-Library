package model

import (
	"library-synapsis/model/domain"
	"library-synapsis/model/web/response"
)

func ToBookResponse(book domain.BookManagement) response.BookManagementResponse {
	return response.BookManagementResponse{
		BookId:       book.BookId,
		CategoryId:   book.CategoryId,
		CategoryName: book.CategoryName,
		AuthorId:     book.AuthorId,
		AuthorName:   book.AuthorName,
		Judul:        book.Judul,
		CodeBook:     book.CodeBook,
		DateofPublic: book.DateofPublic,
	}
}

func ToBookResponses(books []domain.BookManagement) []response.BookManagementResponse {
	var bookResponses []response.BookManagementResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}
	return bookResponses
}
