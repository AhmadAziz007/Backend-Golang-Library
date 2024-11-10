package update

type RoleManagementUpdateRequest struct {
	RoleId   int    `validate:"required",`
	RoleName string `validate:"required,min=1,max=100" json:"roleName"`
}
