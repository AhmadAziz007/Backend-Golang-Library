package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindByCategoryId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAllCategory(ctx context.Context, tx *sql.Tx) []domain.Category
}
