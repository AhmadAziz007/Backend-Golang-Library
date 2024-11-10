package service

import (
	"context"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
)

type CategoryService interface {
	Create(ctx context.Context, request create.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request update.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindByCategoryId(ctx context.Context, categoryId int) response.CategoryResponse
	FindAllCategory(ctx context.Context) []response.CategoryResponse
}
