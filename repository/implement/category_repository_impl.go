package implement

import (
	"context"
	"database/sql"
	"errors"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() repository.CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(category_name) values ($1) returning category_id"
	row := tx.QueryRowContext(ctx, SQL, category.CategoryName)
	err := row.Scan(&category.CategoryId)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set category_name = $1 where category_id = $2"
	_, err := tx.ExecContext(ctx, SQL, category.CategoryName, category.CategoryId)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where category_id = $1"
	_, err := tx.ExecContext(ctx, SQL, category.CategoryId)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindByCategoryId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select category_id, category_name from category where category_id = $1"
	row := tx.QueryRowContext(ctx, SQL, categoryId)

	category := domain.Category{}
	err := row.Scan(&category.CategoryId, &category.CategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return category, errors.New("category is not found")
		}
		helper.PanicIfError(err)
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAllCategory(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select category_id, category_name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.CategoryId, &category.CategoryName)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
