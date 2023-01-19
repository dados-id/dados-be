package validation

import (
	"fmt"
	"net/url"

	"github.com/dados-id/dados-be/model"
	"github.com/dados-id/dados-be/util"
)

func ValidateCreateSchoolRequest(req *model.CreateSchoolRequest) (violations []BadRequest_FieldViolation) {
	if err := validateNameSchool(req.Name); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	if err := validateNicknameSchool(req.NickName); err != nil {
		violations = append(violations, fieldViolation("nickName", err))
	}

	if err := validateCity(req.City); err != nil {
		violations = append(violations, fieldViolation("city", err))
	}

	if err := validateProvince(req.Province); err != nil {
		violations = append(violations, fieldViolation("province", err))
	}

	if _, err := validateWebsite(req.Website); err != nil {
		violations = append(violations, fieldViolation("website", err))
	}

	if err := validateEmail(req.Email); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}

func ValidateUpdateSchoolRequest(req *model.UpdateSchoolStatusRequest) (violations []BadRequest_FieldViolation) {
	if err := validateStatusRequest(req.Status); err != nil {
		violations = append(violations, fieldViolation("status", err))
	}

	return violations
}

func validateStatusRequest(value string) error {
	if ok := util.IsSupportedStatusRequest(value); !ok {
		return fmt.Errorf("not supported statusRequest: %s", value)
	}
	return nil
}

func validateNameSchool(value string) error {
	if err := validateString(value, 1, 64); err != nil {
		return err
	}
	if !isValidSchool(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func validateNicknameSchool(value []string) error {
	for _, val := range value {
		if err := validateString(val, 1, 64); err != nil {
			return err
		}
		if !isValidSchool(val) {
			return fmt.Errorf("must contain only letters or spaces")
		}
	}
	return nil
}

func validateCity(value string) error {
	return validateNameSchool(value)
}

func validateProvince(value string) error {
	return validateNameSchool(value)
}

func validateWebsite(value string) (*url.URL, error) {
	return url.ParseRequestURI(value)
}
