package model

type LoginUserRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}

type CreateUserRequest struct {
	FirstName                string `json:"firstName" binding:"required"`
	LastName                 string `json:"lastName" binding:"required"`
	School                   string `json:"school" binding:"required"`
	ExpectedYearOfGraduation int16  `json:"expectedYearOfGraduation" binding:"required"`
	Email                    string `json:"email" binding:"required,email"`
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type UpdateUserJSONRequest struct {
	FirstName                *string `json:"firstName"`
	LastName                 *string `json:"lastName"`
	School                   *string `json:"school"`
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

func (x *UpdateUserJSONRequest) GetSchool() string {
	if x != nil && x.School != nil {
		return *x.School
	}
	return ""
}

func (x *UpdateUserJSONRequest) GetExpectedYearOfGraduation() int16 {
	if x != nil && x.ExpectedYearOfGraduation != nil {
		return *x.ExpectedYearOfGraduation
	}
	return 0
}

type UpdateUserURIRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type UserListURIRequest struct {
	UserID int64 `uri:"id" binding:"required,min=1"`
}

type UserListQueryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

type UnsaveProfessorURIRequest struct {
	UserID      int64 `uri:"user_id" binding:"required,min=1"`
	ProfessorID int64 `uri:"professor_id" binding:"required,min=1"`
}

type SaveProfessorURIRequest struct {
	UserID      int64 `uri:"user_id" binding:"required,min=1"`
	ProfessorID int64 `uri:"professor_id" binding:"required,min=1"`
}
