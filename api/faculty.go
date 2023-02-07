package api

import (
	"net/http"

	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) listFaculties(ctx *gin.Context) {
	var reqURI model.ListFacultyURIRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	faculties, err := server.query.ListFacultyBySchool(ctx, reqURI.SchoolID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, faculties)
}
