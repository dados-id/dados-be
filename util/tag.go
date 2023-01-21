package util

import (
	db "github.com/dados-id/dados-be/db/sqlc"
)

func GetValidTag() (Tag db.Tag) {
	Tag = db.Tag{
		Name: randomName(),
	}
	return
}
