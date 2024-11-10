package service

import (
	"context"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
)

type AuthorManagementService interface {
	CreateAuthor(ctx context.Context, request create.AuthorManagementCreateRequest) response.AuthorManagementResponse
	UpdateAuthor(ctx context.Context, request update.AuthorManagementUpdateRequest) response.AuthorManagementResponse
	DeleteAuthor(ctx context.Context, AuthorId int)
	FindByAuthorId(ctx context.Context, AuthorId int) response.AuthorManagementResponse
	FindAllAuthor(ctx context.Context) []response.AuthorManagementResponse
}
