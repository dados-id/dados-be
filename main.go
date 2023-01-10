package main

import (
	"github.com/dados-id/dados-be/api"
	"github.com/dados-id/dados-be/config"
	"github.com/dados-id/dados-be/exception"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	configuration := config.LoadConfig(".")
	runGinServer(configuration)
}

func runGinServer(configuration config.Config) {
	server, err := api.NewServer(configuration)
	exception.FatalIfNeeded(err, "cannot create server")

	err = server.Start(configuration.HTTPServerAddress)
	exception.FatalIfNeeded(err, "cannot start server")
}
