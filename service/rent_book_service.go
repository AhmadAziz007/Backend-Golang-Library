package service

import (
	"context"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
)

type RentBookService interface {
	CreateRent(ctx context.Context, request create.RentBookCreateRequest) response.RentBookResponse
	UpdateRent(ctx context.Context, request update.RentBookUpdateRequest) response.RentBookResponse
	DeleteRent(ctx context.Context, rentId int)
	FindByRentId(ctx context.Context, rentId int) response.RentBookResponse
	FindAllRent(ctx context.Context) []response.RentBookResponse
}
