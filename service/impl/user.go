package impl

import (
	"context"
	"database/sql"
	"errhello"
	"errhello/model"
	"errhello/service"
)

type userpg struct {
	db *sql.DB
}

func (u *userpg) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	var query = `select id, username from users where id=$1`

	var user model.User
	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		return nil, &errhello.Error{Code: errhello.ENOTFOUND}
	}
	return &user, nil

}

func (u *userpg) CreateUser(ctx context.Context, user *model.User) error {
	const op = "service.impl.CreateUser"

	err := u.validate(user)
	if err != nil {
		return &errhello.Error{Op: op, Code: errhello.EINVALID, Message: "Username is required", Err: err}
	}

	return nil
}

func (u *userpg) validate(user *model.User) error {
	const op = "service.impl.validate"
	if user.Name == "" {
		return &errhello.Error{Op: op, Code: errhello.EINVALID, Message: "Username is required"}
	}

	return nil
}

func New(db *sql.DB) service.User {
	return &userpg{db: db}
}
