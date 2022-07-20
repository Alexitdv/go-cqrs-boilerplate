package user

import (
	"context"

	"boilerplate/internal/domain/user"
)

type UpdateUserHandler struct {
	userRepo user.Repository
}

func NewUpdateUserHandler(userRepo user.Repository) UpdateUserHandler {
	return UpdateUserHandler{
		userRepo: userRepo,
	}
}

func (h UpdateUserHandler) Handle(ctx context.Context, user *user.User) error {
	if err := h.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
