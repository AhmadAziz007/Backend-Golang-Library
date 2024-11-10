package update

type UserManagementUpdateRequest struct {
	UserId   int    `validate:"required"`
	RoleId   int    `validate:"required"`
	UserName string `validate:"required,min=1,max=100" json:"userName"`
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
