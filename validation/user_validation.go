package validation

import (
	"fmt"

	"github.com/dados-id/dados-be/model"
)

func ValidateCreateUserRequest(req *model.CreateUserRequest) (violations []BadRequest_FieldViolation) {
	if err := validateName(req.FirstName); err != nil {
		violations = append(violations, fieldViolation("firstName", err))
	}

	if err := validateName(req.LastName); err != nil {
		violations = append(violations, fieldViolation("lastName", err))
	}

	if err := validateExpectedYearOfGraduation(int(req.ExpectedYearOfGraduation)); err != nil {
		violations = append(violations, fieldViolation("expectedYearOfGraduation", err))
	}

	if err := validateEmail(req.Email); err != nil {
		violations = append(violations, fieldViolation("email", err))
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
		if err := validateExpectedYearOfGraduation(int(req.GetExpectedYearOfGraduation())); err != nil {
			violations = append(violations, fieldViolation("expectedYearOfGraduation", err))
		}
	}

	if req.School != nil {
		if err := validateSchool(req.GetSchool()); err != nil {
			violations = append(violations, fieldViolation("school", err))
		}
	}

	return violations
}

func validateName(value string) error {
	return validateString(value, 1, 64)
}

func validateExpectedYearOfGraduation(value int) error {
	if err := validateIntNull(value); err != nil {
		return err
	}
	return validateInt(value, 2023, 9999)
}

func validateSchool(value string) error {
	if err := validateString(value, 1, 64); err != nil {
		return err
	}
	if !isValidSchool(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}
