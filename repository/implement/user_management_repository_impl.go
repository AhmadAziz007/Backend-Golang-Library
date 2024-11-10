package implement

import (
	"context"
	"database/sql"
	"errors"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
)

type UserManagementRepositoryImpl struct{}

func (User *UserManagementRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, user domain.UserManagement) domain.UserManagement {
	SQL := "INSERT INTO user_management(role_id, user_name, email, password) VALUES ($1, $2, $3, $4) RETURNING user_id"
	row := tx.QueryRowContext(ctx, SQL, user.RoleId, user.UserName, user.Email, user.Password)
	err := row.Scan(&user.UserId)
	helper.PanicIfError(err)
	return user
}

func (User *UserManagementRepositoryImpl) UpdateUser(ctx context.Context, tx *sql.Tx, user domain.UserManagement) domain.UserManagement {
	SQL := "UPDATE user_management SET role_id = $1, user_name = $2, email = $3, password = $4 WHERE user_id = $5"
	_, err := tx.ExecContext(ctx, SQL,
		user.RoleId,
		user.UserName,
		user.Email,
		user.Password,
		user.UserId)
	helper.PanicIfError(err)
	return user
}

func (User *UserManagementRepositoryImpl) DeleteUser(ctx context.Context, tx *sql.Tx, user domain.UserManagement) {
	SQL := "DELETE FROM user_management WHERE user_id = $1"
	_, err := tx.ExecContext(ctx, SQL, user.UserId)
	helper.PanicIfError(err)
}

func (User *UserManagementRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.UserManagement, error) {
	SQL := `
        SELECT a.user_id, a.role_id, a.user_name, a.email, a.password
        FROM user_management a
        LEFT JOIN role_management c ON a.role_id = c.role_id
        WHERE a.email = $1
    `
	row := tx.QueryRowContext(ctx, SQL, email)
	user := domain.UserManagement{}
	err := row.Scan(
		&user.UserId,
		&user.RoleId,
		&user.UserName,
		&user.Email,
		&user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user is not found")
		}
		helper.PanicIfError(err)
	}
	return user, nil
}

func (User *UserManagementRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.UserManagement, error) {
	SQL := `
		SELECT a.user_id, a.role_id, a.user_name, a.email, a.password
		FROM user_management a
		LEFT JOIN role_management c ON a.role_id = c.role_id
		WHERE a.user_id = $1
	`
	row := tx.QueryRowContext(ctx, SQL, userId)
	user := domain.UserManagement{}
	err := row.Scan(
		&user.UserId,
		&user.RoleId,
		&user.UserName,
		&user.Email,
		&user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user is not found")
		}
		helper.PanicIfError(err)
	}
	return user, nil
}

func (User *UserManagementRepositoryImpl) FindAllUser(ctx context.Context, tx *sql.Tx) []domain.UserManagement {
	SQL := `
		SELECT a.user_id, a.role_id, a.user_name, a.email, a.password
		FROM user_management a
		LEFT JOIN role_management c ON a.role_id = c.role_id
		WHERE a.user_id = $1
	`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.UserManagement
	for rows.Next() {
		user := domain.UserManagement{}
		err := rows.Scan(
			&user.UserId,
			&user.RoleId,
			&user.UserName,
			&user.Email,
			&user.Password)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
