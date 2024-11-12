package create

type LoginCreateRequest struct {
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
