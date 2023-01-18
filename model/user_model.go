package model

type CreateUserRequest struct {
	FirstName                string `json:"firstName" binding:"required"`
	LastName                 string `json:"lastName" binding:"required"`
	School                   string `json:"school" binding:"required"`
	ExpectedYearOfGraduation int16  `json:"expectedYearOfGraduation" binding:"required"`
	Email                    string `json:"email" binding:"required,email"`
}
