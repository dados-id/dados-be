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

func (server *Server) getProfessorRating(ctx *gin.Context) {
	var reqURI model.ProfessorRatingRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.GetProfessorRatingParams{
		ProfessorID:       reqURI.ProfessorID,
		ProfessorRatingID: reqURI.ProfessorRatingID,
	}

	professorRatingInfo, err := server.query.GetProfessorRating(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, professorRatingInfo)
}

func (server *Server) listProfessorRatings(ctx *gin.Context) {
	var reqURI model.ListProfessorRatingURIRequest
	var reqQueryParams model.ListProfessorRatingQueryRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	// Filter By CourseCode
	if reqQueryParams.CourseCode != nil {
		arg := db.ListProfessorRatingsFilterByCourseParams{
			ProfessorID: reqURI.ProfessorID,
			CourseCode:  reqQueryParams.GetCourseCode(),
			Limit:       reqQueryParams.PageSize,
			Offset:      (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
		}

		professorRatings, err := server.query.ListProfessorRatingsFilterByCourse(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		arg2 := db.CountListProfessorRatingsFilterByCourseParams{
			ProfessorID: reqURI.ProfessorID,
			CourseCode:  reqQueryParams.GetCourseCode(),
		}

		totalCount, err := server.query.CountListProfessorRatingsFilterByCourse(ctx, arg2)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
		ctx.JSON(http.StatusOK, professorRatings)
		return
	}

	// Filter By Rating
	if reqQueryParams.Rating != nil {
		arg := db.ListProfessorRatingsFilterByRatingParams{
			ProfessorID: reqURI.ProfessorID,
			Rating:      reqQueryParams.GetRating(),
			Limit:       reqQueryParams.PageSize,
			Offset:      (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
		}

		professorRatings, err := server.query.ListProfessorRatingsFilterByRating(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		arg2 := db.CountListProfessorRatingsFilterByRatingParams{
			ProfessorID: reqURI.ProfessorID,
			Rating:      reqQueryParams.GetRating(),
		}

		totalCount, err := server.query.CountListProfessorRatingsFilterByRating(ctx, arg2)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
		ctx.JSON(http.StatusOK, professorRatings)
		return
	}

	// TODO: Filter By Date

	// No Filter
	arg := db.ListProfessorRatingsParams{
		ProfessorID: reqURI.ProfessorID,
		Limit:       reqQueryParams.PageSize,
		Offset:      (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	professorRatings, err := server.query.ListProfessorRatings(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	totalCount, err := server.query.CountListProfessorRatings(ctx, reqURI.ProfessorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
	ctx.JSON(http.StatusOK, professorRatings)
}

func (server *Server) createProfessorRating(ctx *gin.Context) {
	var reqJSON model.CreateProfessorRatingJSONRequest
	var reqURI model.CreateProfessorRatingURIRequest

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateCreateProfessorRatingRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(string)

	arg := db.CreateProfessorRatingParams{
		Quality:             reqJSON.Quality,
		Difficult:           reqJSON.Difficult,
		WouldTakeAgain:      reqJSON.WouldTakeAgain,
		TakenForCredit:      reqJSON.TakenForCredit,
		UseTextbooks:        reqJSON.UseTextbooks,
		AttendanceMandatory: reqJSON.AttendanceMandatory,
		Grade:               reqJSON.Grade,
		Review:              reqJSON.Review,
		CourseCode:          reqJSON.CourseCode,
		UserID:              userID,
		ProfessorID:         reqURI.ProfessorID,
	}

	professorRating, err := server.query.CreateProfessorRating(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	for i := 0; i < len(reqJSON.Tags); i++ {
		arg := db.CreateProfessorRatingTagsParams{
			TagName:           reqJSON.Tags[i],
			ProfessorRatingID: professorRating.ID,
		}
		err := server.query.CreateProfessorRatingTags(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, professorRating)
}

func (server *Server) updateProfessorRating(ctx *gin.Context) {
	var reqURI model.ProfessorRatingRequest
	var reqJSON model.UpdateProfessorRatingJSONRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateUpdateProfessorRatingRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.UpdateProfessorRatingParams{
		ProfessorID:         reqURI.ProfessorID,
		Quality:             sql.NullString{String: reqJSON.GetQuality(), Valid: reqJSON.Quality != nil},
		Difficult:           sql.NullString{String: reqJSON.GetDifficult(), Valid: reqJSON.Difficult != nil},
		WouldTakeAgain:      sql.NullInt16{Int16: reqJSON.GetWouldTakeAgain(), Valid: reqJSON.WouldTakeAgain != nil},
		TakenForCredit:      sql.NullInt16{Int16: reqJSON.GetTakenForCredit(), Valid: reqJSON.TakenForCredit != nil},
		UseTextbooks:        sql.NullInt16{Int16: reqJSON.GetUseTextbooks(), Valid: reqJSON.UseTextbooks != nil},
		AttendanceMandatory: sql.NullInt16{Int16: reqJSON.GetAttendanceMandatory(), Valid: reqJSON.AttendanceMandatory != nil},
		Grade:               sql.NullString{String: reqJSON.GetGrade(), Valid: reqJSON.Grade != nil},
		CourseCode:          sql.NullString{String: reqJSON.GetCourseCode(), Valid: reqJSON.CourseCode != nil},
		Review:              sql.NullString{String: reqJSON.GetReview(), Valid: reqJSON.Review != nil},
		UpVote:              sql.NullInt32{Int32: reqJSON.GetUpVote(), Valid: reqJSON.UpVote != nil},
		DownVote:            sql.NullInt32{Int32: reqJSON.GetDownVote(), Valid: reqJSON.DownVote != nil},
	}

	professorRating, err := server.query.UpdateProfessorRating(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, professorRating)
}
