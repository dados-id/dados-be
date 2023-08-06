package model

type ListCoursesURIRequest struct {
	ProfessorID int32 `uri:"professor_id" binding:"required,min=1"`
}
