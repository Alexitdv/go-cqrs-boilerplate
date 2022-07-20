package users

import "context"

type UserHandler struct {
	readModel UserModel
}

func NewUserHandler(readModel UserModel) UserHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return UserHandler{readModel: readModel}
}

type UserModel interface {
	GetUserQueryModel(ctx context.Context, id string) (*User, error)
}

func (h UserHandler) Handle(ctx context.Context, id string) (*User, error) {
	return h.readModel.GetUserQueryModel(ctx, id)
}
