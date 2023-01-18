package util

import (
	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidSchoolRating(totalRowUser, totalRowSchool int64) (SchoolRating db.SchoolRating, password string) {
	SchoolRating = db.SchoolRating{
		UserID:        RandomInt(1, totalRowUser),
		SchoolID:      RandomInt(1, totalRowSchool),
		Reputation:    int16(RandomInt(1, 5)),
		Location:      int16(RandomInt(1, 5)),
		Opportunities: int16(RandomInt(1, 5)),
		Facilities:    int16(RandomInt(1, 5)),
		Internet:      int16(RandomInt(1, 5)),
		Food:          int16(RandomInt(1, 5)),
		Clubs:         int16(RandomInt(1, 5)),
		Social:        int16(RandomInt(1, 5)),
		Happiness:     int16(RandomInt(1, 5)),
		Safety:        int16(RandomInt(1, 5)),
		Review:        randomReview(),
	}
	return
}

// randomReview generates a random review of school rating
func randomReview() string {
	return RandomString(100)
}