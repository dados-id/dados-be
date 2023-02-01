package api

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
	"time"

	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/stretchr/testify/require"
)

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user, gotUser)
	require.Equal(t, user.ID, gotUser.ID)
	require.Equal(t, user.FirstName, gotUser.FirstName)
	require.Equal(t, user.LastName, gotUser.LastName)
	require.Equal(t, user.ExpectedYearOfGraduation, gotUser.ExpectedYearOfGraduation)
	require.Equal(t, user.SchoolID, gotUser.SchoolID)
	require.Equal(t, user.Email, gotUser.Email)
	require.WithinDuration(t, user.CreatedAt, gotUser.CreatedAt, time.Second)
}

// func TestGetUserAPI(t *testing.T) {
// 	school := util.GetValidSchool()
// 	user := util.GetValidUser(school.ID)

// 	testCases := []struct {
// 		name          string
// 		userID        string
// 		buildStubs    func(query *mockdb.MockQuerier)
// 		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:   "OK",
// 			userID: user.ID,
// 			buildStubs: func(query *mockdb.MockQuerier) {
// 				query.EXPECT().
// 					GetUser(gomock.Any(), gomock.Eq(user.ID)).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchUser(t, recorder.Body, user)
// 			},
// 		},
// 		{
// 			name:   "NotFound",
// 			userID: user.ID,
// 			buildStubs: func(store *mockdb.MockQuerier) {
// 				store.EXPECT().
// 					GetUser(gomock.Any(), gomock.Eq(user.ID)).
// 					Times(1).
// 					Return(db.User{}, sql.ErrNoRows)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name:   "InternalError",
// 			userID: user.ID,
// 			buildStubs: func(store *mockdb.MockQuerier) {
// 				store.EXPECT().
// 					GetUser(gomock.Any(), gomock.Eq(user.ID)).
// 					Times(1).
// 					Return(db.User{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		{
// 			name:   "InvalidID",
// 			userID: "",
// 			buildStubs: func(store *mockdb.MockQuerier) {
// 				store.EXPECT().
// 					GetUser(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			query := mockdb.NewMockQuerier(ctrl)
// 			tc.buildStubs(query)

// 			server := newTestServer(t, query)
// 			recorder := httptest.NewRecorder()

// 			url := fmt.Sprintf("/users/%s", tc.userID)
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func TestCreateUserAPI(t *testing.T) {
// 	school := util.GetValidSchool()
// 	user := util.GetValidUser(school.ID)

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(query *mockdb.MockQuerier)
// 		checkResponse func(recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"firstName":                user.FirstName,
// 				"lastName":                 user.LastName,
// 				"school":                   user.SchoolID,
// 				"expectedYearOfGraduation": user.ExpectedYearOfGraduation,
// 				"email":                    user.Email,
// 			},
// 			buildStubs: func(query *mockdb.MockQuerier) {
// 				query.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Any()).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchUser(t, recorder.Body, user)
// 			},
// 		},
// 		{
// 			name: "InternalError",
// 			body: gin.H{
// 				"firstName":                user.FirstName,
// 				"lastName":                 user.LastName,
// 				"school":                   user.SchoolID,
// 				"expectedYearOfGraduation": user.ExpectedYearOfGraduation,
// 				"email":                    user.Email,
// 			},
// 			buildStubs: func(query *mockdb.MockQuerier) {
// 				query.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Any()).
// 					Times(1).
// 					Return(db.User{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "DuplicateEmail",
// 			body: gin.H{
// 				"firstName":                user.FirstName,
// 				"lastName":                 user.LastName,
// 				"school":                   user.SchoolID,
// 				"expectedYearOfGraduation": user.ExpectedYearOfGraduation,
// 				"email":                    user.Email,
// 			},
// 			buildStubs: func(query *mockdb.MockQuerier) {
// 				query.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Any()).
// 					Times(1).
// 					Return(db.User{}, &pq.Error{Code: "23505"})
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusForbidden, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "InvalidEmail",
// 			body: gin.H{
// 				"firstName":                user.FirstName,
// 				"lastName":                 user.LastName,
// 				"school":                   user.SchoolID,
// 				"expectedYearOfGraduation": user.ExpectedYearOfGraduation,
// 				"email":                    "invalid-email",
// 			},
// 			buildStubs: func(query *mockdb.MockQuerier) {
// 				query.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			query := mockdb.NewMockQuerier(ctrl)
// 			tc.buildStubs(query)

// 			server := newTestServer(t, query)
// 			recorder := httptest.NewRecorder()

// 			// Marshal body data to JSON
// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			url := "/users"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(recorder)
// 		})
// 	}
// }
