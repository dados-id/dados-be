package util

import (
	"database/sql"

	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidUser(randomSchoolID int64) (user db.User) {
	user = db.User{
		ID:                       RandomString(28),
		FirstName:                randomName(),
		LastName:                 randomName(),
		ExpectedYearOfGraduation: sql.NullInt16{Int16: int16(randomExpectedYearOfGraduation()), Valid: true},
		Email:                    randomEmail(),
		SchoolID:                 sql.NullInt64{Int64: randomSchoolID, Valid: true},
	}
	return
}

// randomExpectedYearOfGraduation generates a random ExpectedYearOfGraduation of transaction
func randomExpectedYearOfGraduation() int64 {
	return RandomInt(2023, 3000)
}
