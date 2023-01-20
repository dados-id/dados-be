package validation

import (
	"math"

	"github.com/dados-id/dados-be/model"
)

func ValidateCreateProfessorRequest(req *model.CreateProfessorRequest) (violations []BadRequest_FieldViolation) {
	if err := validateName(req.FirstName); err != nil {
		violations = append(violations, fieldViolation("firstName", err))
	}

	if err := validateName(req.LastName); err != nil {
		violations = append(violations, fieldViolation("lastName", err))
	}

	if err := validateInt64(req.FacultyID, 1, math.MaxInt64); err != nil {
		violations = append(violations, fieldViolation("facultyID", err))
	}

	if err := validateInt64(req.SchoolID, 1, math.MaxInt64); err != nil {
		violations = append(violations, fieldViolation("schoolID", err))
	}

	return violations
}

func ValidateUpdateProfessorRequest(req *model.UpdateProfessorStatusRequest) (violations []BadRequest_FieldViolation) {
	if err := validateStatusRequest(req.Status); err != nil {
		violations = append(violations, fieldViolation("status", err))
	}

	return violations
}
