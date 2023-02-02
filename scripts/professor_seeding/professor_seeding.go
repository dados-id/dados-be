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

	NDATA := 100
	GOROUTINE := 5

	for i := 1; i <= GOROUTINE; i++ {
		wg.Add(1)
		go createProfessor(NDATA, *queries, &wg)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data Professor to database\n", NDATA*GOROUTINE)
}

func createProfessor(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= NDATA; i++ {
		ctx := context.Background()

		randomSchoolID, err := queries.RandomSchoolID(ctx)
		exception.FatalIfNeeded(err, "Error Count School")

		randomFacultyID, err := queries.RandomFacultyID(ctx)
		exception.FatalIfNeeded(err, "Error Count Faculty")

		professor := util.GetValidProfessor(randomSchoolID, randomFacultyID)

		arg := sqlc.CreateProfessorParams{
			FirstName: professor.FirstName,
			LastName:  professor.LastName,
			FacultyID: professor.FacultyID,
			SchoolID:  professor.SchoolID,
		}

		createdProfessorID, err := queries.CreateProfessor(ctx, arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		listRandomUserID, err := queries.ListRandomUserID(ctx)
		exception.FatalIfNeeded(err, "Error Count User")

		for _, randomUserID := range listRandomUserID {
			arg := sqlc.SaveProfessorParams{
				ProfessorID: createdProfessorID,
				UserID:      randomUserID,
			}

			err = queries.SaveProfessor(ctx, arg)
			if err != nil {
				fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
				continue
			}
		}

		listRandomCourseCode, err := queries.ListRandomCourseCode(ctx)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		for _, randomCourseCode := range listRandomCourseCode {
			arg := sqlc.CreateProfessorCourseAssociationParams{
				CourseCode:  randomCourseCode,
				ProfessorID: createdProfessorID,
			}

			err = queries.CreateProfessorCourseAssociation(ctx, arg)
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
