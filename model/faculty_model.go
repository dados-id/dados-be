package model

type ListFacultyURIRequest struct {
	SchoolID int32 `uri:"school_id" binding:"required,min=1"`
}

type CreateFacultyRequest struct {
	Name     string `json:"name"`
	SchoolID int32  `json:"schoolID"`
}
