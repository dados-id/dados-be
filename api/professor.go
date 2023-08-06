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

func (server *Server) getProfessorInfo(ctx *gin.Context) {
	var reqURI model.GetProfessorRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	listTopTags, err := server.query.ListTopTags(ctx, reqURI.ProfessorID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	listTopCoursesTaught, err := server.query.ListTopCoursesTaught(ctx, reqURI.ProfessorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	courses, err := server.query.ListCoursesByProfessorId(ctx, reqURI.ProfessorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	professorInfo, err := server.query.GetProfessorInfo(ctx, reqURI.ProfessorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	rsp := model.GetProfessorInfoResponse{
		GetProfessorInfoRow: professorInfo,
		TopTags:             listTopTags,
		TopCoursesTaught:    listTopCoursesTaught,
		Courses:             courses,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) createProfessor(ctx *gin.Context) {
	var reqJSON model.CreateProfessorRequest

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateCreateProfessorRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.CreateProfessorParams{
		FirstName: reqJSON.FirstName,
		LastName:  reqJSON.LastName,
		FacultyID: reqJSON.FacultyID,
		SchoolID:  reqJSON.SchoolID,
	}

	professor, err := server.query.CreateProfessor(ctx, arg)
	if err != nil {
		if errorConstraint, ok := exception.IsUniqueViolation(err); ok {
			ctx.JSON(http.StatusForbidden, exception.ViolationUniqueConstraint(errorConstraint))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, professor)
}

func (server *Server) listProfessors(ctx *gin.Context) {
	var reqQueryParams model.ListProfessorsQueryRequest

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	// Search By Specific Name
	if reqQueryParams.Name != nil {
		arg := db.ListProfessorsByNameParams{
			Limit:     reqQueryParams.PageSize,
			Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
			Name:      "%" + reqQueryParams.GetName() + "%",
			SortBy:    reqQueryParams.GetSortBy(),
			SortOrder: reqQueryParams.GetSortOrder(),
		}

		professors, err := server.query.ListProfessorsByName(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		totalCount, err := server.query.CountListProfessorsByName(ctx, "%"+reqQueryParams.GetName()+"%")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
		ctx.JSON(http.StatusOK, professors)
		return
	}

	arg := db.ListProfessorsParams{
		Limit:     reqQueryParams.PageSize,
		Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
		SortBy:    reqQueryParams.GetSortBy(),
		SortOrder: reqQueryParams.GetSortOrder(),
	}

	professors, err := server.query.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	totalCount, err := server.query.CountListProfessors(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
	ctx.JSON(http.StatusOK, professors)
}

func (server *Server) listProfessorsBySchool(ctx *gin.Context) {
	var reqQueryParams model.ListProfessorsQueryBySchoolRequest
	var reqURI model.ListProfessorsURIBySchoolRequest

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	// Search By Specific Name
	if reqQueryParams.Name != nil {
		arg := db.ListProfessorsBySchoolAndNameParams{
			SchoolID:  reqURI.SchoolID,
			Limit:     reqQueryParams.PageSize,
			Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
			Name:      "%" + reqQueryParams.GetName() + "%",
			SortBy:    reqQueryParams.GetSortBy(),
			SortOrder: reqQueryParams.GetSortOrder(),
		}

		professors, err := server.query.ListProfessorsBySchoolAndName(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		arg2 := db.CountListProfessorsBySchoolAndNameParams{
			SchoolID: reqURI.SchoolID,
			Name:     "%" + reqQueryParams.GetName() + "%",
		}

		totalCount, err := server.query.CountListProfessorsBySchoolAndName(ctx, arg2)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
		ctx.JSON(http.StatusOK, professors)
		return
	}

	arg := db.ListProfessorsBySchoolParams{
		SchoolID:  reqURI.SchoolID,
		Limit:     reqQueryParams.PageSize,
		Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
		SortBy:    reqQueryParams.GetSortBy(),
		SortOrder: reqQueryParams.GetSortOrder(),
	}

	professors, err := server.query.ListProfessorsBySchool(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	totalCount, err := server.query.CountListProfessorsBySchool(ctx, reqURI.SchoolID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
	ctx.JSON(http.StatusOK, professors)
}

func (server *Server) listProfessorsBySchoolAndFaculty(ctx *gin.Context) {
	var reqQueryParams model.ListProfessorsQueryByFacultyRequest
	var reqURI model.ListProfessorsURIByFacultyRequest

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	// Search By Specific Name
	if reqQueryParams.Name != nil {
		arg := db.ListProfessorsBySchoolAndNameAndFacultyParams{
			SchoolID:  reqURI.SchoolID,
			FacultyID: reqURI.FacultyID,
			Limit:     reqQueryParams.PageSize,
			Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
			Name:      "%" + reqQueryParams.GetName() + "%",
			SortBy:    reqQueryParams.GetSortBy(),
			SortOrder: reqQueryParams.GetSortOrder(),
		}

		professors, err := server.query.ListProfessorsBySchoolAndNameAndFaculty(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		arg2 := db.CountListProfessorsBySchoolAndNameAndFacultyParams{
			SchoolID:  reqURI.SchoolID,
			FacultyID: reqURI.FacultyID,
			Name:      "%" + reqQueryParams.GetName() + "%",
		}

		totalCount, err := server.query.CountListProfessorsBySchoolAndNameAndFaculty(ctx, arg2)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
		ctx.JSON(http.StatusOK, professors)
		return
	}

	arg := db.ListProfessorsByFacultyAndSchoolParams{
		FacultyID: reqURI.FacultyID,
		SchoolID:  reqURI.SchoolID,
		Limit:     reqQueryParams.PageSize,
		Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
		SortBy:    reqQueryParams.GetSortBy(),
		SortOrder: reqQueryParams.GetSortOrder(),
	}

	professors, err := server.query.ListProfessorsByFacultyAndSchool(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	arg2 := db.CountListProfessorsByFacultyAndSchoolParams{
		FacultyID: reqURI.FacultyID,
		SchoolID:  reqURI.SchoolID,
	}

	totalCount, err := server.query.CountListProfessorsByFacultyAndSchool(ctx, arg2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
	ctx.JSON(http.StatusOK, professors)
}

func (server *Server) updateProfessorStatusRequest(ctx *gin.Context) {
	var reqURI model.GetProfessorRequest
	var reqJSON model.UpdateProfessorStatusRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateUpdateProfessorRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.UpdateProfessorStatusRequestParams{
		Status: db.Statusrequest(reqJSON.Status),
		ID:     reqURI.ProfessorID,
	}

	professor, err := server.query.UpdateProfessorStatusRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, professor)
}
