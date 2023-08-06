package model

type ProfessorRatingRequest struct {
	ProfessorID       int32 `uri:"professor_id" binding:"required,min=1"`
	ProfessorRatingID int32 `uri:"professor_rating_id" binding:"required,min=1"`
}

type ListProfessorRatingURIRequest struct {
	ProfessorID int32 `uri:"professor_id" binding:"required,min=1"`
}

type ListProfessorRatingQueryRequest struct {
	PageID     int32   `form:"page_id" binding:"required,min=1"`
	PageSize   int32   `form:"page_size" binding:"required,min=5,max=10"`
	CourseCode *string `form:"course_code"`
	Rating     *int16  `form:"rating"`
}

func (x *ListProfessorRatingQueryRequest) GetCourseCode() string {
	if x != nil && x.CourseCode != nil {
		return *x.CourseCode
	}
	return ""
}

func (x *ListProfessorRatingQueryRequest) GetRating() int16 {
	if x != nil && x.Rating != nil {
		return *x.Rating
	}
	return 0
}

type CreateProfessorRatingURIRequest struct {
	ProfessorID int32 `uri:"professor_id" binding:"required,min=1"`
}

type CreateProfessorRatingJSONRequest struct {
	Quality             string   `json:"quality"`
	Difficult           string   `json:"difficult"`
	WouldTakeAgain      int16    `json:"wouldTakeAgain"`
	TakenForCredit      int16    `json:"takenForCredit"`
	UseTextbooks        int16    `json:"useTextbooks"`
	AttendanceMandatory int16    `json:"attendanceMandatory"`
	Grade               string   `json:"grade"`
	Tags                []string `json:"tags"`
	Review              string   `json:"review"`
	ProfessorID         int32    `json:"professorId"`
	CourseCode          string   `json:"courseCode"`
}

type UpdateProfessorRatingJSONRequest struct {
	Quality             *string   `json:"quality"`
	Difficult           *string   `json:"difficult"`
	WouldTakeAgain      *int16    `json:"wouldTakeAgain"`
	TakenForCredit      *int16    `json:"takenForCredit"`
	UseTextbooks        *int16    `json:"useTextbooks"`
	AttendanceMandatory *int16    `json:"attendanceMandatory"`
	Grade               *string   `json:"grade"`
	Tags                *[]string `json:"tags"`
	Review              *string   `json:"review"`
	UpVote              *int32    `json:"upVote"`
	DownVote            *int32    `json:"downVote"`
	CourseCode          *string   `json:"courseCode"`
}

func (x *UpdateProfessorRatingJSONRequest) GetQuality() string {
	if x != nil && x.Quality != nil {
		return *x.Quality
	}
	return ""
}

func (x *UpdateProfessorRatingJSONRequest) GetDifficult() string {
	if x != nil && x.Difficult != nil {
		return *x.Difficult
	}
	return ""
}

func (x *UpdateProfessorRatingJSONRequest) GetWouldTakeAgain() int16 {
	if x != nil && x.WouldTakeAgain != nil {
		return *x.WouldTakeAgain
	}
	return 0
}

func (x *UpdateProfessorRatingJSONRequest) GetTakenForCredit() int16 {
	if x != nil && x.TakenForCredit != nil {
		return *x.TakenForCredit
	}
	return 0
}

func (x *UpdateProfessorRatingJSONRequest) GetUseTextbooks() int16 {
	if x != nil && x.UseTextbooks != nil {
		return *x.UseTextbooks
	}
	return 0
}

func (x *UpdateProfessorRatingJSONRequest) GetAttendanceMandatory() int16 {
	if x != nil && x.AttendanceMandatory != nil {
		return *x.AttendanceMandatory
	}
	return 0
}

func (x *UpdateProfessorRatingJSONRequest) GetGrade() string {
	if x != nil && x.Grade != nil {
		return *x.Grade
	}
	return ""
}

func (x *UpdateProfessorRatingJSONRequest) GetTags() []string {
	if x != nil && x.Tags != nil {
		return *x.Tags
	}
	return nil
}

func (x *UpdateProfessorRatingJSONRequest) GetReview() string {
	if x != nil && x.Review != nil {
		return *x.Review
	}
	return ""
}

func (x *UpdateProfessorRatingJSONRequest) GetUpVote() int32 {
	if x != nil && x.UpVote != nil {
		return *x.UpVote
	}
	return 0
}

func (x *UpdateProfessorRatingJSONRequest) GetDownVote() int32 {
	if x != nil && x.DownVote != nil {
		return *x.DownVote
	}
	return 0
}

func (x *UpdateProfessorRatingJSONRequest) GetCourseCode() string {
	if x != nil && x.CourseCode != nil {
		return *x.CourseCode
	}
	return ""
}
