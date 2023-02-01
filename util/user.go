package util

import (
	"database/sql"

	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidUser(randomSchoolID int32) (user db.User) {
	user = db.User{
		ID:                       RandomString(28),
		FirstName:                randomName(),
		LastName:                 randomName(),
		ExpectedYearOfGraduation: sql.NullInt16{Int16: int16(randomExpectedYearOfGraduation()), Valid: true},
		Email:                    randomEmail(),
		SchoolID:                 sql.NullInt32{Int32: randomSchoolID, Valid: true},
	}
	return
}

// randomExpectedYearOfGraduation generates a random ExpectedYearOfGraduation of transaction
func randomExpectedYearOfGraduation() int {
	return RandomInt(2023, 3000)
}
