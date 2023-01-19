package util

import db "github.com/dados-id/dados-be/db/sqlc"

// IsSupportedStatusRequest returns true if the StatusRequest is supported
func IsSupportedStatusRequest(statusRequest string) bool {
	switch statusRequest {
	case string(db.StatusrequestPending), string(db.StatusrequestRejected), string(db.StatusrequestVerified):
		return true
	}
	return false
}
