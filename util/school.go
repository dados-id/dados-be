package util

import (
	"fmt"

	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidSchool() (school db.School) {
	school = db.School{
		Name:     randomName(),
		NickName: []string{randomName(), randomName(), randomName()},
		City:     randomName(),
		Province: randomName(),
		Website:  randomWebsite(),
		Email:    randomEmail(),
	}
	return
}

// randomWebsite generates a random website of school
func randomWebsite() string {
	return fmt.Sprintf("www.%s.id", RandomString(20))
}
