package main

import (
	"context"
	"fmt"

	"github.com/dados-id/dados-be/config"
	sqlc "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/util"
	_ "github.com/lib/pq"
)

func main() {
	configuration := config.LoadConfig(".")
	database := config.NewPostgres(configuration.DBDriver, configuration.DBSource)
	queries := sqlc.New(database)

	NDATA := 3
	createTag(NDATA, *queries)

	fmt.Printf("Successfully added %d data Tag to database\n", NDATA)
}

func createTag(NDATA int, queries sqlc.Queries) {

	for i := 1; i <= NDATA; i++ {
		tag := util.GetValidTag()

		_, err := queries.CreateTag(context.Background(), tag.Name)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
