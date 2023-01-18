package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/dados-id/dados-be/config"
	sqlc "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/util"
	_ "github.com/lib/pq"
)

func main() {
	configuration := config.LoadConfig(".")
	database := config.NewPostgres(configuration.DBDriver, configuration.DBSource)
	queries := sqlc.New(database)

	var wg sync.WaitGroup

	NDATA := 500
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go createUser(NDATA, *queries, &wg)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data User to database\n", NDATA)
}

func createUser(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= NDATA; i++ {
		user := util.GetValidUser()
		arg := sqlc.CreateUserParams{
			FirstName:                user.FirstName,
			LastName:                 user.LastName,
			School:                   user.School,
			ExpectedYearOfGraduation: user.ExpectedYearOfGraduation,
			Email:                    user.Email,
		}
		_, err := queries.CreateUser(context.Background(), arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
