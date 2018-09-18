package customer

import "errors"

func validateName(name string) error {
	if name == "" {
		return errors.New("name can not be empty")
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return errors.New("email can not be empty")
	}
	return nil
}

func validateAge(age uint64) error {
	if age < 1 {
		return errors.New("age can not be less than 1")
	}
	return nil
}
func validateGender(gender string) error {
	if gender == "male" || gender == "female" {
		return nil
	}
	return errors.New("gender can be 'male' or 'female'")
}