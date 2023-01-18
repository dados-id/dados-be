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

func (server *Server) getUser(ctx *gin.Context) {
	var req model.GetUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	User, err := server.query.GetUser(ctx, req.ID)
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
	var req model.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		FirstName:                req.FirstName,
		LastName:                 req.LastName,
		School:                   req.School,
		ExpectedYearOfGraduation: req.ExpectedYearOfGraduation,
		Email:                    req.Email,
	}

	user, err := server.query.CreateUser(ctx, arg)
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
	var req model.SaveProfessorURIRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.SaveProfessorParams{
		UserID:      req.UserID,
		ProfessorID: req.ProfessorID,
	}

	if err := server.query.SaveProfessor(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (server *Server) unsaveProfessor(ctx *gin.Context) {
	var req model.UnsaveProfessorURIRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.ErrorResponse(err))
		return
	}

	arg := db.UnsaveProfessorParams{
		UserID:      req.UserID,
		ProfessorID: req.ProfessorID,
	}

	if err := server.query.UnsaveProfessor(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.ErrorResponse(err))
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
