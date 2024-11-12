package create

type StockManagementCreateRequest struct {
	BookId int `validate:"required"`
	Stock  int `validate:"required" json:"stock"`
}
