package implement

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
)

type AuthorManagementRepositoryImpl struct{}

func NewAuthorManagementRepository() repository.AuthorManagementRepository {
	return &AuthorManagementRepositoryImpl{}
}

func (repository *AuthorManagementRepositoryImpl) CreateAuthor(ctx context.Context, tx *sql.Tx, author domain.AuthorManagement) domain.AuthorManagement {
	SQL := "insert into author_management(author_name) values ($1) returning author_id"
	row := tx.QueryRowContext(ctx, SQL, author.AuthorName)
	err := row.Scan(&author.AuthorId)
	helper.PanicIfError(err)

	return author
}

func (repository *AuthorManagementRepositoryImpl) UpdateAuthor(ctx context.Context, tx *sql.Tx, author domain.AuthorManagement) domain.AuthorManagement {
	SQL := "update author_management set author_name = $1 where author_id = $2"
	_, err := tx.ExecContext(ctx, SQL, author.AuthorName, author.AuthorId)
	helper.PanicIfError(err)

	return author
}

func (repository *AuthorManagementRepositoryImpl) DeleteAuthor(ctx context.Context, tx *sql.Tx, author domain.AuthorManagement) {
	SQL := "delete from author_management where author_id = $1"
	_, err := tx.ExecContext(ctx, SQL, author.AuthorId)
	helper.PanicIfError(err)
}

func (repository *AuthorManagementRepositoryImpl) FindByAuthorId(ctx context.Context, db *sql.DB, authorId int) (domain.AuthorManagement, error) {
	SQL := "SELECT author_id, author_name FROM author_management WHERE author_id = $1"

	row := db.QueryRowContext(ctx, SQL, authorId)

	author := domain.AuthorManagement{}

	err := row.Scan(&author.AuthorId, &author.AuthorName)

	if err != nil {
		if err == sql.ErrNoRows {
			return author, errors.New("author not found")
		}
		return author, fmt.Errorf("gagal mengeksekusi query: %v", err)
	}

	return author, nil
}

func (repository *AuthorManagementRepositoryImpl) FindAllAuthor(ctx context.Context, tx *sql.Tx) []domain.AuthorManagement {
	SQL := "SELECT author_id, author_name FROM author_management"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var authors []domain.AuthorManagement
	for rows.Next() {
		author := domain.AuthorManagement{}
		err := rows.Scan(&author.AuthorId, &author.AuthorName)
		helper.PanicIfError(err)
		authors = append(authors, author)
	}
	return authors
}
