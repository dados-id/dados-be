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

func (server *Server) getSchoolInfoAggregate(ctx *gin.Context) {
	var reqURI model.GetSchoolRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	schoolInfo, err := server.query.GetSchoolInfoAggregate(ctx, reqURI.SchoolID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schoolInfo)
}

func (server *Server) listSchools(ctx *gin.Context) {
	var reqQueryParams1 model.ListSchoolsRequest
	var reqQueryParams2 model.SearchSchoolByNameOrNicknameQueryRequest

	err1 := ctx.ShouldBindQuery(&reqQueryParams1)
	err2 := ctx.ShouldBindQuery(&reqQueryParams2)
	if err1 != nil && err2 != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err2))
		return
	}

	if err1 == nil {
		arg := db.ListSchoolsParams{
			Limit:  reqQueryParams1.PageSize,
			Offset: (reqQueryParams1.PageID - 1) * reqQueryParams1.PageSize,
		}

		schools, err := server.query.ListSchools(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, schools)
		return
	}

	schools, err := server.query.SearchSchoolsByNameOrNickName(ctx, "%"+reqQueryParams2.Name+"%")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schools)
}

func (server *Server) createSchool(ctx *gin.Context) {
	var reqJSON model.CreateSchoolRequest

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateCreateSchoolRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.CreateSchoolParams{
		Name:     reqJSON.Name,
		NickName: reqJSON.NickName,
		City:     reqJSON.City,
		Province: reqJSON.Province,
		Website:  reqJSON.Website,
		Email:    reqJSON.Email,
	}

	school, err := server.query.CreateSchool(ctx, arg)
	if err != nil {
		if errorConstraint, ok := exception.IsUniqueViolation(err); ok {
			ctx.JSON(http.StatusForbidden, exception.ViolationUniqueConstraint(errorConstraint))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school)
}

func (server *Server) updateSchoolStatusRequest(ctx *gin.Context) {
	var reqURI model.GetSchoolRequest
	var reqJSON model.UpdateSchoolStatusRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateUpdateSchoolRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.UpdateSchoolStatusRequestParams{
		Status: db.Statusrequest(reqJSON.Status),
		ID:     reqURI.SchoolID,
	}

	school, err := server.query.UpdateSchoolStatusRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school)
}