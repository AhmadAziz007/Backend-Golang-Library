package domain

import "time"

type RentBook struct {
	RentId     int
	BookId     int
	UserId     int
	NoTiket    string
	Status     string
	Keterangan string
	Jumlah     string
	DateBorrow time.Time
	DateReturn time.Time
}
