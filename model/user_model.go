package model

type LoginUserRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}

type CreateUserRequest struct {
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	ExpectedYearOfGraduation int16  `json:"expectedYearOfGraduation"`
	Email                    string `json:"email"`
	SchoolID                 int64  `json:"schoolId"`
}

type GetUserRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

type UpdateUserJSONRequest struct {
	FirstName                *string `json:"firstName"`
	LastName                 *string `json:"lastName"`
	SchoolID                 *int64  `json:"schoolId"`
	ExpectedYearOfGraduation *int16  `json:"expectedYearOfGraduation"`
}

func (x *UpdateUserJSONRequest) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *UpdateUserJSONRequest) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *UpdateUserJSONRequest) GetSchoolID() int64 {
	if x != nil && x.SchoolID != nil {
		return *x.SchoolID
	}
	return 0
}

func (x *UpdateUserJSONRequest) GetExpectedYearOfGraduation() int16 {
	if x != nil && x.ExpectedYearOfGraduation != nil {
		return *x.ExpectedYearOfGraduation
	}
	return 0
}

type UserListQueryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type UnsaveProfessorURIRequest struct {
	ProfessorID int64 `uri:"professor_id" binding:"required,min=1"`
}

type SaveProfessorURIRequest struct {
	ProfessorID int64 `uri:"professor_id" binding:"required,min=1"`
}
