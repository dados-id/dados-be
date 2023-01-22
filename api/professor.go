package api

import (
	"database/sql"
	"net/http"

	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/model"
	"github.com/dados-id/dados-be/validation"
	"github.com/gin-gonic/gin"
)

func (server *Server) getProfessorInfoAggregate(ctx *gin.Context) {
	var reqURI model.GetProfessorRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	listTop5Tags, err := server.query.ListTop5Tags(ctx, reqURI.ProfessorID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	courses, err := server.query.ListCoursesByProfessorId(ctx, reqURI.ProfessorID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	professorInfo, err := server.query.GetProfessorInfoAggregate(ctx, reqURI.ProfessorID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	rsp := model.GetProfessorInfoResponse{
		GetProfessorInfoAggregateRow: professorInfo,
		Top5Tags:                     listTop5Tags,
		Courses:                      courses,
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

	if reqQueryParams.Name != nil {
		professors, err := server.query.SearchProfessorsByName(ctx, "%"+reqQueryParams.GetName()+"%")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, professors)
		return
	}

	arg := db.ListProfessorsParams{
		Limit:  reqQueryParams.PageSize,
		Offset: (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	professors, err := server.query.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

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

	arg := db.ListProfessorsBySchoolParams{
		SchoolID: reqURI.SchoolID,
		Limit:    reqQueryParams.PageSize,
		Offset:   (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	professors, err := server.query.ListProfessorsBySchool(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, professors)
}

func (server *Server) listProfessorsByFaculty(ctx *gin.Context) {
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

	arg := db.ListProfessorsByFacultyParams{
		FacultyID: reqURI.FacultyID,
		Limit:     reqQueryParams.PageSize,
		Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	professors, err := server.query.ListProfessorsByFaculty(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

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
