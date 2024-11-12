package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type LoginRepository interface {
	FindByRoleName(ctx context.Context, tx *sql.Tx, roleId int) (domain.RoleManagement, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, password string, email string) (domain.UserManagement, error)
}
