package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type BookManagementRepository interface {
	CreateBook(ctx context.Context, tx *sql.Tx, book domain.BookManagement) domain.BookManagement
	UpdateBook(ctx context.Context, tx *sql.Tx, book domain.BookManagement) domain.BookManagement
	DeleteBook(ctx context.Context, tx *sql.Tx, book domain.BookManagement)
	FindByBookId(ctx context.Context, tx *sql.Tx, bookId int) (domain.BookManagement, error)
	FindByBookLikeCriteria(ctx context.Context, tx *sql.Tx, judul, authorName string) ([]domain.BookManagement, error)
	FindAllBook(ctx context.Context, tx *sql.Tx) []domain.BookManagement
}
