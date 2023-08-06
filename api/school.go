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

func (server *Server) getSchoolInfo(ctx *gin.Context) {
	var reqURI model.GetSchoolRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	schoolInfo, err := server.query.GetSchoolInfo(ctx, reqURI.SchoolID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schoolInfo)
}

func (server *Server) listSchools(ctx *gin.Context) {
	var reqQueryParams model.ListSchoolsQueryRequest

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	// Search By Specific Name
	if reqQueryParams.Name != nil {
		arg := db.ListSchoolsByNameParams{
			Limit:     reqQueryParams.PageSize,
			Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
			NickName:  reqQueryParams.GetName(),
			Name:      "%" + reqQueryParams.GetName() + "%",
			SortBy:    reqQueryParams.GetSortBy(),
			SortOrder: reqQueryParams.GetSortOrder(),
		}

		schools, err := server.query.ListSchoolsByName(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		arg2 := db.CountListSchoolsByNameParams{
			NickName: reqQueryParams.GetName(),
			Name:     "%" + reqQueryParams.GetName() + "%",
		}

		totalCount, err := server.query.CountListSchoolsByName(ctx, arg2)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
		ctx.JSON(http.StatusOK, schools)
		return
	}

	if reqQueryParams.PageSize != 0 && reqQueryParams.PageID != 0 {
		arg := db.ListSchoolsParams{
			Limit:     reqQueryParams.PageSize,
			Offset:    (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
			SortBy:    reqQueryParams.GetSortBy(),
			SortOrder: reqQueryParams.GetSortOrder(),
		}

		schools, err := server.query.ListSchools(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		totalCount, err := server.query.CountListSchools(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
			return
		}

		ctx.Header("x-total-count", strconv.Itoa(int(totalCount)))
		ctx.JSON(http.StatusOK, schools)
		return
	}

	arg := db.ListSchoolsAllParams{
		SortBy:    reqQueryParams.GetSortBy(),
		SortOrder: reqQueryParams.GetSortOrder(),
	}

	schools, err := server.query.ListSchoolsAll(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school)
}
