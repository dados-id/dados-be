package validation

import (
	"github.com/dados-id/dados-be/model"
)

func ValidateCreateUserRequest(req *model.CreateUserRequest) (violations []BadRequest_FieldViolation) {
	if err := validateName(req.FirstName); err != nil {
		violations = append(violations, fieldViolation("firstName", err))
	}

	if err := validateName(req.LastName); err != nil {
		violations = append(violations, fieldViolation("lastName", err))
	}

	if err := validateExpectedYearOfGraduation(int32(req.ExpectedYearOfGraduation)); err != nil {
		violations = append(violations, fieldViolation("expectedYearOfGraduation", err))
	}

	if err := validateEmail(req.Email); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := validateSchool(req.SchoolID); err != nil {
		violations = append(violations, fieldViolation("school", err))
	}

	return violations
}

func ValidateUpdateUserRequest(req *model.UpdateUserJSONRequest) (violations []BadRequest_FieldViolation) {
	if req.FirstName != nil {
		if err := validateName(req.GetFirstName()); err != nil {
			violations = append(violations, fieldViolation("firstName", err))
		}
	}

	if req.LastName != nil {
		if err := validateName(req.GetLastName()); err != nil {
			violations = append(violations, fieldViolation("lastName", err))
		}
	}

	if req.ExpectedYearOfGraduation != nil {
		if err := validateExpectedYearOfGraduation(int32(req.GetExpectedYearOfGraduation())); err != nil {
			violations = append(violations, fieldViolation("expectedYearOfGraduation", err))
		}
	}

	if req.SchoolID != nil {
		if err := validateSchool(req.GetSchoolID()); err != nil {
			violations = append(violations, fieldViolation("school", err))
		}
	}

	return violations
}

func validateName(value string) error {
	return validateString(value, 1, 64)
}

func validateExpectedYearOfGraduation(value int32) error {
	if err := validateInt32Null(value); err != nil {
		return err
	}
	return validateInt32(value, 2023, 9999)
}

func validateSchool(value int32) error {
	if err := validateInt32Null(value); err != nil {
		return err
	}

	return nil
}
