package repository

import (
	"context"
	"database/sql"
	"todo-app-be/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	GetById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
