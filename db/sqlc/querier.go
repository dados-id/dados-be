// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateCorrection(ctx context.Context, arg CreateCorrectionParams) (CorrectionForm, error)
	CreateProfessor(ctx context.Context, arg CreateProfessorParams) (Professor, error)
	CreateProfessorCourseAssociation(ctx context.Context, arg CreateProfessorCourseAssociationParams) error
	CreateProfessorRating(ctx context.Context, arg CreateProfessorRatingParams) (ProfessorRating, error)
	CreateReport(ctx context.Context, arg CreateReportParams) (ReportForm, error)
	CreateSchool(ctx context.Context, arg CreateSchoolParams) (School, error)
	CreateSchoolFacultyAssociation(ctx context.Context, arg CreateSchoolFacultyAssociationParams) error
	CreateSchoolRating(ctx context.Context, arg CreateSchoolRatingParams) (SchoolRating, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserSaveProfessor(ctx context.Context, arg CreateUserSaveProfessorParams) error
	DeleteUserSaveProfessor(ctx context.Context, arg DeleteUserSaveProfessorParams) error
	GetProfessor(ctx context.Context, id int64) (Professor, error)
	GetProfessorInfoAggregate(ctx context.Context, id int64) (GetProfessorInfoAggregateRow, error)
	GetProfessorRating(ctx context.Context, arg GetProfessorRatingParams) (GetProfessorRatingRow, error)
	GetSchool(ctx context.Context, id int64) (School, error)
	GetSchoolInfoAggregate(ctx context.Context, id int64) (GetSchoolInfoAggregateRow, error)
	GetSchoolRating(ctx context.Context, arg GetSchoolRatingParams) (GetSchoolRatingRow, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListCorrection(ctx context.Context, arg ListCorrectionParams) ([]CorrectionForm, error)
	ListProfessorRatings(ctx context.Context, arg ListProfessorRatingsParams) ([]ListProfessorRatingsRow, error)
	ListProfessorRatingsJoinProfessorFilterByCourse(ctx context.Context, arg ListProfessorRatingsJoinProfessorFilterByCourseParams) ([]ListProfessorRatingsJoinProfessorFilterByCourseRow, error)
	ListProfessorRatingsJoinProfessorFilterByRating(ctx context.Context, arg ListProfessorRatingsJoinProfessorFilterByRatingParams) ([]ListProfessorRatingsJoinProfessorFilterByRatingRow, error)
	ListProfessors(ctx context.Context, arg ListProfessorsParams) ([]ListProfessorsRow, error)
	ListProfessorsByFaculty(ctx context.Context, arg ListProfessorsByFacultyParams) ([]ListProfessorsByFacultyRow, error)
	ListProfessorsByFacultyAndSchool(ctx context.Context, arg ListProfessorsByFacultyAndSchoolParams) ([]ListProfessorsByFacultyAndSchoolRow, error)
	ListProfessorsBySchool(ctx context.Context, arg ListProfessorsBySchoolParams) ([]ListProfessorsBySchoolRow, error)
	ListReport(ctx context.Context, arg ListReportParams) ([]ReportForm, error)
	ListSchoolRatings(ctx context.Context, arg ListSchoolRatingsParams) ([]ListSchoolRatingsRow, error)
	ListSchools(ctx context.Context, arg ListSchoolsParams) ([]School, error)
	ListTop5Tags(ctx context.Context, professorID int64) ([]string, error)
	SearchProfessorsByName(ctx context.Context, firstName string) ([]Professor, error)
	SearchSchoolsByNameOrNickName(ctx context.Context, name string) ([]School, error)
	UpdateCorrection(ctx context.Context, arg UpdateCorrectionParams) (CorrectionForm, error)
	UpdateProfessorRating(ctx context.Context, arg UpdateProfessorRatingParams) (ProfessorRating, error)
	UpdateProfessorStatusRequest(ctx context.Context, arg UpdateProfessorStatusRequestParams) (Professor, error)
	UpdateReport(ctx context.Context, arg UpdateReportParams) (ReportForm, error)
	UpdateSchoolRating(ctx context.Context, arg UpdateSchoolRatingParams) (SchoolRating, error)
	UpdateSchoolStatusRequest(ctx context.Context, arg UpdateSchoolStatusRequestParams) (School, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UserListProfessorRatings(ctx context.Context, arg UserListProfessorRatingsParams) ([]UserListProfessorRatingsRow, error)
	UserListSavedProfessors(ctx context.Context, arg UserListSavedProfessorsParams) ([]UserListSavedProfessorsRow, error)
	UserListSchoolRatings(ctx context.Context, arg UserListSchoolRatingsParams) ([]UserListSchoolRatingsRow, error)
}

var _ Querier = (*Queries)(nil)