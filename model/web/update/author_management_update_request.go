package update

type AuthorManagementUpdateRequest struct {
	AuthorId   int    `validate:"required"`
	AuthorName string `validate:"required,min=1,max=100" json:"authorName"`
}
