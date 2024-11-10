package update

type CategoryUpdateRequest struct {
	CategoryId   int    `validate:"required"`
	CategoryName string `validate:"required,max=200,min=1" json:"categoryName"`
}
