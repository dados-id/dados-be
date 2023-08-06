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
	SchoolID int32 `uri:"school_id" binding:"required,min=1"`
}

type ListSchoolsQueryRequest struct {
	PageID    int32   `form:"page_id"`
	PageSize  int32   `form:"page_size"`
	Name      *string `form:"name"`
	SortBy    *string `form:"sort_by"`
	SortOrder *string `form:"sort_order"`
}

func (x *ListSchoolsQueryRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ListSchoolsQueryRequest) GetSortBy() string {
	if x != nil && x.SortBy != nil {
		return *x.SortBy
	}
	return ""
}

func (x *ListSchoolsQueryRequest) GetSortOrder() string {
	if x != nil && x.SortOrder != nil {
		return *x.SortOrder
	}
	return ""
}

type UpdateSchoolStatusRequest struct {
	Status string `json:"status"`
}
