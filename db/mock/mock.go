// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dados-id/dados-be/db/sqlc (interfaces: Querier)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/dados-id/dados-be/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CountFaculty mocks base method.
func (m *MockQuerier) CountFaculty(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountFaculty", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountFaculty indicates an expected call of CountFaculty.
func (mr *MockQuerierMockRecorder) CountFaculty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountFaculty", reflect.TypeOf((*MockQuerier)(nil).CountFaculty), arg0)
}

// CountProfessor mocks base method.
func (m *MockQuerier) CountProfessor(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProfessor", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProfessor indicates an expected call of CountProfessor.
func (mr *MockQuerierMockRecorder) CountProfessor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProfessor", reflect.TypeOf((*MockQuerier)(nil).CountProfessor), arg0)
}

// CountSchool mocks base method.
func (m *MockQuerier) CountSchool(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountSchool", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountSchool indicates an expected call of CountSchool.
func (mr *MockQuerierMockRecorder) CountSchool(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountSchool", reflect.TypeOf((*MockQuerier)(nil).CountSchool), arg0)
}

// CountUser mocks base method.
func (m *MockQuerier) CountUser(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUser indicates an expected call of CountUser.
func (mr *MockQuerierMockRecorder) CountUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUser", reflect.TypeOf((*MockQuerier)(nil).CountUser), arg0)
}

// CreateCorrection mocks base method.
func (m *MockQuerier) CreateCorrection(arg0 context.Context, arg1 db.CreateCorrectionParams) (db.CorrectionForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCorrection", arg0, arg1)
	ret0, _ := ret[0].(db.CorrectionForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCorrection indicates an expected call of CreateCorrection.
func (mr *MockQuerierMockRecorder) CreateCorrection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCorrection", reflect.TypeOf((*MockQuerier)(nil).CreateCorrection), arg0, arg1)
}

// CreateCourse mocks base method.
func (m *MockQuerier) CreateCourse(arg0 context.Context, arg1 db.CreateCourseParams) (db.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCourse", arg0, arg1)
	ret0, _ := ret[0].(db.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCourse indicates an expected call of CreateCourse.
func (mr *MockQuerierMockRecorder) CreateCourse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCourse", reflect.TypeOf((*MockQuerier)(nil).CreateCourse), arg0, arg1)
}

// CreateFaculty mocks base method.
func (m *MockQuerier) CreateFaculty(arg0 context.Context, arg1 string) (db.Faculty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFaculty", arg0, arg1)
	ret0, _ := ret[0].(db.Faculty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFaculty indicates an expected call of CreateFaculty.
func (mr *MockQuerierMockRecorder) CreateFaculty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFaculty", reflect.TypeOf((*MockQuerier)(nil).CreateFaculty), arg0, arg1)
}

// CreateProfessor mocks base method.
func (m *MockQuerier) CreateProfessor(arg0 context.Context, arg1 db.CreateProfessorParams) (db.Professor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfessor", arg0, arg1)
	ret0, _ := ret[0].(db.Professor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfessor indicates an expected call of CreateProfessor.
func (mr *MockQuerierMockRecorder) CreateProfessor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfessor", reflect.TypeOf((*MockQuerier)(nil).CreateProfessor), arg0, arg1)
}

// CreateProfessorCourseAssociation mocks base method.
func (m *MockQuerier) CreateProfessorCourseAssociation(arg0 context.Context, arg1 db.CreateProfessorCourseAssociationParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfessorCourseAssociation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfessorCourseAssociation indicates an expected call of CreateProfessorCourseAssociation.
func (mr *MockQuerierMockRecorder) CreateProfessorCourseAssociation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfessorCourseAssociation", reflect.TypeOf((*MockQuerier)(nil).CreateProfessorCourseAssociation), arg0, arg1)
}

// CreateProfessorRating mocks base method.
func (m *MockQuerier) CreateProfessorRating(arg0 context.Context, arg1 db.CreateProfessorRatingParams) (db.ProfessorRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfessorRating", arg0, arg1)
	ret0, _ := ret[0].(db.ProfessorRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfessorRating indicates an expected call of CreateProfessorRating.
func (mr *MockQuerierMockRecorder) CreateProfessorRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfessorRating", reflect.TypeOf((*MockQuerier)(nil).CreateProfessorRating), arg0, arg1)
}

// CreateProfessorRatingTags mocks base method.
func (m *MockQuerier) CreateProfessorRatingTags(arg0 context.Context, arg1 db.CreateProfessorRatingTagsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfessorRatingTags", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfessorRatingTags indicates an expected call of CreateProfessorRatingTags.
func (mr *MockQuerierMockRecorder) CreateProfessorRatingTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfessorRatingTags", reflect.TypeOf((*MockQuerier)(nil).CreateProfessorRatingTags), arg0, arg1)
}

// CreateReport mocks base method.
func (m *MockQuerier) CreateReport(arg0 context.Context, arg1 db.CreateReportParams) (db.ReportForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReport", arg0, arg1)
	ret0, _ := ret[0].(db.ReportForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReport indicates an expected call of CreateReport.
func (mr *MockQuerierMockRecorder) CreateReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReport", reflect.TypeOf((*MockQuerier)(nil).CreateReport), arg0, arg1)
}

// CreateSchool mocks base method.
func (m *MockQuerier) CreateSchool(arg0 context.Context, arg1 db.CreateSchoolParams) (db.School, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchool", arg0, arg1)
	ret0, _ := ret[0].(db.School)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchool indicates an expected call of CreateSchool.
func (mr *MockQuerierMockRecorder) CreateSchool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchool", reflect.TypeOf((*MockQuerier)(nil).CreateSchool), arg0, arg1)
}

// CreateSchoolFacultyAssociation mocks base method.
func (m *MockQuerier) CreateSchoolFacultyAssociation(arg0 context.Context, arg1 db.CreateSchoolFacultyAssociationParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchoolFacultyAssociation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSchoolFacultyAssociation indicates an expected call of CreateSchoolFacultyAssociation.
func (mr *MockQuerierMockRecorder) CreateSchoolFacultyAssociation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchoolFacultyAssociation", reflect.TypeOf((*MockQuerier)(nil).CreateSchoolFacultyAssociation), arg0, arg1)
}

// CreateSchoolRating mocks base method.
func (m *MockQuerier) CreateSchoolRating(arg0 context.Context, arg1 db.CreateSchoolRatingParams) (db.SchoolRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchoolRating", arg0, arg1)
	ret0, _ := ret[0].(db.SchoolRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchoolRating indicates an expected call of CreateSchoolRating.
func (mr *MockQuerierMockRecorder) CreateSchoolRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchoolRating", reflect.TypeOf((*MockQuerier)(nil).CreateSchoolRating), arg0, arg1)
}

// CreateTag mocks base method.
func (m *MockQuerier) CreateTag(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockQuerierMockRecorder) CreateTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockQuerier)(nil).CreateTag), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), arg0, arg1)
}

// GetProfessorInfoAggregate mocks base method.
func (m *MockQuerier) GetProfessorInfoAggregate(arg0 context.Context, arg1 int64) (db.GetProfessorInfoAggregateRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfessorInfoAggregate", arg0, arg1)
	ret0, _ := ret[0].(db.GetProfessorInfoAggregateRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfessorInfoAggregate indicates an expected call of GetProfessorInfoAggregate.
func (mr *MockQuerierMockRecorder) GetProfessorInfoAggregate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfessorInfoAggregate", reflect.TypeOf((*MockQuerier)(nil).GetProfessorInfoAggregate), arg0, arg1)
}

// GetProfessorRating mocks base method.
func (m *MockQuerier) GetProfessorRating(arg0 context.Context, arg1 db.GetProfessorRatingParams) (db.GetProfessorRatingRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfessorRating", arg0, arg1)
	ret0, _ := ret[0].(db.GetProfessorRatingRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfessorRating indicates an expected call of GetProfessorRating.
func (mr *MockQuerierMockRecorder) GetProfessorRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfessorRating", reflect.TypeOf((*MockQuerier)(nil).GetProfessorRating), arg0, arg1)
}

// GetSchoolInfoAggregate mocks base method.
func (m *MockQuerier) GetSchoolInfoAggregate(arg0 context.Context, arg1 int64) (db.GetSchoolInfoAggregateRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchoolInfoAggregate", arg0, arg1)
	ret0, _ := ret[0].(db.GetSchoolInfoAggregateRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchoolInfoAggregate indicates an expected call of GetSchoolInfoAggregate.
func (mr *MockQuerierMockRecorder) GetSchoolInfoAggregate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchoolInfoAggregate", reflect.TypeOf((*MockQuerier)(nil).GetSchoolInfoAggregate), arg0, arg1)
}

// GetSchoolRating mocks base method.
func (m *MockQuerier) GetSchoolRating(arg0 context.Context, arg1 db.GetSchoolRatingParams) (db.GetSchoolRatingRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchoolRating", arg0, arg1)
	ret0, _ := ret[0].(db.GetSchoolRatingRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchoolRating indicates an expected call of GetSchoolRating.
func (mr *MockQuerierMockRecorder) GetSchoolRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchoolRating", reflect.TypeOf((*MockQuerier)(nil).GetSchoolRating), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockQuerier) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockQuerierMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockQuerier)(nil).GetUser), arg0, arg1)
}

// ListCorrection mocks base method.
func (m *MockQuerier) ListCorrection(arg0 context.Context, arg1 db.ListCorrectionParams) ([]db.CorrectionForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCorrection", arg0, arg1)
	ret0, _ := ret[0].([]db.CorrectionForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCorrection indicates an expected call of ListCorrection.
func (mr *MockQuerierMockRecorder) ListCorrection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCorrection", reflect.TypeOf((*MockQuerier)(nil).ListCorrection), arg0, arg1)
}

// ListCoursesByProfessorId mocks base method.
func (m *MockQuerier) ListCoursesByProfessorId(arg0 context.Context, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCoursesByProfessorId", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCoursesByProfessorId indicates an expected call of ListCoursesByProfessorId.
func (mr *MockQuerierMockRecorder) ListCoursesByProfessorId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCoursesByProfessorId", reflect.TypeOf((*MockQuerier)(nil).ListCoursesByProfessorId), arg0, arg1)
}

// ListProfessorRatings mocks base method.
func (m *MockQuerier) ListProfessorRatings(arg0 context.Context, arg1 db.ListProfessorRatingsParams) ([]db.ListProfessorRatingsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProfessorRatings", arg0, arg1)
	ret0, _ := ret[0].([]db.ListProfessorRatingsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfessorRatings indicates an expected call of ListProfessorRatings.
func (mr *MockQuerierMockRecorder) ListProfessorRatings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfessorRatings", reflect.TypeOf((*MockQuerier)(nil).ListProfessorRatings), arg0, arg1)
}

// ListProfessorRatingsFilterByCourse mocks base method.
func (m *MockQuerier) ListProfessorRatingsFilterByCourse(arg0 context.Context, arg1 db.ListProfessorRatingsFilterByCourseParams) ([]db.ListProfessorRatingsFilterByCourseRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProfessorRatingsFilterByCourse", arg0, arg1)
	ret0, _ := ret[0].([]db.ListProfessorRatingsFilterByCourseRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfessorRatingsFilterByCourse indicates an expected call of ListProfessorRatingsFilterByCourse.
func (mr *MockQuerierMockRecorder) ListProfessorRatingsFilterByCourse(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfessorRatingsFilterByCourse", reflect.TypeOf((*MockQuerier)(nil).ListProfessorRatingsFilterByCourse), arg0, arg1)
}

// ListProfessorRatingsFilterByRating mocks base method.
func (m *MockQuerier) ListProfessorRatingsFilterByRating(arg0 context.Context, arg1 db.ListProfessorRatingsFilterByRatingParams) ([]db.ListProfessorRatingsFilterByRatingRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProfessorRatingsFilterByRating", arg0, arg1)
	ret0, _ := ret[0].([]db.ListProfessorRatingsFilterByRatingRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfessorRatingsFilterByRating indicates an expected call of ListProfessorRatingsFilterByRating.
func (mr *MockQuerierMockRecorder) ListProfessorRatingsFilterByRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfessorRatingsFilterByRating", reflect.TypeOf((*MockQuerier)(nil).ListProfessorRatingsFilterByRating), arg0, arg1)
}

// ListProfessors mocks base method.
func (m *MockQuerier) ListProfessors(arg0 context.Context, arg1 db.ListProfessorsParams) ([]db.ListProfessorsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProfessors", arg0, arg1)
	ret0, _ := ret[0].([]db.ListProfessorsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfessors indicates an expected call of ListProfessors.
func (mr *MockQuerierMockRecorder) ListProfessors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfessors", reflect.TypeOf((*MockQuerier)(nil).ListProfessors), arg0, arg1)
}

// ListProfessorsByFaculty mocks base method.
func (m *MockQuerier) ListProfessorsByFaculty(arg0 context.Context, arg1 db.ListProfessorsByFacultyParams) ([]db.ListProfessorsByFacultyRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProfessorsByFaculty", arg0, arg1)
	ret0, _ := ret[0].([]db.ListProfessorsByFacultyRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfessorsByFaculty indicates an expected call of ListProfessorsByFaculty.
func (mr *MockQuerierMockRecorder) ListProfessorsByFaculty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfessorsByFaculty", reflect.TypeOf((*MockQuerier)(nil).ListProfessorsByFaculty), arg0, arg1)
}

// ListProfessorsByFacultyAndSchool mocks base method.
func (m *MockQuerier) ListProfessorsByFacultyAndSchool(arg0 context.Context, arg1 db.ListProfessorsByFacultyAndSchoolParams) ([]db.ListProfessorsByFacultyAndSchoolRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProfessorsByFacultyAndSchool", arg0, arg1)
	ret0, _ := ret[0].([]db.ListProfessorsByFacultyAndSchoolRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfessorsByFacultyAndSchool indicates an expected call of ListProfessorsByFacultyAndSchool.
func (mr *MockQuerierMockRecorder) ListProfessorsByFacultyAndSchool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfessorsByFacultyAndSchool", reflect.TypeOf((*MockQuerier)(nil).ListProfessorsByFacultyAndSchool), arg0, arg1)
}

// ListProfessorsBySchool mocks base method.
func (m *MockQuerier) ListProfessorsBySchool(arg0 context.Context, arg1 db.ListProfessorsBySchoolParams) ([]db.ListProfessorsBySchoolRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProfessorsBySchool", arg0, arg1)
	ret0, _ := ret[0].([]db.ListProfessorsBySchoolRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProfessorsBySchool indicates an expected call of ListProfessorsBySchool.
func (mr *MockQuerierMockRecorder) ListProfessorsBySchool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProfessorsBySchool", reflect.TypeOf((*MockQuerier)(nil).ListProfessorsBySchool), arg0, arg1)
}

// ListReport mocks base method.
func (m *MockQuerier) ListReport(arg0 context.Context, arg1 db.ListReportParams) ([]db.ReportForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReport", arg0, arg1)
	ret0, _ := ret[0].([]db.ReportForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReport indicates an expected call of ListReport.
func (mr *MockQuerierMockRecorder) ListReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReport", reflect.TypeOf((*MockQuerier)(nil).ListReport), arg0, arg1)
}

// ListSchoolRatings mocks base method.
func (m *MockQuerier) ListSchoolRatings(arg0 context.Context, arg1 db.ListSchoolRatingsParams) ([]db.ListSchoolRatingsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchoolRatings", arg0, arg1)
	ret0, _ := ret[0].([]db.ListSchoolRatingsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchoolRatings indicates an expected call of ListSchoolRatings.
func (mr *MockQuerierMockRecorder) ListSchoolRatings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchoolRatings", reflect.TypeOf((*MockQuerier)(nil).ListSchoolRatings), arg0, arg1)
}

// ListSchools mocks base method.
func (m *MockQuerier) ListSchools(arg0 context.Context, arg1 db.ListSchoolsParams) ([]db.ListSchoolsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchools", arg0, arg1)
	ret0, _ := ret[0].([]db.ListSchoolsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchools indicates an expected call of ListSchools.
func (mr *MockQuerierMockRecorder) ListSchools(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchools", reflect.TypeOf((*MockQuerier)(nil).ListSchools), arg0, arg1)
}

// ListTagsByProfessorRatingId mocks base method.
func (m *MockQuerier) ListTagsByProfessorRatingId(arg0 context.Context, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsByProfessorRatingId", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsByProfessorRatingId indicates an expected call of ListTagsByProfessorRatingId.
func (mr *MockQuerierMockRecorder) ListTagsByProfessorRatingId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsByProfessorRatingId", reflect.TypeOf((*MockQuerier)(nil).ListTagsByProfessorRatingId), arg0, arg1)
}

// ListTop5Tags mocks base method.
func (m *MockQuerier) ListTop5Tags(arg0 context.Context, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTop5Tags", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTop5Tags indicates an expected call of ListTop5Tags.
func (mr *MockQuerierMockRecorder) ListTop5Tags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTop5Tags", reflect.TypeOf((*MockQuerier)(nil).ListTop5Tags), arg0, arg1)
}

// RandomCourseCode mocks base method.
func (m *MockQuerier) RandomCourseCode(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomCourseCode", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RandomCourseCode indicates an expected call of RandomCourseCode.
func (mr *MockQuerierMockRecorder) RandomCourseCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomCourseCode", reflect.TypeOf((*MockQuerier)(nil).RandomCourseCode), arg0)
}

// RandomTag mocks base method.
func (m *MockQuerier) RandomTag(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomTag", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RandomTag indicates an expected call of RandomTag.
func (mr *MockQuerierMockRecorder) RandomTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomTag", reflect.TypeOf((*MockQuerier)(nil).RandomTag), arg0)
}

// SaveProfessor mocks base method.
func (m *MockQuerier) SaveProfessor(arg0 context.Context, arg1 db.SaveProfessorParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveProfessor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveProfessor indicates an expected call of SaveProfessor.
func (mr *MockQuerierMockRecorder) SaveProfessor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProfessor", reflect.TypeOf((*MockQuerier)(nil).SaveProfessor), arg0, arg1)
}

// SearchProfessorsByName mocks base method.
func (m *MockQuerier) SearchProfessorsByName(arg0 context.Context, arg1 string) ([]db.SearchProfessorsByNameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProfessorsByName", arg0, arg1)
	ret0, _ := ret[0].([]db.SearchProfessorsByNameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProfessorsByName indicates an expected call of SearchProfessorsByName.
func (mr *MockQuerierMockRecorder) SearchProfessorsByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProfessorsByName", reflect.TypeOf((*MockQuerier)(nil).SearchProfessorsByName), arg0, arg1)
}

// SearchSchoolsByNameOrNickName mocks base method.
func (m *MockQuerier) SearchSchoolsByNameOrNickName(arg0 context.Context, arg1 db.SearchSchoolsByNameOrNickNameParams) ([]db.SearchSchoolsByNameOrNickNameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchSchoolsByNameOrNickName", arg0, arg1)
	ret0, _ := ret[0].([]db.SearchSchoolsByNameOrNickNameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchSchoolsByNameOrNickName indicates an expected call of SearchSchoolsByNameOrNickName.
func (mr *MockQuerierMockRecorder) SearchSchoolsByNameOrNickName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSchoolsByNameOrNickName", reflect.TypeOf((*MockQuerier)(nil).SearchSchoolsByNameOrNickName), arg0, arg1)
}

// UnsaveProfessor mocks base method.
func (m *MockQuerier) UnsaveProfessor(arg0 context.Context, arg1 db.UnsaveProfessorParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsaveProfessor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsaveProfessor indicates an expected call of UnsaveProfessor.
func (mr *MockQuerierMockRecorder) UnsaveProfessor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsaveProfessor", reflect.TypeOf((*MockQuerier)(nil).UnsaveProfessor), arg0, arg1)
}

// UpdateCorrection mocks base method.
func (m *MockQuerier) UpdateCorrection(arg0 context.Context, arg1 db.UpdateCorrectionParams) (db.CorrectionForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCorrection", arg0, arg1)
	ret0, _ := ret[0].(db.CorrectionForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCorrection indicates an expected call of UpdateCorrection.
func (mr *MockQuerierMockRecorder) UpdateCorrection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCorrection", reflect.TypeOf((*MockQuerier)(nil).UpdateCorrection), arg0, arg1)
}

// UpdateProfessorRating mocks base method.
func (m *MockQuerier) UpdateProfessorRating(arg0 context.Context, arg1 db.UpdateProfessorRatingParams) (db.ProfessorRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfessorRating", arg0, arg1)
	ret0, _ := ret[0].(db.ProfessorRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfessorRating indicates an expected call of UpdateProfessorRating.
func (mr *MockQuerierMockRecorder) UpdateProfessorRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfessorRating", reflect.TypeOf((*MockQuerier)(nil).UpdateProfessorRating), arg0, arg1)
}

// UpdateProfessorStatusRequest mocks base method.
func (m *MockQuerier) UpdateProfessorStatusRequest(arg0 context.Context, arg1 db.UpdateProfessorStatusRequestParams) (db.Professor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfessorStatusRequest", arg0, arg1)
	ret0, _ := ret[0].(db.Professor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfessorStatusRequest indicates an expected call of UpdateProfessorStatusRequest.
func (mr *MockQuerierMockRecorder) UpdateProfessorStatusRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfessorStatusRequest", reflect.TypeOf((*MockQuerier)(nil).UpdateProfessorStatusRequest), arg0, arg1)
}

// UpdateReport mocks base method.
func (m *MockQuerier) UpdateReport(arg0 context.Context, arg1 db.UpdateReportParams) (db.ReportForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReport", arg0, arg1)
	ret0, _ := ret[0].(db.ReportForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReport indicates an expected call of UpdateReport.
func (mr *MockQuerierMockRecorder) UpdateReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReport", reflect.TypeOf((*MockQuerier)(nil).UpdateReport), arg0, arg1)
}

// UpdateSchoolRating mocks base method.
func (m *MockQuerier) UpdateSchoolRating(arg0 context.Context, arg1 db.UpdateSchoolRatingParams) (db.SchoolRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchoolRating", arg0, arg1)
	ret0, _ := ret[0].(db.SchoolRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSchoolRating indicates an expected call of UpdateSchoolRating.
func (mr *MockQuerierMockRecorder) UpdateSchoolRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchoolRating", reflect.TypeOf((*MockQuerier)(nil).UpdateSchoolRating), arg0, arg1)
}

// UpdateSchoolStatusRequest mocks base method.
func (m *MockQuerier) UpdateSchoolStatusRequest(arg0 context.Context, arg1 db.UpdateSchoolStatusRequestParams) (db.School, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchoolStatusRequest", arg0, arg1)
	ret0, _ := ret[0].(db.School)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSchoolStatusRequest indicates an expected call of UpdateSchoolStatusRequest.
func (mr *MockQuerierMockRecorder) UpdateSchoolStatusRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchoolStatusRequest", reflect.TypeOf((*MockQuerier)(nil).UpdateSchoolStatusRequest), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockQuerier) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockQuerierMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockQuerier)(nil).UpdateUser), arg0, arg1)
}

// UserListProfessorRatings mocks base method.
func (m *MockQuerier) UserListProfessorRatings(arg0 context.Context, arg1 db.UserListProfessorRatingsParams) ([]db.UserListProfessorRatingsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserListProfessorRatings", arg0, arg1)
	ret0, _ := ret[0].([]db.UserListProfessorRatingsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserListProfessorRatings indicates an expected call of UserListProfessorRatings.
func (mr *MockQuerierMockRecorder) UserListProfessorRatings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserListProfessorRatings", reflect.TypeOf((*MockQuerier)(nil).UserListProfessorRatings), arg0, arg1)
}

// UserListSavedProfessors mocks base method.
func (m *MockQuerier) UserListSavedProfessors(arg0 context.Context, arg1 db.UserListSavedProfessorsParams) ([]db.UserListSavedProfessorsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserListSavedProfessors", arg0, arg1)
	ret0, _ := ret[0].([]db.UserListSavedProfessorsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserListSavedProfessors indicates an expected call of UserListSavedProfessors.
func (mr *MockQuerierMockRecorder) UserListSavedProfessors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserListSavedProfessors", reflect.TypeOf((*MockQuerier)(nil).UserListSavedProfessors), arg0, arg1)
}

// UserListSchoolRatings mocks base method.
func (m *MockQuerier) UserListSchoolRatings(arg0 context.Context, arg1 db.UserListSchoolRatingsParams) ([]db.UserListSchoolRatingsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserListSchoolRatings", arg0, arg1)
	ret0, _ := ret[0].([]db.UserListSchoolRatingsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserListSchoolRatings indicates an expected call of UserListSchoolRatings.
func (mr *MockQuerierMockRecorder) UserListSchoolRatings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserListSchoolRatings", reflect.TypeOf((*MockQuerier)(nil).UserListSchoolRatings), arg0, arg1)
}
