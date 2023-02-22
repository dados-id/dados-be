package config

import (
	"database/sql"

	"github.com/dados-id/dados-be/exception"
	"github.com/rs/zerolog/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewPostgres(dbDriver, dbSource string) *sql.DB {
	conn, err := sql.Open(dbDriver, dbSource)
	exception.FatalIfNeeded(err, "cannot connect to db")

	return conn
}

func RunDBMigration(migrationURL, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	exception.FatalIfNeeded(err, "cannot create new migrate instance")

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
