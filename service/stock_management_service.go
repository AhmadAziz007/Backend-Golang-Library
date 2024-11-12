package service

import (
	"context"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
)

type StockManagementService interface {
	CreateStock(ctx context.Context, request create.StockManagementCreateRequest) response.StockManagementResponse
	UpdateStock(ctx context.Context, request update.StockManagementUpdateRequest) response.StockManagementResponse
	DeleteStock(ctx context.Context, stockId int)
	FindByStockId(ctx context.Context, stockId int) response.StockManagementResponse
	FindAllStock(ctx context.Context) []response.StockManagementResponse
}
