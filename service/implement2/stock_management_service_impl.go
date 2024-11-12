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

type StockManagementServiceImpl struct {
	StockRepository repository.StockManagementRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewStockManagementService(stockRepository repository.StockManagementRepository, db *sql.DB, validate *validator.Validate) service.StockManagementService {
	return &StockManagementServiceImpl{
		StockRepository: stockRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service *StockManagementServiceImpl) CreateStock(ctx context.Context, request create.StockManagementCreateRequest) response.StockManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stock := domain.StockManagement{
		BookId: request.BookId,
		Stock:  request.Stock,
	}

	stock = service.StockRepository.CreateStock(ctx, tx, stock)
	return model.ToStockResponse(stock)
}

func (service *StockManagementServiceImpl) UpdateStock(ctx context.Context, request update.StockManagementUpdateRequest) response.StockManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stock, err := service.StockRepository.FindByStockId(ctx, tx, request.StockId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	stock.StockId = request.StockId
	stock.BookId = request.BookId
	stock.Stock = request.Stock

	stock = service.StockRepository.UpdateStock(ctx, tx, stock)
	return model.ToStockResponse(stock)
}

func (service *StockManagementServiceImpl) DeleteStock(ctx context.Context, stockId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stock, err := service.StockRepository.FindByStockId(ctx, tx, stockId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.StockRepository.DeleteStock(ctx, tx, stock)
}

func (service *StockManagementServiceImpl) FindByStockId(ctx context.Context, stockId int) response.StockManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stock, err := service.StockRepository.FindByStockId(ctx, tx, stockId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToStockResponse(stock)
}

func (service *StockManagementServiceImpl) FindAllStock(ctx context.Context) []response.StockManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	stock := service.StockRepository.FindAllStock(ctx, tx)
	return model.ToStockResponses(stock)
}
