package create

type RoleManagementCreateRequest struct {
	RoleName string `validate:"required,min=1,max=100" json:"roleName"`
}
