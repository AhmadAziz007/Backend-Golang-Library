package response

type UserManagementResponse struct {
	UserId   int    `json:"userId"`
	RoleId   int    `json:"roleId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
