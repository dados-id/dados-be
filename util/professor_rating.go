package util

import db "github.com/dados-id/dados-be/db/sqlc"

func GetValidProfessorRating(totalRowUser, totalRowProfessor int64, courseCode string) (ProfessorRating db.ProfessorRating) {
	ProfessorRating = db.ProfessorRating{
		Quality:             convertIntToStr(RandomFloat(0, 5)),
		Difficult:           convertIntToStr(RandomFloat(0, 5)),
		WouldTakeAgain:      int16(RandomInt(0, 1)),
		TakenForCredit:      int16(RandomInt(0, 2)),
		UseTextbooks:        int16(RandomInt(0, 2)),
		AttendanceMandatory: int16(RandomInt(0, 2)),
		Grade:               RandomString(5),
		CourseCode:          courseCode,
		UserID:              RandomInt(1, totalRowUser),
		ProfessorID:         RandomInt(1, totalRowProfessor),
		Review:              randomReview(),
	}
	return
}
