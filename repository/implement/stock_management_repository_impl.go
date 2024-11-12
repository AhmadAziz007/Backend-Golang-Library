package implement

import (
	"context"
	"database/sql"
	"errors"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
)

type StockManagementRepositoryImpl struct{}

func NewStockManagementRepository() repository.StockManagementRepository {
	return &StockManagementRepositoryImpl{}
}

func (StockManagementRepositoryImpl) CreateStock(ctx context.Context, tx *sql.Tx, stock domain.StockManagement) domain.StockManagement {
	SQL := "INSERT INTO stock_management(book_id, stock) VALUES ($1, $2) RETURNING book_id"
	row := tx.QueryRowContext(ctx, SQL, stock.BookId, stock.Stock)
	err := row.Scan(&stock.StockId)
	helper.PanicIfError(err)
	return stock
}

func (StockManagementRepositoryImpl) UpdateStock(ctx context.Context, tx *sql.Tx, stock domain.StockManagement) domain.StockManagement {
	SQL := "UPDATE stock_management SET book_id = $1, stock = $2 WHERE book_id = $3"
	_, err := tx.ExecContext(ctx, SQL,
		stock.BookId,
		stock.Stock,
		stock.StockId)
	helper.PanicIfError(err)
	return stock
}

func (StockManagementRepositoryImpl) DeleteStock(ctx context.Context, tx *sql.Tx, stock domain.StockManagement) {
	SQL := "DELETE FROM stock_management WHERE stock_id = $1"
	_, err := tx.ExecContext(ctx, SQL, stock.StockId)
	helper.PanicIfError(err)
}

func (StockManagementRepositoryImpl) FindByStockId(ctx context.Context, tx *sql.Tx, stockId int) (domain.StockManagement, error) {
	SQL := `
        SELECT a.stock_id, a.book_id, b.judul, a.stock 
        FROM stock_management a 
        LEFT JOIN book_management b ON a.book_id = b.book_id
        WHERE a.stock_id = $1
        ORDER BY a.stock_id ASC
    `
	row := tx.QueryRowContext(ctx, SQL, stockId)
	stock := domain.StockManagement{}
	err := row.Scan(
		&stock.StockId,
		&stock.BookId,
		&stock.Judul,
		&stock.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return stock, errors.New("stock is not found")
		}
		helper.PanicIfError(err)
	}
	return stock, nil
}

func (StockManagementRepositoryImpl) FindAllStock(ctx context.Context, tx *sql.Tx) []domain.StockManagement {
	SQL := `
        SELECT a.stock_id, a.book_id, b.judul, a.stock 
        FROM stock_management a 
        LEFT JOIN book_management b ON a.book_id = b.book_id
        ORDER BY a.stock_id ASC
    `
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var stocks []domain.StockManagement
	for rows.Next() {
		stock := domain.StockManagement{}
		err := rows.Scan(
			&stock.StockId,
			&stock.BookId,
			&stock.Judul,
			&stock.Stock)
		helper.PanicIfError(err)
		stocks = append(stocks, stock)
	}
	return stocks
}
