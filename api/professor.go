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
	var reqQueryParams1 model.ListProfessorsRequest
	var reqQueryParams2 model.SearchProfessorByNameQueryRequest

	err1 := ctx.ShouldBindQuery(&reqQueryParams1)
	err2 := ctx.ShouldBindQuery(&reqQueryParams2)
	if err1 != nil && err2 != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err2))
		return
	}

	if err1 == nil {
		arg := db.ListProfessorsParams{
			Limit:  reqQueryParams1.PageSize,
			Offset: (reqQueryParams1.PageID - 1) * reqQueryParams1.PageSize,
		}

		professors, err := server.query.ListProfessors(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, professors)
		return
	}

	professors, err := server.query.SearchProfessorsByName(ctx, "%"+reqQueryParams2.Name+"%")
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
