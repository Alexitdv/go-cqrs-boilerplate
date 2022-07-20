package user

import (
	queryUsers "boilerplate/internal/app/query/users"
	"boilerplate/internal/domain/user/phone"
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/doug-martin/goqu/v9"

	"boilerplate/internal/domain/user"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql" // to register Mysql driver
	"github.com/jmoiron/sqlx"
)

type PGRepository struct {
	db      *sqlx.DB
	builder goqu.DialectWrapper
}

func NewPGRepository(db *sqlx.DB) PGRepository {
	return PGRepository{
		db:      db,
		builder: goqu.Dialect("postgres"),
	}
}

func (r PGRepository) GetUserQueryModel(_ context.Context, id string) (*queryUsers.User, error) {
	var qUser queryUsers.User
	err := r.db.Get(
		&qUser,
		"SELECT uuid, phone, email, name, lastName, created_at FROM users WHERE uuid = $1",
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrorNotFound
		}
		return nil, err
	}
	return &qUser, nil
}

func (r PGRepository) SaveUser(_ context.Context, u *user.User) error {
	query := `INSERT INTO users (uuid, password, email, phone, name, lastname) VALUES(:uuid, :password, :email, :phone, :name, :lastname)`
	if _, err := r.db.NamedExec(query, u); err != nil {
		return err
	}
	return nil
}

func (r PGRepository) GetUser(_ context.Context, id string) (*user.User, error) {
	userObj := user.NewEmptyUser()
	err := r.db.Get(
		userObj,
		"SELECT uuid, password, phone, email, name, lastname FROM users WHERE uuid = $1",
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrorNotFound
		}
		return nil, err
	}
	return userObj, nil
}

func (r PGRepository) GetUserByPhone(_ context.Context, phone phone.Phone) (*user.User, error) {
	userObj := user.NewEmptyUser()
	err := r.db.Get(
		userObj,
		"SELECT uuid, password, phone, email, name, lastname FROM users WHERE phone = $1",
		phone,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrorNotFound
		}
		return nil, err
	}
	return userObj, nil
}

func (r PGRepository) DeleteUser(_ context.Context, id string) error {
	query := `DELETE FROM users WHERE uuid = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	cols, _ := res.RowsAffected()
	if cols == 0 {
		return errors.New("nothing to delete")
	}
	return nil
}

func (r PGRepository) UpdateUser(_ context.Context, u *user.User) error {
	fields, err := u.ChangedTagMapped("db", *u)
	if err != nil {
		return err
	}
	query, _, _ := r.builder.Update("users").Set(fields).Where(goqu.C("uuid").Eq(u.UUID)).ToSQL()
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
