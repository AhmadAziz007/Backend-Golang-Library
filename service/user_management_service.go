package service

import (
	"context"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
)

type UserManagementService interface {
	CreateUser(ctx context.Context, request create.UserManagementCreateRequest) response.UserManagementResponse
	UpdateUser(ctx context.Context, request update.UserManagementUpdateRequest) response.UserManagementResponse
	DeleteUser(ctx context.Context, categoryId int)
	FindByUserId(ctx context.Context, categoryId int) response.UserManagementResponse
	FindAllUser(ctx context.Context) []response.UserManagementResponse
}
