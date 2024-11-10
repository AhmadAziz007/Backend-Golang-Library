package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type AuthorManagementRepository interface {
	CreateAuthor(ctx context.Context, tx *sql.Tx, author domain.AuthorManagement) domain.AuthorManagement
	UpdateAuthor(ctx context.Context, tx *sql.Tx, author domain.AuthorManagement) domain.AuthorManagement
	DeleteAuthor(ctx context.Context, tx *sql.Tx, author domain.AuthorManagement)
	FindByAuthorId(ctx context.Context, db *sql.DB, AuthorId int) (domain.AuthorManagement, error)
	FindAllAuthor(ctx context.Context, tx *sql.Tx) []domain.AuthorManagement
}
