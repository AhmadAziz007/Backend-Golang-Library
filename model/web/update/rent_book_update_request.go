package update

import "time"

type RentBookUpdateRequest struct {
	RentId     int       `validate:"required"`
	BookId     int       `validate:"required"`
	UserId     int       `validate:"required"`
	NoTiket    string    `validate:"required,min=1,max=100" json:"noTiket"`
	Status     string    `validate:"required,min=1,max=100" json:"status"`
	Keterangan string    `validate:"required,min=1,max=100" json:"keterangan"`
	Jumlah     string    `validate:"required,min=1,max=100" json:"jumlah"`
	DateBorrow time.Time `validate:"required,min=1,max=100" json:"dateBorrow"`
	DateReturn time.Time `validate:"required,min=1,max=100" json:"dateReturn"`
}
