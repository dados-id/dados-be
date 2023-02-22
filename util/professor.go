package util

import (
	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidProfessor(totalRowSchool, totalRowFaculty int32) (professor db.Professor) {
	professor = db.Professor{
		ID:        int32(RandomInt(1, 1000)),
		FirstName: randomName(),
		LastName:  randomName(),
		FacultyID: totalRowFaculty,
		SchoolID:  totalRowSchool,
	}
	return
}
