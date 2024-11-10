package repository

import (
	"context"
	"database/sql"
	"library-synapsis/model/domain"
)

type UserManagementRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, user domain.UserManagement) domain.UserManagement
	UpdateUser(ctx context.Context, tx *sql.Tx, user domain.UserManagement) domain.UserManagement
	DeleteUser(ctx context.Context, tx *sql.Tx, user domain.UserManagement)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.UserManagement, error)
	FindAllUser(ctx context.Context, tx *sql.Tx) []domain.UserManagement
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.UserManagement, error)
}
