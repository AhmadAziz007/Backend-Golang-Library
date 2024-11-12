package implement2

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"library-synapsis/exception"
	"library-synapsis/helper"
	"library-synapsis/helper/model"
	"library-synapsis/model/domain"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
	"library-synapsis/repository"
	"library-synapsis/service"
)

type RentBookServiceImpl struct {
	RentBookRepository repository.RentBookRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewRentBookService(rentBookRepository repository.RentBookRepository, db *sql.DB, validate *validator.Validate) service.RentBookService {
	return &RentBookServiceImpl{
		RentBookRepository: rentBookRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *RentBookServiceImpl) CreateRent(ctx context.Context, request create.RentBookCreateRequest) response.RentBookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rent := domain.RentBook{
		BookId:     request.BookId,
		UserId:     request.UserId,
		NoTiket:    request.NoTiket,
		Keterangan: request.Keterangan,
		Jumlah:     request.Jumlah,
		DateBorrow: request.DateBorrow,
		DateReturn: request.DateReturn,
	}

	rent = service.RentBookRepository.CreateRent(ctx, tx, rent)
	return model.ToRentBookResponse(rent)
}

func (service *RentBookServiceImpl) UpdateRent(ctx context.Context, request update.RentBookUpdateRequest) response.RentBookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rent, err := service.RentBookRepository.FindByRentId(ctx, tx, request.RentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	rent.BookId = request.BookId
	rent.UserId = request.UserId
	rent.NoTiket = request.NoTiket
	rent.Keterangan = request.Keterangan
	rent.Jumlah = request.Jumlah
	rent.DateBorrow = request.DateBorrow
	rent.DateReturn = request.DateReturn

	rent = service.RentBookRepository.UpdateRent(ctx, tx, rent)
	return model.ToRentBookResponse(rent)
}

func (service *RentBookServiceImpl) DeleteRent(ctx context.Context, rentId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rent, err := service.RentBookRepository.FindByRentId(ctx, tx, rentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.RentBookRepository.DeleteRent(ctx, tx, rent)
}

func (service *RentBookServiceImpl) FindByRentId(ctx context.Context, rentId int) response.RentBookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rent, err := service.RentBookRepository.FindByRentId(ctx, tx, rentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToRentBookResponse(rent)
}

func (service *RentBookServiceImpl) FindAllRent(ctx context.Context) []response.RentBookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rents := service.RentBookRepository.FindAllRent(ctx, tx)
	return model.ToRentBookResponses(rents)
}
