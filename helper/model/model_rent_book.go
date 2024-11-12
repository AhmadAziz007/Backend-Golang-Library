package model

import (
	"library-synapsis/model/domain"
	"library-synapsis/model/web/response"
)

func ToRentBookResponse(rentBook domain.RentBook) response.RentBookResponse {
	return response.RentBookResponse{
		RentId:     rentBook.RentId,
		BookId:     rentBook.BookId,
		UserId:     rentBook.UserId,
		NoTiket:    rentBook.NoTiket,
		Status:     rentBook.Status,
		Keterangan: rentBook.Keterangan,
		Jumlah:     rentBook.Jumlah,
		DateBorrow: rentBook.DateBorrow,
		DateReturn: rentBook.DateReturn,
	}
}

func ToRentBookResponses(rentBooks []domain.RentBook) []response.RentBookResponse {
	var rentResponses []response.RentBookResponse
	for _, rentBook := range rentBooks {
		rentResponses = append(rentResponses, ToRentBookResponse(rentBook))
	}
	return rentResponses
}
