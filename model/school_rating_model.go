package model

type SchoolRatingURIRequest struct {
	SchoolID       int64 `uri:"school_id" binding:"required,min=1"`
	SchoolRatingID int64 `uri:"school_rating_id" binding:"required,min=1"`
}

type ListSchoolRatingsQueryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

type ListSchoolRatingsURIRequest struct {
	SchoolID int64 `uri:"school_id" binding:"required,min=1"`
}

type CreateSchoolRatingURIRequest struct {
	SchoolID int64 `uri:"school_id" binding:"required,min=1"`
}

type CreateSchoolRatingJSONRequest struct {
	UserID        int64  `json:"userID" binding:"required"`
	Reputation    int16  `json:"reputation" binding:"required,min=1,max=5"`
	Location      int16  `json:"location" binding:"required,min=1,max=5"`
	Opportunities int16  `json:"opportunities" binding:"required,min=1,max=5"`
	Facilities    int16  `json:"facilities" binding:"required,min=1,max=5"`
	Internet      int16  `json:"internet" binding:"required,min=1,max=5"`
	Food          int16  `json:"food" binding:"required,min=1,max=5"`
	Clubs         int16  `json:"clubs" binding:"required,min=1,max=5"`
	Social        int16  `json:"social" binding:"required,min=1,max=5"`
	Happiness     int16  `json:"happiness" binding:"required,min=1,max=5"`
	Safety        int16  `json:"safety" binding:"required,min=1,max=5"`
	Review        string `json:"review" binding:"required"`
}

type UpdateSchoolRatingJSONRequest struct {
	Reputation    *int16  `json:"reputation" binding:"max=5"`
	Location      *int16  `json:"location" binding:"max=5"`
	Opportunities *int16  `json:"opportunities" binding:"max=5"`
	Facilities    *int16  `json:"facilities" binding:"max=5"`
	Internet      *int16  `json:"internet" binding:"max=5"`
	Food          *int16  `json:"food" binding:"max=5"`
	Clubs         *int16  `json:"clubs" binding:"max=5"`
	Social        *int16  `json:"social" binding:"max=5"`
	Happiness     *int16  `json:"happiness" binding:"max=5"`
	Safety        *int16  `json:"safety" binding:"max=5"`
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
