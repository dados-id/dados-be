// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CountListProfessorRatings(ctx context.Context, professorID int32) (int32, error)
	CountListProfessorRatingsFilterByCourse(ctx context.Context, arg CountListProfessorRatingsFilterByCourseParams) (int32, error)
	CountListProfessorRatingsFilterByRating(ctx context.Context, arg CountListProfessorRatingsFilterByRatingParams) (int32, error)
	CountListProfessors(ctx context.Context) (int32, error)
	CountListProfessorsByFaculty(ctx context.Context, facultyID int32) (int32, error)
	CountListProfessorsByFacultyAndSchool(ctx context.Context, arg CountListProfessorsByFacultyAndSchoolParams) (int32, error)
	CountListProfessorsByName(ctx context.Context, name string) (int32, error)
	CountListProfessorsBySchool(ctx context.Context, schoolID int32) (int32, error)
	CountListProfessorsBySchoolAndName(ctx context.Context, arg CountListProfessorsBySchoolAndNameParams) (int32, error)
	CountListProfessorsBySchoolAndNameAndFaculty(ctx context.Context, arg CountListProfessorsBySchoolAndNameAndFacultyParams) (int32, error)
	CountListSchoolRatings(ctx context.Context, schoolID int32) (int32, error)
	CountListSchools(ctx context.Context) (int32, error)
	CountListSchoolsByName(ctx context.Context, arg CountListSchoolsByNameParams) (int32, error)
	CreateCorrection(ctx context.Context, arg CreateCorrectionParams) (CorrectionForm, error)
	CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error)
	CreateFaculty(ctx context.Context, name string) (Faculty, error)
	CreateProfessor(ctx context.Context, arg CreateProfessorParams) (int32, error)
	CreateProfessorCourseAssociation(ctx context.Context, arg CreateProfessorCourseAssociationParams) error
	CreateProfessorRating(ctx context.Context, arg CreateProfessorRatingParams) (ProfessorRating, error)
	CreateProfessorRatingTags(ctx context.Context, arg CreateProfessorRatingTagsParams) error
	CreateReport(ctx context.Context, arg CreateReportParams) (ReportForm, error)
	CreateSchool(ctx context.Context, arg CreateSchoolParams) (int32, error)
	CreateSchoolFacultyAssociation(ctx context.Context, arg CreateSchoolFacultyAssociationParams) error
	CreateSchoolRating(ctx context.Context, arg CreateSchoolRatingParams) (SchoolRating, error)
	CreateTag(ctx context.Context, name string) (string, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetProfessorInfo(ctx context.Context, id int32) (GetProfessorInfoRow, error)
	GetProfessorRating(ctx context.Context, arg GetProfessorRatingParams) (GetProfessorRatingRow, error)
	GetSchoolInfo(ctx context.Context, id int32) (GetSchoolInfoRow, error)
	GetSchoolRating(ctx context.Context, arg GetSchoolRatingParams) (GetSchoolRatingRow, error)
	GetUser(ctx context.Context, id string) (GetUserRow, error)
	ListCorrection(ctx context.Context, arg ListCorrectionParams) ([]CorrectionForm, error)
	ListCoursesByProfessorId(ctx context.Context, professorID int32) ([]string, error)
	ListFacultyBySchool(ctx context.Context, schoolID int32) ([]Faculty, error)
	ListProfessorRatings(ctx context.Context, arg ListProfessorRatingsParams) ([]ListProfessorRatingsRow, error)
	ListProfessorRatingsFilterByCourse(ctx context.Context, arg ListProfessorRatingsFilterByCourseParams) ([]ListProfessorRatingsFilterByCourseRow, error)
	ListProfessorRatingsFilterByRating(ctx context.Context, arg ListProfessorRatingsFilterByRatingParams) ([]ListProfessorRatingsFilterByRatingRow, error)
	ListProfessors(ctx context.Context, arg ListProfessorsParams) ([]ListProfessorsRow, error)
	ListProfessorsByFaculty(ctx context.Context, arg ListProfessorsByFacultyParams) ([]ListProfessorsByFacultyRow, error)
	ListProfessorsByFacultyAndSchool(ctx context.Context, arg ListProfessorsByFacultyAndSchoolParams) ([]ListProfessorsByFacultyAndSchoolRow, error)
	ListProfessorsByName(ctx context.Context, arg ListProfessorsByNameParams) ([]ListProfessorsByNameRow, error)
	ListProfessorsBySchool(ctx context.Context, arg ListProfessorsBySchoolParams) ([]ListProfessorsBySchoolRow, error)
	ListProfessorsBySchoolAndName(ctx context.Context, arg ListProfessorsBySchoolAndNameParams) ([]ListProfessorsBySchoolAndNameRow, error)
	ListProfessorsBySchoolAndNameAndFaculty(ctx context.Context, arg ListProfessorsBySchoolAndNameAndFacultyParams) ([]ListProfessorsBySchoolAndNameAndFacultyRow, error)
	ListRandomCourseCode(ctx context.Context) ([]string, error)
	ListRandomFacultyID(ctx context.Context) ([]int32, error)
	ListRandomTag(ctx context.Context) ([]string, error)
	ListRandomUserID(ctx context.Context) ([]string, error)
	ListReport(ctx context.Context, arg ListReportParams) ([]ReportForm, error)
	ListSchoolRatings(ctx context.Context, arg ListSchoolRatingsParams) ([]ListSchoolRatingsRow, error)
	ListSchools(ctx context.Context, arg ListSchoolsParams) ([]ListSchoolsRow, error)
	ListSchoolsAll(ctx context.Context, arg ListSchoolsAllParams) ([]ListSchoolsAllRow, error)
	ListSchoolsByName(ctx context.Context, arg ListSchoolsByNameParams) ([]ListSchoolsByNameRow, error)
	ListTopCoursesTaught(ctx context.Context, professorID int32) ([]string, error)
	ListTopTags(ctx context.Context, professorID int32) ([]string, error)
	RandomFacultyID(ctx context.Context) (int32, error)
	RandomProfessorID(ctx context.Context) (int32, error)
	RandomSchoolID(ctx context.Context) (int32, error)
	SaveProfessor(ctx context.Context, arg SaveProfessorParams) error
	UnsaveProfessor(ctx context.Context, arg UnsaveProfessorParams) error
	UpdateCorrection(ctx context.Context, arg UpdateCorrectionParams) (CorrectionForm, error)
	UpdateProfessorRating(ctx context.Context, arg UpdateProfessorRatingParams) (ProfessorRating, error)
	UpdateProfessorStatusRequest(ctx context.Context, arg UpdateProfessorStatusRequestParams) (Statusrequest, error)
	UpdateReport(ctx context.Context, arg UpdateReportParams) (ReportForm, error)
	UpdateSchoolRating(ctx context.Context, arg UpdateSchoolRatingParams) (SchoolRating, error)
	UpdateSchoolStatusRequest(ctx context.Context, arg UpdateSchoolStatusRequestParams) (School, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UserListProfessorRatings(ctx context.Context, arg UserListProfessorRatingsParams) ([]UserListProfessorRatingsRow, error)
	UserListSavedProfessors(ctx context.Context, arg UserListSavedProfessorsParams) ([]UserListSavedProfessorsRow, error)
	UserListSchoolRatings(ctx context.Context, arg UserListSchoolRatingsParams) ([]UserListSchoolRatingsRow, error)
}

var _ Querier = (*Queries)(nil)
