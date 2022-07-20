package app

import (
	"boilerplate/internal/app/query/users"

	"github.com/sirupsen/logrus"

	"boilerplate/internal/adapters/user"
	userCmd "boilerplate/internal/app/command/user"
)

type Application struct {
	Options           *Options
	Commands          Commands
	Queries           Queries
	ShutdownFunctions []func() error
}

// TODO: Provide jwt token to service meta
// TODO: Design ACL

type Commands struct {
	AddUser        userCmd.AddUserHandler
	UpdateUser     userCmd.UpdateUserHandler
	Login          userCmd.LoginHandler
	UpdatePassword userCmd.UpdatePasswordHandler
	DeleteUser     userCmd.DeleteUserHandler
}

type Queries struct {
	GetUser users.UserHandler
}

func NewApplication(options *Options) (*Application, error) {
	userRepo := user.NewPGRepository(options.DB)
	return &Application{
		Options: options,
		Commands: Commands{
			// Users
			AddUser:        userCmd.NewSignupHandler(userRepo),
			UpdateUser:     userCmd.NewUpdateUserHandler(userRepo),
			UpdatePassword: userCmd.NewUpdatePasswordHandler(userRepo),
			Login:          userCmd.NewLoginHandler(userRepo),
			DeleteUser:     userCmd.NewDeleteHandler(userRepo),
		},
		Queries: Queries{
			GetUser: users.NewUserHandler(userRepo),
		},
		ShutdownFunctions: []func() error{},
	}, nil
}

func (a *Application) Shutdown() {
	logrus.Info("Application shutdown start")
	for _, fn := range a.ShutdownFunctions {
		if err := fn(); err != nil {
			logrus.Error(err.Error())
		}
	}
	logrus.Info("Application shutdown done")
}
