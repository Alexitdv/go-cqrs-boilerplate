package user

import (
	"boilerplate/internal/domain/user"
	"boilerplate/internal/domain/user/phone"
	"context"

	"github.com/pkg/errors"
)

type UpdatePasswordHandler struct {
	userRepo user.Repository
}

func NewUpdatePasswordHandler(userRepo user.Repository) UpdatePasswordHandler {
	return UpdatePasswordHandler{
		userRepo: userRepo,
	}
}

func (h UpdatePasswordHandler) Handle(ctx context.Context, phoneStr, oldPassword, newPassword string) error {
	ph, err := phone.NewPhone(phoneStr)
	if err != nil {
		return err
	}
	u, err := h.userRepo.GetUserByPhone(ctx, ph)
	if err != nil {
		return err
	}
	if !u.ComparePassword(oldPassword) {
		return errors.New("current password is incorrect")
	}
	err = user.WithPassword(newPassword)(u)
	if err != nil {
		return err
	}
	err = h.userRepo.UpdateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
