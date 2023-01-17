// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type StatusRequest string

const (
	StatusRequestPending  StatusRequest = "pending"
	StatusRequestVerified StatusRequest = "verified"
	StatusRequestRejected StatusRequest = "rejected"
)

func (e *StatusRequest) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StatusRequest(s)
	case string:
		*e = StatusRequest(s)
	default:
		return fmt.Errorf("unsupported scan type for StatusRequest: %T", src)
	}
	return nil
}

type NullStatusRequest struct {
	StatusRequest StatusRequest
	Valid         bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatusRequest) Scan(value interface{}) error {
	if value == nil {
		ns.StatusRequest, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StatusRequest.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatusRequest) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.StatusRequest, nil
}

type CorrectionForm struct {
	ID           int64          `json:"id"`
	Problem      sql.NullString `json:"problem"`
	CorrectInfo  sql.NullString `json:"correct_info"`
	Email        sql.NullString `json:"email"`
	Status       interface{}    `json:"status"`
	RequestDate  time.Time      `json:"request_date"`
	VerifiedDate time.Time      `json:"verified_date"`
	UserID       int64          `json:"user_id"`
}

type Course struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Faculty struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//
//       List of derived attribute:
//       1. top tags
//       2. 5 distribusi nilai (from quality)
//
type Professor struct {
	ID                int64       `json:"id"`
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Rating            int16       `json:"rating"`
	TotalReview       int32       `json:"total_review"`
	WouldTakeAgain    int16       `json:"would_take_again"`
	LevelOfDifficulty string      `json:"level_of_difficulty"`
	CreatedAt         time.Time   `json:"created_at"`
	Status            interface{} `json:"status"`
	VerifiedDate      time.Time   `json:"verified_date"`
	FacultyID         int64       `json:"faculty_id"`
	SchoolID          int64       `json:"school_id"`
}

type ProfessorCourseAssociation struct {
	ProfessorID int64  `json:"professor_id"`
	CourseCode  string `json:"course_code"`
}

//
//       attendance_mandatory:
//       1. true
//       2. false
//       3. unknown
//
//       would_take_again:
//       0. false
//       1. true
//
type ProfessorRating struct {
	ID                  int64          `json:"id"`
	Quality             string         `json:"quality"`
	Difficult           string         `json:"difficult"`
	WouldTakeAgain      int16          `json:"would_take_again"`
	TakenForCredit      sql.NullBool   `json:"taken_for_credit"`
	UseTextbooks        sql.NullBool   `json:"use_textbooks"`
	AttendanceMandatory int16          `json:"attendance_mandatory"`
	Grade               sql.NullString `json:"grade"`
	Tags                []string       `json:"tags"`
	Review              string         `json:"review"`
	UpVote              int32          `json:"up_vote"`
	DownVote            int32          `json:"down_vote"`
	CreatedAt           time.Time      `json:"created_at"`
	EditedAt            time.Time      `json:"edited_at"`
	ProfessorID         int64          `json:"professor_id"`
	CourseCode          string         `json:"course_code"`
	UserID              int64          `json:"user_id"`
	Verified            sql.NullBool   `json:"verified"`
}

type ProfessorRatingTag struct {
	TagID       int64 `json:"tag_id"`
	ProfessorID int64 `json:"professor_id"`
}

type ReportForm struct {
	ID                int64       `json:"id"`
	Comment           string      `json:"comment"`
	Status            interface{} `json:"status"`
	RequestDate       time.Time   `json:"request_date"`
	VerifiedDate      time.Time   `json:"verified_date"`
	ProfessorRatingID int64       `json:"professor_rating_id"`
	UserID            int64       `json:"user_id"`
}

//
//       List of derived attribute:
//       1. 10 avg field school rating
//       2. avg Overall Quality
//
type School struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	NickName     []string    `json:"nick_name"`
	Country      string      `json:"country"`
	Province     string      `json:"province"`
	Website      string      `json:"website"`
	Email        string      `json:"email"`
	Status       interface{} `json:"status"`
	VerifiedDate time.Time   `json:"verified_date"`
}

type SchoolFacultyAssociation struct {
	FacultyID int64 `json:"faculty_id"`
	SchoolID  int64 `json:"school_id"`
}

type SchoolRating struct {
	ID            int64        `json:"id"`
	UserID        int64        `json:"user_id"`
	SchoolID      int64        `json:"school_id"`
	Reputation    int16        `json:"reputation"`
	Location      int16        `json:"location"`
	Opportunities int16        `json:"opportunities"`
	Facilities    int16        `json:"facilities"`
	Internet      int16        `json:"internet"`
	Food          int16        `json:"food"`
	Clubs         int16        `json:"clubs"`
	Social        int16        `json:"social"`
	Happiness     int16        `json:"happiness"`
	Safety        int16        `json:"safety"`
	Review        string       `json:"review"`
	UpVote        int32        `json:"up_vote"`
	DownVote      int32        `json:"down_vote"`
	OverallRating string       `json:"overall_rating"`
	CreatedAt     time.Time    `json:"created_at"`
	EditedAt      time.Time    `json:"edited_at"`
	Verified      sql.NullBool `json:"verified"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID                       int64          `json:"id"`
	FirstName                sql.NullString `json:"first_name"`
	LastName                 sql.NullString `json:"last_name"`
	School                   sql.NullString `json:"school"`
	ExpectedYearOfGraduation sql.NullInt16  `json:"expected_year_of_graduation"`
	Email                    string         `json:"email"`
	CreatedAt                time.Time      `json:"created_at"`
}

type UserSaveProfessor struct {
	ProfessorID int64 `json:"professor_id"`
	UserID      int64 `json:"user_id"`
}
