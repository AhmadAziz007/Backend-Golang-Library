package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type RoleManagementRepository interface {
	CreateRole(ctx context.Context, tx *sql.Tx, role domain.RoleManagement) domain.RoleManagement
	UpdateRole(ctx context.Context, tx *sql.Tx, role domain.RoleManagement) domain.RoleManagement
	DeleteRole(ctx context.Context, tx *sql.Tx, role domain.RoleManagement)
	FindByRoleId(ctx context.Context, tx *sql.Tx, roleId int) (domain.RoleManagement, error)
	FindByRoleName(ctx context.Context, tx *sql.Tx, roleName string) (domain.RoleManagement, error)
	FindAllRole(ctx context.Context, tx *sql.Tx) []domain.RoleManagement
}
