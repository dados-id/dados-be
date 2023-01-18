package util

import (
	"fmt"

	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidUser() (user db.User) {
	user = db.User{
		FirstName:                randomName(),
		LastName:                 randomName(),
		School:                   randomSchool(),
		ExpectedYearOfGraduation: int16(randomExpectedYearOfGraduation()),
		Email:                    randomEmail(),
	}
	return
}

// randomSchool generates a random school of user
func randomSchool() string {
	return fmt.Sprintf("univesity of %s", RandomString(20))
}

// randomExpectedYearOfGraduation generates a random ExpectedYearOfGraduation of transaction
func randomExpectedYearOfGraduation() int64 {
	return RandomInt(2023, 3000)
}
