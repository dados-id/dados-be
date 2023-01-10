package config

import (
	"database/sql"

	"github.com/dados-id/dados-be/exception"
)

func NewPostgres(dbDriver, dbSource string) *sql.DB {
	conn, err := sql.Open(dbDriver, dbSource)
	exception.FatalIfNeeded(err, "cannot connect to db")

	return conn
}
