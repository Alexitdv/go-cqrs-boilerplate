package user

import (
	"boilerplate/internal/domain/user/phone"
	"boilerplate/pkg/struct/changed"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	changed.Changed
	UUID     uuid.UUID   `db:"uuid"`
	Password string      `db:"password"`
	Email    string      `db:"email"`
	Phone    phone.Phone `db:"phone"`
	Name     string      `db:"name"`
	LastName string      `db:"lastname"`
}

func NewUser(opts ...Option) (*User, error) {
	user := &User{UUID: uuid.New(), Changed: map[string]interface{}{}}
	for _, opt := range opts {
		err := opt(user)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func NewEmptyUser() *User {
	return &User{Changed: map[string]interface{}{}}
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
