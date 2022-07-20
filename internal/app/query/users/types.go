package users

import "time"

type User struct {
	UUID     string    `db:"uuid"`
	Phone    string    `db:"phone"`
	Name     string    `db:"name"`
	LastName string    `db:"lastname"`
	Email    string    `db:"email"`
	Created  time.Time `db:"created_at"`
}
