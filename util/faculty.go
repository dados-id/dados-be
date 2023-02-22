package util

import (
	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidFaculty() (faculty db.Faculty) {
	faculty = db.Faculty{
		Name: randomName(),
	}
	return
}
