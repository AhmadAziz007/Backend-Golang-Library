package update

type StockManagementUpdateRequest struct {
	StockId int `validate:"required"`
	BookId  int `validate:"required"`
	Stock   int `validate:"required" json:"stock"`
}
