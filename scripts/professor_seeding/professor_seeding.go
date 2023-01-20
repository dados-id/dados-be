package main

import (
	"context"
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
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go createProfessor(NDATA, *queries, &wg)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data Professor to database\n", NDATA)
}

func createProfessor(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup) {
	defer wg.Done()

	totalRowSchool, err := queries.CountSchool(context.Background())
	exception.FatalIfNeeded(err, "Error Count School")

	totalRowFaculty, err := queries.CountFaculty(context.Background())
	exception.FatalIfNeeded(err, "Error Count Faculty")

	for i := 1; i <= NDATA; i++ {
		professor := util.GetValidProfessor(totalRowSchool, totalRowFaculty)

		arg := sqlc.CreateProfessorParams{
			FirstName: professor.FirstName,
			LastName:  professor.LastName,
			FacultyID: professor.FacultyID,
			SchoolID:  professor.SchoolID,
		}

		_, err := queries.CreateProfessor(context.Background(), arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
