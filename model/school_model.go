package model

import db "github.com/dados-id/dados-be/db/sqlc"

type CreateSchoolRequest struct {
	Name     string   `json:"name" binding:"required"`
	NickName []string `json:"nickName" binding:"required"`
	City     string   `json:"city" binding:"required"`
	Province string   `json:"province" binding:"required"`
	Website  string   `json:"website" binding:"required"`
	Email    string   `json:"email" binding:"required,email"`
}

type GetSchoolRequest struct {
	SchoolID int64 `uri:"school_id" binding:"required,min=1"`
}

type ListSchoolsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

type SearchSchoolByNameOrNicknameQueryRequest struct {
	Name string `form:"name" binding:"required,alpha"`
}

type UpdateSchoolStatusRequest struct {
	Status db.Statusrequest `json:"status" binding:"required"`
}
