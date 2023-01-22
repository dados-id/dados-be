package model

type CreateSchoolRequest struct {
	Name     string   `json:"name"`
	NickName []string `json:"nickName"`
	City     string   `json:"city"`
	Province string   `json:"province"`
	Website  string   `json:"website"`
	Email    string   `json:"email"`
}

type GetSchoolRequest struct {
	SchoolID int64 `uri:"school_id" binding:"required,min=1"`
}

type ListSchoolsQueryRequest struct {
	PageID   int32   `form:"page_id" binding:"required,min=1"`
	PageSize int32   `form:"page_size" binding:"required,min=5"`
	Name     *string `form:"name"`
}

func (x *ListSchoolsQueryRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type UpdateSchoolStatusRequest struct {
	Status string `json:"status"`
}
