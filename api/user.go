package api

import (
	"database/sql"
	"net/http"
	"strings"

	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/model"
	"github.com/dados-id/dados-be/validation"
	"github.com/gin-gonic/gin"
)

func (server *Server) loginUser(ctx *gin.Context) {
	var reqJSON model.LoginUserRequest

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	// Verify the ID token passed by the client
	token := reqJSON.IDToken
	tokenInfo, err := server.firebaseClient.VerifyIDToken(ctx, token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, exception.ErrorResponse(err))
		return
	}

	name := tokenInfo.Claims["name"].(string)
	email := tokenInfo.Claims["email"].(string)

	fullName := strings.Fields(name)

	firstName := fullName[0]
	lastName := strings.Join(fullName[1:], " ")

	arg := db.CreateUserParams{
		FirstName:                firstName,
		LastName:                 lastName,
		School:                   "",
		ExpectedYearOfGraduation: 0,
		Email:                    email,
	}

	user, err := server.query.CreateUser(ctx, arg)
	if err != nil {
		if errorConstraint, ok := exception.IsUniqueViolation(err); ok {
			ctx.JSON(http.StatusForbidden, exception.ViolationUniqueConstraint(errorConstraint))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) getUser(ctx *gin.Context) {
	var reqURI model.GetUserRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	User, err := server.query.GetUser(ctx, reqURI.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, exception.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, User)
}

func (server *Server) createUser(ctx *gin.Context) {
	var reqJSON model.CreateUserRequest

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateCreateUserRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.CreateUserParams{
		FirstName:                reqJSON.FirstName,
		LastName:                 reqJSON.LastName,
		School:                   reqJSON.School,
		ExpectedYearOfGraduation: reqJSON.ExpectedYearOfGraduation,
		Email:                    reqJSON.Email,
	}

	user, err := server.query.CreateUser(ctx, arg)
	if err != nil {
		if errorConstraint, ok := exception.IsUniqueViolation(err); ok {
			ctx.JSON(http.StatusForbidden, exception.ViolationUniqueConstraint(errorConstraint))
			return
		}

		ctx.JSON(http.StatusInternalServerError, exception.ServerErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) updateUser(ctx *gin.Context) {
	var reqURI model.UpdateUserURIRequest
	var reqJSON model.UpdateUserJSONRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	violations := validation.ValidateUpdateUserRequest(&reqJSON)
	if violations != nil {
		ctx.JSON(http.StatusBadRequest, exception.ViolationsFieldValidation(violations))
		return
	}

	arg := db.UpdateUserParams{
		ID: reqURI.ID,
		FirstName: sql.NullString{
			String: reqJSON.GetFirstName(),
			Valid:  reqJSON.FirstName != nil,
		},
		LastName: sql.NullString{
			String: reqJSON.GetLastName(),
			Valid:  reqJSON.LastName != nil,
		},
		School: sql.NullString{
			String: reqJSON.GetSchool(),
			Valid:  reqJSON.School != nil,
		},
		ExpectedYearOfGraduation: sql.NullInt16{
			Int16: reqJSON.GetExpectedYearOfGraduation(),
			Valid: reqJSON.ExpectedYearOfGraduation != nil,
		},
	}

	user, err := server.query.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) userListProfessorRatings(ctx *gin.Context) {
	var reqURI model.UserListURIRequest
	var reqQueryParams model.UserListQueryRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.UserListProfessorRatingsParams{
		UserID: reqURI.UserID,
		Limit:  reqQueryParams.PageSize,
		Offset: (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	users, err := server.query.UserListProfessorRatings(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (server *Server) userListSchoolRatings(ctx *gin.Context) {
	var reqURI model.UserListURIRequest
	var reqQueryParams model.UserListQueryRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.UserListSchoolRatingsParams{
		UserID: reqURI.UserID,
		Limit:  reqQueryParams.PageSize,
		Offset: (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	users, err := server.query.UserListSchoolRatings(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (server *Server) userListSavedProfessors(ctx *gin.Context) {
	var reqURI model.UserListURIRequest
	var reqQueryParams model.UserListQueryRequest

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.UserListSavedProfessorsParams{
		UserID: reqURI.UserID,
		Limit:  reqQueryParams.PageSize,
		Offset: (reqQueryParams.PageID - 1) * reqQueryParams.PageSize,
	}

	users, err := server.query.UserListSavedProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (server *Server) saveProfessor(ctx *gin.Context) {
	var reqURI model.SaveProfessorURIRequest
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.SaveProfessorParams{
		UserID:      reqURI.UserID,
		ProfessorID: reqURI.ProfessorID,
	}

	if err := server.query.SaveProfessor(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (server *Server) unsaveProfessor(ctx *gin.Context) {
	var reqURI model.UnsaveProfessorURIRequest
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.UnsaveProfessorParams{
		UserID:      reqURI.UserID,
		ProfessorID: reqURI.ProfessorID,
	}

	if err := server.query.UnsaveProfessor(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
