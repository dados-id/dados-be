package model

type CreateUserRequest struct {
	FirstName                string `json:"first_name" binding:"required"`
	LastName                 string `json:"last_name" binding:"required"`
	School                   string `json:"school" binding:"required"`
	ExpectedYearOfGraduation int16  `json:"expected_year_of_graduation" binding:"required"`
	Email                    string `json:"email" binding:"required,email"`
}
