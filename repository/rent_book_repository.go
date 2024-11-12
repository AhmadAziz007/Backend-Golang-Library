package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type RentBookRepository interface {
	CreateRent(ctx context.Context, tx *sql.Tx, stock domain.RentBook) domain.RentBook
	UpdateRent(ctx context.Context, tx *sql.Tx, stock domain.RentBook) domain.RentBook
	DeleteRent(ctx context.Context, tx *sql.Tx, stock domain.RentBook)
	FindByRentId(ctx context.Context, tx *sql.Tx, stockId int) (domain.RentBook, error)
	FindAllRent(ctx context.Context, tx *sql.Tx) []domain.RentBook
}
