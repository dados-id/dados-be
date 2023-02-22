package validation

import (
	"fmt"
	"strconv"

	"github.com/dados-id/dados-be/model"
)

func ValidateCreateProfessorRatingRequest(req *model.CreateProfessorRatingJSONRequest) (violations []BadRequest_FieldViolation) {
	if err := validateQuality(req.Quality); err != nil {
		violations = append(violations, fieldViolation("quality", err))
	}

	if err := validateDifficult(req.Difficult); err != nil {
		violations = append(violations, fieldViolation("difficult", err))
	}

	if err := validateWouldTakeAgain(req.WouldTakeAgain); err != nil {
		violations = append(violations, fieldViolation("wouldTakeAgain", err))
	}

	if err := validateTakenForCredit(req.TakenForCredit); err != nil {
		violations = append(violations, fieldViolation("takenForCredit", err))
	}

	if err := validateUseTextbooks(req.UseTextbooks); err != nil {
		violations = append(violations, fieldViolation("useTextbooks", err))
	}

	if err := validateAttendanceMandatory(req.AttendanceMandatory); err != nil {
		violations = append(violations, fieldViolation("attendanceMandatory", err))
	}

	if err := validateGrade(req.Grade); err != nil {
		violations = append(violations, fieldViolation("Grade", err))
	}

	if err := validateTags(req.Tags); err != nil {
		violations = append(violations, fieldViolation("tags", err))
	}

	if err := validateReview(req.Review); err != nil {
		violations = append(violations, fieldViolation("review", err))
	}

	if err := validateCourseCode(req.CourseCode); err != nil {
		violations = append(violations, fieldViolation("courseCode", err))
	}

	return violations
}

func ValidateUpdateProfessorRatingRequest(req *model.UpdateProfessorRatingJSONRequest) (violations []BadRequest_FieldViolation) {
	if req.Quality != nil {
		if err := validateQuality(req.GetQuality()); err != nil {
			violations = append(violations, fieldViolation("quality", err))
		}
	}
	if req.Difficult != nil {
		if err := validateDifficult(req.GetDifficult()); err != nil {
			violations = append(violations, fieldViolation("difficult", err))
		}
	}

	if req.WouldTakeAgain != nil {
		if err := validateWouldTakeAgain(req.GetWouldTakeAgain()); err != nil {
			violations = append(violations, fieldViolation("wouldTakeAgain", err))
		}
	}

	if req.TakenForCredit != nil {
		if err := validateTakenForCredit(req.GetTakenForCredit()); err != nil {
			violations = append(violations, fieldViolation("takenForCredit", err))
		}
	}

	if req.UseTextbooks != nil {
		if err := validateUseTextbooks(req.GetUseTextbooks()); err != nil {
			violations = append(violations, fieldViolation("useTextbooks", err))
		}
	}

	if req.AttendanceMandatory != nil {
		if err := validateAttendanceMandatory(req.GetAttendanceMandatory()); err != nil {
			violations = append(violations, fieldViolation("attendanceMandatory", err))
		}
	}

	if req.Grade != nil {
		if err := validateGrade(req.GetGrade()); err != nil {
			violations = append(violations, fieldViolation("grade", err))
		}
	}

	if req.Tags != nil {
		if err := validateTags(req.GetTags()); err != nil {
			violations = append(violations, fieldViolation("tags", err))
		}
	}

	if req.Review != nil {
		if err := validateReview(req.GetReview()); err != nil {
			violations = append(violations, fieldViolation("review", err))
		}
	}

	if req.CourseCode != nil {
		if err := validateCourseCode(req.GetCourseCode()); err != nil {
			violations = append(violations, fieldViolation("courseCode", err))
		}
	}

	return violations
}

func validateQuality(value string) error {
	num, err := parseStringToInt(value)
	if err != nil {
		return err
	}
	if err := validateInt(num, 1, 5); err != nil {
		return err
	}
	return nil
}

func validateDifficult(value string) error {
	return validateQuality(value)
}

func validateWouldTakeAgain(value int16) error {
	return validateInt(int(value), 0, 1)
}

func validateTakenForCredit(value int16) error {
	return validateInt(int(value), 0, 2)
}

func validateUseTextbooks(value int16) error {
	return validateInt(int(value), 0, 2)
}

func validateAttendanceMandatory(value int16) error {
	return validateInt(int(value), 0, 2)
}

func validateGrade(value string) error {
	return validateStringNull(value)
}

func validateTags(value []string) error {
	for _, str := range value {
		if err := validateStringNull(str); err != nil {
			return err
		}
	}
	return nil
}

func validateCourseCode(value string) error {
	return validateStringNull(value)
}

func parseStringToInt(value string) (int, error) {
	if num, err := strconv.Atoi(value); err == nil {
		return num, nil
	}
	return 0, fmt.Errorf("%q not a number", value)
}
