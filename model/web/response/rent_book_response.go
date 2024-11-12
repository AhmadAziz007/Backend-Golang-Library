package response

import "time"

type RentBookResponse struct {
	RentId     int       `json:"rentId"`
	BookId     int       `json:"bookId""`
	UserId     int       `json:"userId"`
	NoTiket    string    `json:"noTiket"`
	Status     string    `json:"status"`
	Keterangan string    `json:"keterangan"`
	Jumlah     string    `json:"jumlah"`
	DateBorrow time.Time `json:"dateBorrow"`
	DateReturn time.Time `json:"dateReturn"`
}
