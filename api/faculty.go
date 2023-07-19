package api

import (
	"net/http"

	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/model"
	"github.com/dados-id/dados-be/validation"
	"github.com/gin-gonic/gin"
)

func (server *Server) listFacultiesBySchool(ctx *gin.Context) {
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

func (server *Server) createFaculty(ctx *gin.Context) {
	var reqJSON model.CreateFacultyRequest

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse((err)))
		return
	}

	violations := validation.ValidateCreateFacultyRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	name := reqJSON.Name

	faculty, err := server.query.CreateFaculty(ctx, name)
	if err != nil {
		if errorConstraint, ok := exception.IsUniqueViolation(err); ok {
			ctx.JSON(http.StatusForbidden, exception.ViolationUniqueConstraint(errorConstraint))
			return
		}
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	arg := db.CreateSchoolFacultyAssociationParams{
		FacultyID: faculty.ID,
		SchoolID:  reqJSON.SchoolID,
	}
	errSchoolFaculty := server.query.CreateSchoolFacultyAssociation(ctx, arg)
	if errSchoolFaculty != nil {
		if errorConstraintSchoolFaculty, ok := exception.IsUniqueViolation(errSchoolFaculty); ok {
			ctx.JSON(http.StatusForbidden, exception.ViolationUniqueConstraint(errorConstraintSchoolFaculty))
			return
		}
		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(errSchoolFaculty))
		return
	}

	data := db.FacultySchool{
		ID:       faculty.ID,
		Name:     faculty.Name,
		SchoolID: reqJSON.SchoolID,
	}

	ctx.JSON(http.StatusOK, data)
}

// func (ser)

// func (server *Server) createFaculty(ctx *gin.Context) {

// 	var reqJSON model.CreateSchoolRequest

// 	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
// 		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
// 		return
// 	}

// 	violations := validation.ValidateCreateSchoolRequest(&reqJSON)
// 	if violations != nil {
// 		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
// 		return
// 	}

// 	arg := db.CreateSchoolParams{
// 		Name:     reqJSON.Name,
// 		NickName: reqJSON.NickName,
// 		City:     reqJSON.City,
// 		Province: reqJSON.Province,
// 		Website:  reqJSON.Website,
// 		Email:    reqJSON.Email,
// 	}

// 	school, err := server.query.CreateSchool(ctx, arg)
// 	if err != nil {
// 		if errorConstraint, ok := exception.IsUniqueViolation(err); ok {
// 			ctx.JSON(http.StatusForbidden, exception.ViolationUniqueConstraint(errorConstraint))
// 			return
// 		}

// 		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, faculties)
// }
