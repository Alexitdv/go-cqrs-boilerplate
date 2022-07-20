-- +goose Up
-- +goose StatementBegin
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS users
(
    uuid       CHAR(36)     NOT NULL PRIMARY KEY,
    password   CHAR(60)     NOT NULL,
    phone      CHAR(16)     NOT NULL,
    email      VARCHAR(100) NOT NULL,
    name       VARCHAR(100) NOT NULL DEFAULT '',
    lastName   VARCHAR(100) NOT NULL DEFAULT '',
    active     SMALLINT     NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    UNIQUE (phone),
    UNIQUE (email)
);

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
DROP TABLE users;