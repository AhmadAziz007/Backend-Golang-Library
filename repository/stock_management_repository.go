package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type StockManagementRepository interface {
	CreateStock(ctx context.Context, tx *sql.Tx, rent domain.StockManagement) domain.StockManagement
	UpdateStock(ctx context.Context, tx *sql.Tx, rent domain.StockManagement) domain.StockManagement
	DeleteStock(ctx context.Context, tx *sql.Tx, rent domain.StockManagement)
	FindByStockId(ctx context.Context, tx *sql.Tx, rentId int) (domain.StockManagement, error)
	FindAllStock(ctx context.Context, tx *sql.Tx) []domain.StockManagement
}
