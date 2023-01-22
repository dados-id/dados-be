package validation

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/dados-id/dados-be/util"
)

var (
	isValidSchool = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

type BadRequest_FieldViolation struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

type BadRequest struct {
	FieldViolations []*BadRequest_FieldViolation
}

func fieldViolation(field string, err error) BadRequest_FieldViolation {
	return BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func validateStringNull(value string) error {
	if len(value) == 0 {
		return fmt.Errorf("must not null value")
	}
	return nil
}

func validateIntNull(value int) error {
	if value == 0 {
		return fmt.Errorf("must not null value")
	}
	return nil
}

func validateInt64Null(value int64) error {
	if value == 0 {
		return fmt.Errorf("must not null value")
	}
	return nil
}

func validateString(value string, minLength, maxLength int) error {
	if err := validateStringNull(value); err != nil {
		return err
	}
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func validateInt(value int, minValue, maxValue int) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("value must between %d to %d", minValue, maxValue)
	}
	return nil
}

func validateInt64(value int64, minValue, maxValue int64) error {
	if err := validateInt64Null(value); err != nil {
		return err
	}
	if value < minValue || value > maxValue {
		return fmt.Errorf("value must between %d to %d", minValue, maxValue)
	}
	return nil
}

func validateEmail(value string) error {
	if err := validateString(value, 6, 30); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func validateStatusRequest(value string) error {
	if err := validateStringNull(value); err != nil {
		return err
	}
	if ok := util.IsSupportedStatusRequest(value); !ok {
		return fmt.Errorf("not supported statusRequest: %s", value)
	}
	return nil
}

func validateReview(value string) error {
	return validateString(value, 1, 256)
}
