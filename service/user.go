package service

import (
	"context"
	"errhello/model"
)

// UserService represents a service for managing users.
type User interface {
	FindUserByID(ctx context.Context, id int) (*model.User, error)

	CreateUser(ctx context.Context, user *model.User) error
}
