package create

type AuthorManagementCreateRequest struct {
	AuthorName string `validate:"required,min=1,max=100" json:"authorName"`
}
