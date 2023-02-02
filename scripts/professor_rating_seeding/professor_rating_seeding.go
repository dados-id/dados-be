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
	GOROUTINE := 5

	for i := 1; i <= GOROUTINE; i++ {
		wg.Add(1)
		go createProfessorRating(NDATA, *queries, &wg)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data ProfessorRating to database\n", NDATA*GOROUTINE)
}

func createProfessorRating(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= NDATA; i++ {
		ctx := context.Background()

		listRandomUserID, err := queries.ListRandomUserID(context.Background())
		exception.FatalIfNeeded(err, "Error Get Count User")

		randomProfessorID, err := queries.RandomProfessorID(context.Background())
		exception.FatalIfNeeded(err, "Error Get Count Professor")

		listCoursesByProfessorID, err := queries.ListCoursesByProfessorId(context.Background(), randomProfessorID)
		exception.FatalIfNeeded(err, "Error Get Random Course Code")

		randomCourseCode := util.RandomPickArrayStr(listCoursesByProfessorID)
		randomUserID := util.RandomPickArrayStr(listRandomUserID)

		professorRating := util.GetValidProfessorRating(randomUserID, randomProfessorID, randomCourseCode)

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

		listRandomTag, err := queries.ListRandomTag(ctx)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data on Create Tag\n %s", i, err.Error())
			continue
		}

		for _, randomTag := range listRandomTag {
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
