package response

type StockManagementResponse struct {
	StockId int    `json:"stockId"`
	BookId  int    `json:"bookId"`
	Judul   string `json:"judul"`
	Stock   int    `json:"stock"`
}
