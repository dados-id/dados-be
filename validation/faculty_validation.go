package validation

import (
	"fmt"
	"math"

	"github.com/dados-id/dados-be/model"
)

func ValidateCreateFacultyRequest(req *model.CreateFacultyRequest) (violations []BadRequest_FieldViolation) {
	if err := validateFacultyName(req.Name); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := validateInt32(req.SchoolID, 1, math.MaxInt32); err != nil {
		violations = append(violations, fieldViolation("schoolID", err))
	}

	return violations
}

func validateFacultyName(value string) error {
	if err := validateString(value, 1, 64); err != nil {
		return err
	}
	if !isValidSchool(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}
