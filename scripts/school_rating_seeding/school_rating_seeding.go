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
		go createSchoolRating(NDATA, *queries, &wg)
	}
	wg.Wait()

	fmt.Printf("Successfully added %d data SchoolRating to database\n", NDATA)
}

func createSchoolRating(NDATA int, queries sqlc.Queries, wg *sync.WaitGroup) {
	defer wg.Done()

	totalRowUser, err := queries.CountUser(context.Background())
	exception.FatalIfNeeded(err, "Error Count User")

	totalRowSchool, err := queries.CountSchool(context.Background())
	exception.FatalIfNeeded(err, "Error Count User")

	for i := 1; i <= NDATA; i++ {
		schoolRating := util.GetValidSchoolRating(totalRowUser, totalRowSchool)
		arg := sqlc.CreateSchoolRatingParams{
			UserID:        schoolRating.UserID,
			SchoolID:      schoolRating.SchoolID,
			Reputation:    schoolRating.Reputation,
			Location:      schoolRating.Location,
			Opportunities: schoolRating.Opportunities,
			Facilities:    schoolRating.Facilities,
			Internet:      schoolRating.Internet,
			Food:          schoolRating.Food,
			Clubs:         schoolRating.Clubs,
			Social:        schoolRating.Social,
			Happiness:     schoolRating.Happiness,
			Safety:        schoolRating.Safety,
			Review:        schoolRating.Review,
		}

		_, err := queries.CreateSchoolRating(context.Background(), arg)
		if err != nil {
			fmt.Printf("Error seeded on the %dth data\n %s", i, err.Error())
			continue
		}

		if i%100 == 0 {
			fmt.Printf("Seeded %d data\n", i)
		}
	}
}
