package implement

import (
	"context"
	"database/sql"
	"errors"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
)

type LoginRepositoryImpl struct{}

func NewLoginRepository() repository.LoginRepository {
	return &LoginRepositoryImpl{}
}

func (repository *LoginRepositoryImpl) FindByRoleName(ctx context.Context, tx *sql.Tx, roleId int) (domain.RoleManagement, error) {
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

func (repository *LoginRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, password string, email string) (domain.UserManagement, error) {
	SQL := `
        SELECT a.user_id, a.role_id, a.user_name, a.email, a.password
        FROM user_management a
        LEFT JOIN role_management c ON a.role_id = c.role_id
        WHERE a.email = $1
        AND a.password = $2
        ORDER BY a.user_id ASC
    `
	row := tx.QueryRowContext(ctx, SQL, email, password)
	user := domain.UserManagement{}
	err := row.Scan(
		&user.UserId,
		&user.RoleId,
		&user.UserName,
		&user.Email,
		&user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("email is not found")
		}
		helper.PanicIfError(err)
	}
	return user, nil
}
