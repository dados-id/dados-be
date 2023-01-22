package model

import db "github.com/dados-id/dados-be/db/sqlc"

type CreateProfessorRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FacultyID int64  `json:"facultyID"`
	SchoolID  int64  `json:"schoolID"`
}

type UpdateProfessorStatusRequest struct {
	Status string `json:"status"`
}

type GetProfessorRequest struct {
	ProfessorID int64 `uri:"professor_id" binding:"required,min=1"`
}

type GetProfessorInfoResponse struct {
	db.GetProfessorInfoAggregateRow
	Top5Tags []string `json:"top5Tags"`
	Courses  []string `json:"courses"`
}

type ListProfessorsQueryRequest struct {
	PageID   int32   `form:"page_id" binding:"required,min=1"`
	PageSize int32   `form:"page_size" binding:"required,min=5"`
	Name     *string `form:"name"`
}

func (x *ListProfessorsQueryRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type ListProfessorsQueryBySchoolRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

type ListProfessorsURIBySchoolRequest struct {
	SchoolID int64 `uri:"school_id" binding:"required,min=1"`
}

type ListProfessorsQueryByFacultyRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

type ListProfessorsURIByFacultyRequest struct {
	FacultyID int64 `uri:"faculty_id" binding:"required,min=1"`
}
