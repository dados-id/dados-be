package model

type SchoolRatingURIRequest struct {
	SchoolID       int32 `uri:"school_id" binding:"required,min=1"`
	SchoolRatingID int32 `uri:"school_rating_id" binding:"required,min=1"`
}

type ListSchoolRatingsQueryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type ListSchoolRatingsURIRequest struct {
	SchoolID int32 `uri:"school_id" binding:"required,min=1"`
}

type CreateSchoolRatingURIRequest struct {
	SchoolID int32 `uri:"school_id" binding:"required,min=1"`
}

type CreateSchoolRatingJSONRequest struct {
	Reputation    int16  `json:"reputation"`
	Location      int16  `json:"location"`
	Opportunities int16  `json:"opportunities"`
	Facilities    int16  `json:"facilities"`
	Internet      int16  `json:"internet"`
	Food          int16  `json:"food"`
	Clubs         int16  `json:"clubs"`
	Social        int16  `json:"social"`
	Happiness     int16  `json:"happiness"`
	Safety        int16  `json:"safety"`
	Review        string `json:"review"`
}

type UpdateSchoolRatingJSONRequest struct {
	Reputation    *int16  `json:"reputation"`
	Location      *int16  `json:"location"`
	Opportunities *int16  `json:"opportunities"`
	Facilities    *int16  `json:"facilities"`
	Internet      *int16  `json:"internet"`
	Food          *int16  `json:"food"`
	Clubs         *int16  `json:"clubs"`
	Social        *int16  `json:"social"`
	Happiness     *int16  `json:"happiness"`
	Safety        *int16  `json:"safety"`
	Review        *string `json:"review"`
	UpVote        *int32  `json:"upVote"`
	DownVote      *int32  `json:"downVote"`
}

func (x *UpdateSchoolRatingJSONRequest) GetReputation() int16 {
	if x != nil && x.Reputation != nil {
		return *x.Reputation
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetLocation() int16 {
	if x != nil && x.Location != nil {
		return *x.Location
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetOpportunities() int16 {
	if x != nil && x.Opportunities != nil {
		return *x.Opportunities
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetFacilities() int16 {
	if x != nil && x.Facilities != nil {
		return *x.Facilities
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetInternet() int16 {
	if x != nil && x.Internet != nil {
		return *x.Internet
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetFood() int16 {
	if x != nil && x.Food != nil {
		return *x.Food
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetClubs() int16 {
	if x != nil && x.Clubs != nil {
		return *x.Clubs
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetSocial() int16 {
	if x != nil && x.Social != nil {
		return *x.Social
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetHappiness() int16 {
	if x != nil && x.Happiness != nil {
		return *x.Happiness
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetSafety() int16 {
	if x != nil && x.Safety != nil {
		return *x.Safety
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetReview() string {
	if x != nil && x.Review != nil {
		return *x.Review
	}
	return ""
}

func (x *UpdateSchoolRatingJSONRequest) GetUpVote() int32 {
	if x != nil && x.UpVote != nil {
		return *x.UpVote
	}
	return 0
}

func (x *UpdateSchoolRatingJSONRequest) GetDownVote() int32 {
	if x != nil && x.DownVote != nil {
		return *x.DownVote
	}
	return 0
}
