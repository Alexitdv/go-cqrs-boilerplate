package user

import (
	"boilerplate/internal/domain/user"
	"boilerplate/internal/domain/user/phone"
	"context"

	"github.com/davecgh/go-spew/spew"

	"github.com/pkg/errors"
)

var (
	ErrIncorrectPassword = errors.New("login or password is incorrect")
)

type LoginHandler struct {
	userRepo user.Repository
}

func NewLoginHandler(userRepo user.Repository) LoginHandler {
	return LoginHandler{
		userRepo: userRepo,
	}
}

func (h LoginHandler) Handle(ctx context.Context, phoneStr, password string) (result bool, err error) {
	userPhone, err := phone.NewPhone(phoneStr)
	if err != nil {
		return false, err
	}
	userObj, err := h.userRepo.GetUserByPhone(ctx, userPhone)
	spew.Dump(userObj)
	if err != nil {
		return
	}
	if !userObj.ComparePassword(password) {
		err = ErrIncorrectPassword
		return
	}
	return true, nil
}
