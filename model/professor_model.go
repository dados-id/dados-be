package model

import db "github.com/dados-id/dados-be/db/sqlc"

type CreateProfessorRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FacultyID int32  `json:"facultyID"`
	SchoolID  int32  `json:"schoolID"`
}

type UpdateProfessorStatusRequest struct {
	Status string `json:"status"`
}

type GetProfessorRequest struct {
	ProfessorID int32 `uri:"professor_id" binding:"required,min=1"`
}

type GetProfessorInfoResponse struct {
	db.GetProfessorInfoRow
	TopTags          []string `json:"topTags"`
	TopCoursesTaught []string `json:"topCoursesTaught"`
	Courses          []string `json:"courses"`
}

type ListProfessorsQueryRequest struct {
	PageID    int32   `form:"page_id" binding:"required,min=1"`
	PageSize  int32   `form:"page_size" binding:"required,min=5,max=10"`
	Name      *string `form:"name"`
	SortBy    *string `form:"sort_by"`
	SortOrder *string `form:"sort_order"`
}

func (x *ListProfessorsQueryRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ListProfessorsQueryRequest) GetSortBy() string {
	if x != nil && x.SortBy != nil {
		return *x.SortBy
	}
	return ""
}

func (x *ListProfessorsQueryRequest) GetSortOrder() string {
	if x != nil && x.SortOrder != nil {
		return *x.SortOrder
	}
	return ""
}

type ListProfessorsQueryBySchoolRequest struct {
	PageID    int32   `form:"page_id" binding:"required,min=1"`
	PageSize  int32   `form:"page_size" binding:"required,min=5,max=10"`
	SortBy    *string `form:"sort_by"`
	SortOrder *string `form:"sort_order"`
}

func (x *ListProfessorsQueryBySchoolRequest) GetSortBy() string {
	if x != nil && x.SortBy != nil {
		return *x.SortBy
	}
	return ""
}

func (x *ListProfessorsQueryBySchoolRequest) GetSortOrder() string {
	if x != nil && x.SortOrder != nil {
		return *x.SortOrder
	}
	return ""
}

type ListProfessorsURIBySchoolRequest struct {
	SchoolID int32 `uri:"school_id" binding:"required,min=1"`
}

type ListProfessorsQueryByFacultyRequest struct {
	PageID    int32   `form:"page_id" binding:"required,min=1"`
	PageSize  int32   `form:"page_size" binding:"required,min=5,max=10"`
	SortBy    *string `form:"sort_by"`
	SortOrder *string `form:"sort_order"`
}

func (x *ListProfessorsQueryByFacultyRequest) GetSortBy() string {
	if x != nil && x.SortBy != nil {
		return *x.SortBy
	}
	return ""
}

func (x *ListProfessorsQueryByFacultyRequest) GetSortOrder() string {
	if x != nil && x.SortOrder != nil {
		return *x.SortOrder
	}
	return ""
}

type ListProfessorsURIByFacultyRequest struct {
	FacultyID int32 `uri:"faculty_id" binding:"required,min=1"`
}
