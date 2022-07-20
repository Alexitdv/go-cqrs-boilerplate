package user

import (
	"boilerplate/internal/domain/user/phone"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrOptionsShortPassword  = errors.New("password required min 6 symbols")
	ErrOptionsHashPassword   = errors.New("can't hash password")
	ErrOptionsIncorrectEmail = errors.New("email is incorrect")
)

type Option func(*User) error
type Options []Option

func NewOptions(opts ...Option) Options {
	return opts
}

func (o *Options) Append(opt Option) {
	*o = append(*o, opt)
}

func WithID(id string) Option {
	return func(u *User) error {
		uid, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		u.UUID = uid
		u.SetChanged("UUID", uid)
		return nil
	}
}

func WithName(name string) Option {
	return func(u *User) error {
		u.Name = name
		u.SetChanged("Name", name)
		return nil
	}
}

func WithLastName(name string) Option {
	return func(u *User) error {
		u.LastName = name
		u.SetChanged("LastName", name)
		return nil
	}
}

func WithPassword(password string) Option {
	return func(u *User) error {
		if len(password) < 6 {
			return ErrOptionsShortPassword
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("%s: %w", err.Error(), ErrOptionsHashPassword)
		}
		u.Password = string(hashedPassword)
		u.SetChanged("Password", string(hashedPassword))
		return nil
	}
}

func WithPhone(str string) Option {
	return func(u *User) error {
		ph, err := phone.NewPhone(str)
		if err != nil {
			return err
		}
		u.Phone = ph
		u.SetChanged("Phone", ph)
		return nil
	}
}

func WithEmail(email string) Option {
	return func(u *User) error {
		address, err := mail.ParseAddress(email)
		if err != nil {
			return fmt.Errorf("%s: %w", err.Error(), ErrOptionsIncorrectEmail)
		}
		u.Email = address.Address
		u.SetChanged("Email", address.Address)
		return nil
	}
}
