package main

import (
	"github.com/dados-id/dados-be/api"
	"github.com/dados-id/dados-be/config"
	db "github.com/dados-id/dados-be/db/sqlc"
)

func main() {
	configuration := config.LoadConfig(".")
	database := config.NewPostgres(configuration.DBDriver, configuration.DBSource)
	firebaseClient := config.NewFireBase()

	config.RunDBMigration(configuration.MigrationURL, configuration.DBSource)

	query := db.New(database)

	api.RunGinServer(configuration, query, *firebaseClient)
}
