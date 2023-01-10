package main

import (
	"github.com/dados-id/dados-be/api"
	"github.com/dados-id/dados-be/config"
)

func main() {
	configuration := config.LoadConfig(".")
	api.RunGinServer(configuration)
}
