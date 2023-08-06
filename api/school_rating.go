package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/model"
	"github.com/dados-id/dados-be/validation"
	"github.com/gin-gonic/gin"
)

func (server *Server) getSchoolRating(ctx *gin.Context) {
	var req model.SchoolRatingURIRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.GetSchoolRatingParams{
		SchoolID:       req.SchoolID,
		SchoolRatingID: req.SchoolRatingID,
	}

	school_rating, err := server.query.GetSchoolRating(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school_rating)
}

func (server *Server) listSchoolRatings(ctx *gin.Context) {
	var reqURI model.ListSchoolRatingsURIRequest
	var reqQueryParams model.ListSchoolRatingsQueryRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.ListSchoolRatingsParams{
		SchoolID: reqURI.SchoolID,
		Limit:    reqQueryParams.PageSize,
		Offset:   (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	schoolRatings, err := server.query.ListSchoolRatings(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	totalCount, err := server.query.CountListSchoolRatings(ctx, reqURI.SchoolID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
	ctx.JSON(http.StatusOK, schoolRatings)
}

func (server *Server) createSchoolRating(ctx *gin.Context) {
	var reqJSON model.CreateSchoolRatingJSONRequest
	var reqURI model.CreateSchoolRatingURIRequest

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateCreateSchoolRatingRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(string)

	arg := db.CreateSchoolRatingParams{
		UserID:        userID,
		SchoolID:      reqURI.SchoolID,
		Reputation:    reqJSON.Reputation,
		Location:      reqJSON.Location,
		Opportunities: reqJSON.Opportunities,
		Facilities:    reqJSON.Facilities,
		Internet:      reqJSON.Internet,
		Food:          reqJSON.Food,
		Clubs:         reqJSON.Clubs,
		Social:        reqJSON.Social,
		Happiness:     reqJSON.Happiness,
		Safety:        reqJSON.Safety,
		Review:        reqJSON.Review,
	}

	schoolRating, err := server.query.CreateSchoolRating(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schoolRating)
}

func (server *Server) updateSchoolRating(ctx *gin.Context) {
	var reqURI model.SchoolRatingURIRequest
	var reqJSON model.UpdateSchoolRatingJSONRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateUpdateSchoolRatingRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.UpdateSchoolRatingParams{
		Reputation:     sql.NullInt16{Int16: reqJSON.GetReputation(), Valid: reqJSON.Reputation != nil},
		Location:       sql.NullInt16{Int16: reqJSON.GetLocation(), Valid: reqJSON.Location != nil},
		Opportunities:  sql.NullInt16{Int16: reqJSON.GetOpportunities(), Valid: reqJSON.Opportunities != nil},
		Facilities:     sql.NullInt16{Int16: reqJSON.GetFacilities(), Valid: reqJSON.Facilities != nil},
		Internet:       sql.NullInt16{Int16: reqJSON.GetInternet(), Valid: reqJSON.Internet != nil},
		Food:           sql.NullInt16{Int16: reqJSON.GetFood(), Valid: reqJSON.Food != nil},
		Clubs:          sql.NullInt16{Int16: reqJSON.GetClubs(), Valid: reqJSON.Clubs != nil},
		Social:         sql.NullInt16{Int16: reqJSON.GetSocial(), Valid: reqJSON.Social != nil},
		Happiness:      sql.NullInt16{Int16: reqJSON.GetHappiness(), Valid: reqJSON.Happiness != nil},
		Safety:         sql.NullInt16{Int16: reqJSON.GetSafety(), Valid: reqJSON.Safety != nil},
		Review:         sql.NullString{String: reqJSON.GetReview(), Valid: reqJSON.Review != nil},
		UpVote:         sql.NullInt32{Int32: reqJSON.GetUpVote(), Valid: reqJSON.UpVote != nil},
		DownVote:       sql.NullInt32{Int32: reqJSON.GetDownVote(), Valid: reqJSON.DownVote != nil},
		SchoolRatingID: reqURI.SchoolRatingID,
		SchoolID:       reqURI.SchoolID,
	}

	SchoolRating, err := server.query.UpdateSchoolRating(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SchoolRating)
}
