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

	totalRowUser, err := queries.CountUser(context.Background())
	exception.FatalIfNeeded(err, "Error Get Count User")

	totalRowProfessor, err := queries.CountProfessor(context.Background())
	exception.FatalIfNeeded(err, "Error Get Count Professor")

	randomCourseCode, err := queries.RandomCourseCode(context.Background())
	exception.FatalIfNeeded(err, "Error Get Random Course Code")

	var wg sync.WaitGroup

	NDATA := 500
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go createProfessorRating(NDATA, *queries, &wg, totalRowUser, totalRowProfessor, randomCourseCode)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data ProfessorRating to database\n", NDATA)
}

func createProfessorRating(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup, totalRowUser, totalRowProfessor int64, courseCode string) {
	defer wg.Done()

	for i := 1; i <= NDATA; i++ {
		ctx := context.Background()
		professorRating := util.GetValidProfessorRating(totalRowUser, totalRowProfessor, courseCode)

		arg := sqlc.CreateProfessorRatingParams{
			Quality:             professorRating.Quality,
			Difficult:           professorRating.Difficult,
			WouldTakeAgain:      professorRating.WouldTakeAgain,
			TakenForCredit:      professorRating.TakenForCredit,
			UseTextbooks:        professorRating.UseTextbooks,
			AttendanceMandatory: professorRating.AttendanceMandatory,
			Grade:               professorRating.Grade,
			Review:              professorRating.Review,
			ProfessorID:         professorRating.ProfessorID,
			CourseCode:          professorRating.CourseCode,
			UserID:              professorRating.UserID,
		}

		createdProfessorRating, err := queries.CreateProfessorRating(ctx, arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		for j := 1; j <= 3; j++ {
			randomTag, err := queries.RandomTag(ctx)
			if err != nil {
				fmt.Printf("Error seeded on the %dth data on Create Tag\n %s", i, err.Error())
				continue
			}

			arg := sqlc.CreateProfessorRatingTagsParams{
				TagName:           randomTag,
				ProfessorRatingID: createdProfessorRating.ID,
			}

			err = queries.CreateProfessorRatingTags(ctx, arg)
			if err != nil {
				fmt.Printf("Error seeded on the %dth data on ProfessorTag\n %s", i, err.Error())
				continue
			}
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
