package create

type CategoryCreateRequest struct {
	CategoryName string `validate:"required,min=1,max=100" json:"categoryName"`
}
