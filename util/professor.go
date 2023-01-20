package util

import (
	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidProfessor(totalRowSchool, totalRowFaculty int64) (professor db.Professor) {
	professor = db.Professor{
		FirstName: randomName(),
		LastName:  randomName(),
		FacultyID: RandomInt(1, totalRowFaculty),
		SchoolID:  RandomInt(1, totalRowSchool),
	}
	return
}
