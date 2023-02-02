package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/dados-id/dados-be/db/mock"
	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetProfessorInfoAPI(t *testing.T) {
	professor := util.GetValidProfessor(1, 1)

	testCases := []struct {
		name          string
		professorID   int32
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			professorID: professor.ID,
			buildStubs: func(store *mockdb.MockQuerier) {
				professorExpected := db.GetProfessorInfoRow{
					ID:        professor.ID,
					FirstName: professor.FirstName,
					LastName:  professor.LastName,
				}

				store.EXPECT().ListTopTags(gomock.Any(), gomock.Eq(professor.ID)).Times(1)
				store.EXPECT().ListTopCoursesTaught(gomock.Any(), gomock.Eq(professor.ID)).Times(1)
				store.EXPECT().ListCoursesByProfessorId(gomock.Any(), gomock.Eq(professor.ID)).Times(1)

				store.EXPECT().
					GetProfessorInfo(gomock.Any(), gomock.Eq(professor.ID)).
					Times(1).
					Return(professorExpected, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				professorExpected := db.GetProfessorInfoRow{
					ID:        professor.ID,
					FirstName: professor.FirstName,
					LastName:  professor.LastName,
				}
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProfessorInfo(t, recorder.Body, professorExpected)
			},
		},
		{
			name:        "NotFound",
			professorID: professor.ID,

			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					ListTopTags(gomock.Any(), gomock.Eq(professor.ID)).
					Times(1).
					Return([]string{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InternalError",
			professorID: professor.ID,

			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().ListTopTags(gomock.Any(), gomock.Eq(professor.ID)).Times(1)
				store.EXPECT().ListTopCoursesTaught(gomock.Any(), gomock.Eq(professor.ID)).Times(1)
				store.EXPECT().ListCoursesByProfessorId(gomock.Any(), gomock.Eq(professor.ID)).Times(1)

				store.EXPECT().
					GetProfessorInfo(gomock.Any(), gomock.Eq(professor.ID)).
					Times(1).
					Return(db.GetProfessorInfoRow{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:        "InvalidID",
			professorID: 0,

			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().ListTopTags(gomock.Any(), gomock.Eq(professor.ID)).Times(0)
				store.EXPECT().ListTopCoursesTaught(gomock.Any(), gomock.Eq(professor.ID)).Times(0)
				store.EXPECT().ListCoursesByProfessorId(gomock.Any(), gomock.Eq(professor.ID)).Times(0)

				store.EXPECT().
					GetProfessorInfo(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/professors/%d", tc.professorID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListProfessorsAPI(t *testing.T) {
	n := 5
	professors := make([]db.Professor, n)
	for i := 0; i < n; i++ {
		professors[i] = util.GetValidProfessor(1, 1)
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockQuerier) {
				arg := db.ListProfessorsParams{
					Limit:  int32(n),
					Offset: 0,
				}

				professorsExpected := make([]db.ListProfessorsRow, n)
				for i := 0; i < n; i++ {
					professorsExpected[i] = db.ListProfessorsRow{
						ID:        professors[i].ID,
						FirstName: professors[i].FirstName,
						LastName:  professors[i].LastName,
					}
				}

				store.EXPECT().
					ListProfessors(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(professorsExpected, nil)

				store.EXPECT().CountListProfessors(gomock.Any()).Times(1).Return(int32(5), nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				professorsExpected := make([]db.ListProfessorsRow, n)
				for i := 0; i < n; i++ {
					professorsExpected[i] = db.ListProfessorsRow{
						ID:        professors[i].ID,
						FirstName: professors[i].FirstName,
						LastName:  professors[i].LastName,
					}
				}
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProfessors(t, recorder.Body, professorsExpected)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					ListProfessors(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.ListProfessorsRow{}, sql.ErrConnDone)

				store.EXPECT().CountListProfessors(gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					ListProfessors(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().CountListProfessors(gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					ListProfessors(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().CountListProfessors(gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/professors"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchProfessorInfo(t *testing.T, body *bytes.Buffer, professor db.GetProfessorInfoRow) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProfessor db.GetProfessorInfoRow
	err = json.Unmarshal(data, &gotProfessor)
	require.NoError(t, err)
	require.Equal(t, professor, gotProfessor)
}

func requireBodyMatchProfessors(t *testing.T, body *bytes.Buffer, professors []db.ListProfessorsRow) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var potProfessors []db.ListProfessorsRow
	err = json.Unmarshal(data, &potProfessors)
	require.NoError(t, err)
	require.Equal(t, professors, potProfessors)
}
