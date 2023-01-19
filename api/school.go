package api

import (
	"database/sql"
	"net/http"

	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/model"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) getSchoolInfoAggregate(ctx *gin.Context) {
	var req model.GetSchoolRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	schoolInfo, err := server.query.GetSchoolInfoAggregate(ctx, req.SchoolID)
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
	var req model.CreateSchoolRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.CreateSchoolParams{
		Name:     req.Name,
		NickName: req.NickName,
		City:     req.City,
		Province: req.Province,
		Website:  req.Website,
		Email:    req.Email,
	}

	school, err := server.query.CreateSchool(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, exception.ErrorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
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

	arg := db.UpdateSchoolStatusRequestParams{
		Status: reqJSON.Status,
		ID:     reqURI.SchoolID,
	}

	school, err := server.query.UpdateSchoolStatusRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school)
}
