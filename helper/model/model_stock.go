package model

import (
	"library-synapsis/model/domain"
	"library-synapsis/model/web/response"
)

func ToStockResponse(stock domain.StockManagement) response.StockManagementResponse {
	return response.StockManagementResponse{
		StockId: stock.StockId,
		BookId:  stock.BookId,
		Stock:   stock.Stock,
	}
}

func ToStockResponses(stocks []domain.StockManagement) []response.StockManagementResponse {
	var stockResponses []response.StockManagementResponse
	for _, stock := range stocks {
		stockResponses = append(stockResponses, ToStockResponse(stock))
	}
	return stockResponses
}
