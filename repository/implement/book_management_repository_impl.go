package implement

import (
	"context"
	"database/sql"
	"errors"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
)

type BookManagementRepositoryImpl struct{}

func NewBookManagementRepository() repository.BookManagementRepository {
	return &BookManagementRepositoryImpl{}
}

func (repository *BookManagementRepositoryImpl) CreateBook(ctx context.Context, tx *sql.Tx, book domain.BookManagement) domain.BookManagement {
	SQL := "INSERT INTO book_management(category_id, author_id, judul, code_book, date_of_publication) VALUES ($1, $2, $3, $4, $5) RETURNING book_id"
	row := tx.QueryRowContext(ctx, SQL, book.CategoryId, book.AuthorId, book.Judul, book.CodeBook, book.DateofPublic)
	err := row.Scan(&book.BookId)
	helper.PanicIfError(err)
	return book
}

func (repository *BookManagementRepositoryImpl) FindByBookId(ctx context.Context, tx *sql.Tx, bookId int) (domain.BookManagement, error) {
	SQL := `
		SELECT a.book_id, a.category_id, a.author_id, d.author_name, a.code_book, a.date_of_publication
		FROM book_management a
		LEFT JOIN category c ON a.category_id = c.category_id
		LEFT JOIN author_management d ON a.author_id = d.author_id
		WHERE a.book_id = $1
	`
	row := tx.QueryRowContext(ctx, SQL, bookId)
	book := domain.BookManagement{}
	err := row.Scan(
		&book.BookId,
		&book.CategoryId,
		&book.AuthorId,
		&book.Judul,
		&book.CodeBook,
		&book.DateofPublic)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("book is not found")
		}
		helper.PanicIfError(err)
	}
	return book, nil
}

func (repository *BookManagementRepositoryImpl) UpdateBook(ctx context.Context, tx *sql.Tx, book domain.BookManagement) domain.BookManagement {
	SQL := "UPDATE book_management SET category_id = $1, author_id = $2, judul = $3, code_book = $4, date_of_publication = $5 WHERE book_id = $6"
	_, err := tx.ExecContext(ctx, SQL,
		book.CategoryId,
		book.AuthorId,
		book.Judul,
		book.CodeBook,
		book.DateofPublic,
		book.BookId)
	helper.PanicIfError(err)
	return book
}

func (repository *BookManagementRepositoryImpl) DeleteBook(ctx context.Context, tx *sql.Tx, book domain.BookManagement) {
	SQL := "DELETE FROM book_management WHERE book_id = $1"
	_, err := tx.ExecContext(ctx, SQL, book.BookId)
	helper.PanicIfError(err)
}

func (repository *BookManagementRepositoryImpl) FindByBookLikeCriteria(ctx context.Context, tx *sql.Tx, judul, authorName string) ([]domain.BookManagement, error) {
	SQL := `
		SELECT a.book_id, a.category_id, c.category_name, a.author_id, d.author_name, a.judul, a.code_book, a.date_of_publication
		FROM book_management a
		LEFT JOIN category c ON a.category_id = c.category_id
		LEFT JOIN author_management d ON a.author_id = d.author_id
		WHERE a.judul LIKE $1 AND d.author_name LIKE $2
		ORDER BY a.book_id ASC
	`
	rows, err := tx.QueryContext(ctx, SQL, "%"+judul+"%", "%"+authorName+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.BookManagement
	for rows.Next() {
		book := domain.BookManagement{}
		err := rows.Scan(
			&book.BookId,
			&book.CategoryId,
			&book.CategoryName,
			&book.AuthorId,
			&book.AuthorName,
			&book.Judul,
			&book.CodeBook,
			&book.DateofPublic)
		helper.PanicIfError(err)
		books = append(books, book)
	}
	return books, nil
}

func (BookManagementRepositoryImpl) FindAllBook(ctx context.Context, tx *sql.Tx) []domain.BookManagement {
	SQL := `
		SELECT a.book_id, a.category_id, c.category_name, a.author_id, d.author_name, a.judul, a.code_book, a.date_of_publication
		FROM book_management a
		LEFT JOIN category c ON a.category_id = c.category_id
		LEFT JOIN author_management d ON a.author_id = d.author_id
		ORDER BY a.book_id ASC
	`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.BookManagement
	for rows.Next() {
		book := domain.BookManagement{}
		err := rows.Scan(
			&book.BookId,
			&book.CategoryId,
			&book.CategoryName,
			&book.AuthorId,
			&book.AuthorName,
			&book.Judul,
			&book.CodeBook,
			&book.DateofPublic)
		helper.PanicIfError(err)
		books = append(books, book)
	}
	return books
}
