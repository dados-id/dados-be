package util

import (
	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidCourse() (course db.Course) {
	course = db.Course{
		Code: RandomString(30),
		Name: randomName(),
	}
	return
}
