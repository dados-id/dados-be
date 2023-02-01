package util

import (
	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidProfessor(totalRowSchool, totalRowFaculty int32) (professor db.Professor) {
	professor = db.Professor{
		FirstName: randomName(),
		LastName:  randomName(),
		FacultyID: totalRowFaculty,
		SchoolID:  totalRowSchool,
	}
	return
}
