package repository

import (
	"context"
	"database/sql"
	"errors"
	"todo-app-be/helper"
	"todo-app-be/model/domain"

	"github.com/sirupsen/logrus"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(email, full_name ,password) values (?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, user.Email, user.FullName, user.Password)
	helper.PanicIfError(err)

	logrus.Debugln("error:", result)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

func (repository *UserRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select * from users where id = ?"

	logrus.WithField("SQL:", SQL).Infoln("get user by id")

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.FullName, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	}

	return user, errors.New("user is not found")
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set email = ?, full_name = ?, password = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.FullName, user.Password, user.Id)
	logrus.Errorln(err)
	helper.PanicIfError(err)

	return user
}
