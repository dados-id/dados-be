package util

import db "github.com/dados-id/dados-be/db/sqlc"

func GetValidProfessorRating(randomUserID string, totalRowProfessor int64, courseCode string) (ProfessorRating db.ProfessorRating) {
	ProfessorRating = db.ProfessorRating{
		Quality:             convertIntToStr(RandomFloat(1, 5)),
		Difficult:           convertIntToStr(RandomFloat(1, 5)),
		WouldTakeAgain:      int16(RandomInt(0, 1)),
		TakenForCredit:      int16(RandomInt(0, 2)),
		UseTextbooks:        int16(RandomInt(0, 2)),
		AttendanceMandatory: int16(RandomInt(0, 2)),
		Grade:               RandomString(5),
		CourseCode:          courseCode,
		UserID:              randomUserID,
		ProfessorID:         RandomInt(1, totalRowProfessor),
		Review:              randomReview(),
	}
	return
}
