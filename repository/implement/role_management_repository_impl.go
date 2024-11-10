package implement

import (
	"context"
	"database/sql"
	"errors"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
)

type RoleManagementRepositoryImpl struct{}

func NewRoleManagementRepository() repository.RoleManagementRepository {
	return &RoleManagementRepositoryImpl{}
}

func (repository *RoleManagementRepositoryImpl) CreateRole(ctx context.Context, tx *sql.Tx, role domain.RoleManagement) domain.RoleManagement {
	SQL := "insert into role_management(role_name) values ($1) returning role_id"
	row := tx.QueryRowContext(ctx, SQL, role.RoleName)
	err := row.Scan(&role.RoleId)
	helper.PanicIfError(err)

	return role
}

func (repository *RoleManagementRepositoryImpl) UpdateRole(ctx context.Context, tx *sql.Tx, role domain.RoleManagement) domain.RoleManagement {
	SQL := "update role_management set role_name = $1 where role_id = $2"
	_, err := tx.ExecContext(ctx, SQL, role.RoleName, role.RoleId)
	helper.PanicIfError(err)

	return role
}

func (repository *RoleManagementRepositoryImpl) DeleteRole(ctx context.Context, tx *sql.Tx, role domain.RoleManagement) {
	SQL := "delete from role_management where role_id = $1"
	_, err := tx.ExecContext(ctx, SQL, role.RoleId)
	helper.PanicIfError(err)
}

func (repository *RoleManagementRepositoryImpl) FindByRoleId(ctx context.Context, tx *sql.Tx, roleId int) (domain.RoleManagement, error) {
	SQL := "SELECT role_id, role_name FROM role_management WHERE role_id = $1"
	row := tx.QueryRowContext(ctx, SQL, roleId)

	role := domain.RoleManagement{}
	err := row.Scan(&role.RoleId, &role.RoleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return role, errors.New("role not found")
		}
		helper.PanicIfError(err)
	}
	return role, nil
}

func (repository *RoleManagementRepositoryImpl) FindByRoleName(ctx context.Context, tx *sql.Tx, roleName string) (domain.RoleManagement, error) {
	SQL := "SELECT role_id, role_name FROM role_management WHERE role_name = $1"
	row := tx.QueryRowContext(ctx, SQL, roleName)

	role := domain.RoleManagement{}
	err := row.Scan(&role.RoleId, &role.RoleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return role, errors.New("role not found")
		}
		helper.PanicIfError(err)
	}
	return role, nil
}

func (repository *RoleManagementRepositoryImpl) FindAllRole(ctx context.Context, tx *sql.Tx) []domain.RoleManagement {
	SQL := "select role_id, role_name from role_management"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var roles []domain.RoleManagement
	for rows.Next() {
		role := domain.RoleManagement{}
		err := rows.Scan(&role.RoleId, &role.RoleName)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}
	return roles
}
