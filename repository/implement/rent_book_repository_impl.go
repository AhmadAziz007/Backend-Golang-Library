package implement

import (
	"context"
	"database/sql"
	"errors"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
)

type RentBookRepositoryImpl struct{}

func NewRentBookRepository() repository.RentBookRepository {
	return &RentBookRepositoryImpl{}
}

func (repository *RentBookRepositoryImpl) CreateRent(ctx context.Context, tx *sql.Tx, rent domain.RentBook) domain.RentBook {
	SQL := "INSERT INTO rent_Book(book_id, user_id, no_tiket, keterangan, jumlah, date_borrow, date_return) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING rent_id"
	row := tx.QueryRowContext(ctx, SQL,
		rent.BookId,
		rent.UserId,
		rent.NoTiket,
		rent.Keterangan,
		rent.Jumlah,
		rent.DateBorrow,
		rent.DateReturn)
	err := row.Scan(&rent.RentId)
	helper.PanicIfError(err)
	return rent
}

func (repository *RentBookRepositoryImpl) UpdateRent(ctx context.Context, tx *sql.Tx, rent domain.RentBook) domain.RentBook {
	SQL := "UPDATE rent_book SET book_id = $1, user_id = $2, no_tiket = $3, keterangan = $4, jumlah = $5, date_borrow = $6, date_return = $7 WHERE rent_id = $8"
	_, err := tx.ExecContext(ctx, SQL,
		rent.BookId,
		rent.UserId,
		rent.NoTiket,
		rent.Keterangan,
		rent.Jumlah,
		rent.DateBorrow,
		rent.DateReturn,
		rent.RentId)
	helper.PanicIfError(err)
	return rent
}

func (repository *RentBookRepositoryImpl) DeleteRent(ctx context.Context, tx *sql.Tx, rent domain.RentBook) {
	SQL := "DELETE FROM rent_book WHERE rent_id = $1"
	_, err := tx.ExecContext(ctx, SQL, rent.RentId)
	helper.PanicIfError(err)
}

func (repository *RentBookRepositoryImpl) FindByRentId(ctx context.Context, tx *sql.Tx, rentId int) (domain.RentBook, error) {
	SQL := `
         SELECT a.rent_id a.book_id, a.user_id, a.no_tiket, a.status, a.keterangan, a.jumlah, a.date_borrow, a.date_return
         FROM rent_book a
         LEFT JOIN book_management b ON a.book_id = b.book_id
         LEFT JOIN user_management c ON a.user_id = c.user_id
         WHERE a.rent_id = $1 
    `
	row := tx.QueryRowContext(ctx, SQL, rentId)
	rent := domain.RentBook{}
	err := row.Scan(
		&rent.RentId,
		&rent.BookId,
		&rent.UserId,
		&rent.NoTiket,
		&rent.Status,
		&rent.Keterangan,
		&rent.Jumlah,
		&rent.DateBorrow,
		&rent.DateReturn)
	if err != nil {
		if err == sql.ErrNoRows {
			return rent, errors.New("rent book is not found")
		}
		helper.PanicIfError(err)
	}
	return rent, nil
}

func (repository *RentBookRepositoryImpl) FindAllRent(ctx context.Context, tx *sql.Tx) []domain.RentBook {
	SQL := `
         SELECT a.rent_id a.book_id, a.user_id, a.no_tiket, a.status, a.keterangan, a.jumlah, a.date_borrow, a.date_return
         FROM rent_book a
         LEFT JOIN book_management b ON a.book_id = b.book_id
         LEFT JOIN user_management c ON a.user_id = c.user_id
         ORDER BY a.rent_id ASC 
    `
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var rents []domain.RentBook
	for rows.Next() {
		rent := domain.RentBook{}
		err := rows.Scan(
			&rent.RentId,
			&rent.BookId,
			&rent.UserId,
			&rent.NoTiket,
			&rent.Status,
			&rent.Keterangan,
			&rent.Jumlah,
			&rent.DateBorrow,
			&rent.DateReturn)
		helper.PanicIfError(err)
		rents = append(rents, rent)
	}
	return rents
}
