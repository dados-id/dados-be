package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/dados-id/dados-be/config"
	sqlc "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/util"
	_ "github.com/lib/pq"
)

func main() {
	configuration := config.LoadConfig(".")
	database := config.NewPostgres(configuration.DBDriver, configuration.DBSource)
	queries := sqlc.New(database)

	var wg sync.WaitGroup

	NDATA := 500
	GOROUTINE := 5

	for i := 1; i <= GOROUTINE; i++ {
		wg.Add(1)
		go createUser(NDATA, *queries, &wg)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data User to database\n", NDATA*GOROUTINE)
}

func createUser(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= NDATA; i++ {
		randomSchoolID, err := queries.RandomSchoolID(context.Background())
		exception.FatalIfNeeded(err, "Error Count School")

		user := util.GetValidUser(randomSchoolID)

		arg := sqlc.CreateUserParams{
			ID:                       user.ID,
			FirstName:                user.FirstName,
			LastName:                 user.LastName,
			ExpectedYearOfGraduation: user.ExpectedYearOfGraduation,
			Email:                    user.Email,
			SchoolID:                 sql.NullInt32{Int32: randomSchoolID, Valid: true},
		}

		_, err = queries.CreateUser(context.Background(), arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
