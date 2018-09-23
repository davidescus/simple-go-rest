package customer

func validateName(name string) string {
	if name == "" {
		return "name can not be empty"
	}
	return ""
}

func validateEmail(email string) string {
	if email == "" {
		return "email can not be empty"
	}
	return ""
}

func validateAge(age uint64) string {
	if age < 1 {
		return "age can not be less than 1"
	}
	return ""
}
func validateGender(gender string) string {
	if gender == "male" || gender == "female" {
		return ""
	}

	return "gender can be 'male' or 'female'"
}