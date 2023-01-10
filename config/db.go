package config

import (
	"database/sql"

	"github.com/dados-id/dados-be/exception"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewPostgres(dbDriver, dbSource string) *sql.DB {
	conn, err := sql.Open(dbDriver, dbSource)
	exception.FatalIfNeeded(err, "cannot connect to db")

	return conn
}
