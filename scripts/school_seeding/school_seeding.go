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

	NDATA := 1000
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
		ctx := context.Background()
		school := util.GetValidSchool()

		arg := sqlc.CreateSchoolParams{
			Name:     school.Name,
			NickName: school.NickName,
			City:     school.City,
			Province: school.Province,
			Website:  school.Website,
			Email:    school.Email,
		}

		schoolID, err := queries.CreateSchool(ctx, arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		listRandomFacultyID, err := queries.ListRandomFacultyID(ctx)
		exception.FatalIfNeeded(err, "Error Count User")

		for _, randomFacultyID := range listRandomFacultyID {
			arg := sqlc.CreateSchoolFacultyAssociationParams{
				FacultyID: randomFacultyID,
				SchoolID:  schoolID,
			}

			err = queries.CreateSchoolFacultyAssociation(ctx, arg)
			if err != nil {
				fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
				continue
			}
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
