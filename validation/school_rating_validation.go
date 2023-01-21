package validation

import "github.com/dados-id/dados-be/model"

func ValidateCreateSchoolRatingRequest(req *model.CreateSchoolRatingJSONRequest) (violations []BadRequest_FieldViolation) {
	if err := validateSchoolRatingField(req.Clubs); err != nil {
		violations = append(violations, fieldViolation("clubs", err))
	}

	if err := validateSchoolRatingField(req.Facilities); err != nil {
		violations = append(violations, fieldViolation("facilities", err))
	}

	if err := validateSchoolRatingField(req.Food); err != nil {
		violations = append(violations, fieldViolation("food", err))
	}

	if err := validateSchoolRatingField(req.Happiness); err != nil {
		violations = append(violations, fieldViolation("happiness", err))
	}

	if err := validateSchoolRatingField(req.Internet); err != nil {
		violations = append(violations, fieldViolation("internet", err))
	}

	if err := validateSchoolRatingField(req.Location); err != nil {
		violations = append(violations, fieldViolation("location", err))
	}

	if err := validateSchoolRatingField(req.Opportunities); err != nil {
		violations = append(violations, fieldViolation("opportunities", err))
	}

	if err := validateSchoolRatingField(req.Reputation); err != nil {
		violations = append(violations, fieldViolation("reputation", err))
	}

	if err := validateSchoolRatingField(req.Safety); err != nil {
		violations = append(violations, fieldViolation("safety", err))
	}

	if err := validateSchoolRatingField(req.Social); err != nil {
		violations = append(violations, fieldViolation("social", err))
	}

	if err := validateReview(req.Review); err != nil {
		violations = append(violations, fieldViolation("review", err))
	}

	return violations
}

func ValidateUpdateSchoolRatingRequest(req *model.UpdateSchoolRatingJSONRequest) (violations []BadRequest_FieldViolation) {
	if req.Clubs != nil {
		if err := validateSchoolRatingField(req.GetClubs()); err != nil {
			violations = append(violations, fieldViolation("clubs", err))
		}
	}
	if req.Facilities != nil {
		if err := validateSchoolRatingField(req.GetFacilities()); err != nil {
			violations = append(violations, fieldViolation("facilities", err))
		}
	}

	if req.Food != nil {
		if err := validateSchoolRatingField(req.GetFood()); err != nil {
			violations = append(violations, fieldViolation("food", err))
		}
	}

	if req.Happiness != nil {
		if err := validateSchoolRatingField(req.GetHappiness()); err != nil {
			violations = append(violations, fieldViolation("happiness", err))
		}
	}

	if req.Internet != nil {
		if err := validateSchoolRatingField(req.GetInternet()); err != nil {
			violations = append(violations, fieldViolation("internet", err))
		}
	}

	if req.Location != nil {
		if err := validateSchoolRatingField(req.GetLocation()); err != nil {
			violations = append(violations, fieldViolation("location", err))
		}
	}

	if req.Opportunities != nil {
		if err := validateSchoolRatingField(req.GetOpportunities()); err != nil {
			violations = append(violations, fieldViolation("opportunities", err))
		}
	}

	if req.Reputation != nil {
		if err := validateSchoolRatingField(req.GetReputation()); err != nil {
			violations = append(violations, fieldViolation("reputation", err))
		}
	}

	if req.Safety != nil {
		if err := validateSchoolRatingField(req.GetSafety()); err != nil {
			violations = append(violations, fieldViolation("safety", err))
		}
	}

	if req.Social != nil {
		if err := validateSchoolRatingField(req.GetSocial()); err != nil {
			violations = append(violations, fieldViolation("social", err))
		}
	}

	if req.Review != nil {
		if err := validateReview(req.GetReview()); err != nil {
			violations = append(violations, fieldViolation("review", err))
		}
	}

	return violations
}

func validateSchoolRatingField(value int16) error {
	if err := validateIntNull(int(value)); err != nil {
		return err
	}
	return validateInt(int(value), 1, 5)
}
