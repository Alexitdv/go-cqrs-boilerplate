package user

import (
	"boilerplate/internal/domain/user/phone"
	"context"
	"errors"
)

type Repository interface {
	SaveUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*User, error)
	GetUserByPhone(ctx context.Context, phone phone.Phone) (*User, error)
}

var (
	ErrorNotFound = errors.New("user wasn't found")
)
