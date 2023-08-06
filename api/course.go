package api

import (
	"net/http"

	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) listCoursesByProfessorId(ctx *gin.Context) {
	var reqURI model.ListCoursesURIRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	courses, err := server.query.ListCoursesByProfessorId(ctx, reqURI.ProfessorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	if len(courses) == 0 {
		ctx.JSON(http.StatusNotFound, exception.ErrorResponseMessage("professor_id not found"))
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
