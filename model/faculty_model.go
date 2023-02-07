package model

type ListFacultyURIRequest struct {
	SchoolID int32 `uri:"school_id" binding:"required,min=1"`
}
