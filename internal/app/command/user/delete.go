package user

import (
	"boilerplate/internal/domain/user"
	"context"

	"github.com/google/uuid"
)

type DeleteUserHandler struct {
	userRepo user.Repository
}

func NewDeleteHandler(userRepo user.Repository) DeleteUserHandler {
	return DeleteUserHandler{
		userRepo: userRepo,
	}
}

func (h DeleteUserHandler) Handle(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = h.userRepo.DeleteUser(ctx, uid.String())
	if err != nil {
		return err
	}
	return nil
}
