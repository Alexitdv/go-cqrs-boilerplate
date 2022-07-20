package app

import "github.com/jmoiron/sqlx"

type Options struct {
	DB *sqlx.DB
}
