package user

import (
	"boilerplate/internal/domain/user"
	"context"
)

type AddUserHandler struct {
	userRepo user.Repository
}

func NewSignupHandler(userRepo user.Repository) AddUserHandler {
	return AddUserHandler{
		userRepo: userRepo,
	}
}

func (h AddUserHandler) Handle(ctx context.Context, user *user.User) (string, error) {
	err := h.userRepo.SaveUser(ctx, user)
	if err != nil {
		return "", err
	}
	return user.UUID.String(), nil
}
