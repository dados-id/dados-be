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

	NDATA := 5000
	GOROUTINE := 5

	for i := 1; i <= GOROUTINE; i++ {
		wg.Add(1)
		go createSchool(NDATA, *queries, &wg)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data School to database\n", NDATA*GOROUTINE)
}

func createSchool(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= NDATA; i++ {
		school := util.GetValidSchool()

		arg := sqlc.CreateSchoolParams{
			Name:     school.Name,
			NickName: school.NickName,
			City:     school.City,
			Province: school.Province,
			Website:  school.Website,
			Email:    school.Email,
		}

		_, err := queries.CreateSchool(context.Background(), arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
