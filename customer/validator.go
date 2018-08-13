package customer

func ValidateName(name string) string {
	if name == "" {
		return "Name can not be empty\n"
	}
	return ""
}

func ValidateEmail(email string) string {
	if email == "" {
		return "Email can not be empty"
	}
	return ""
}

func ValidateAge(age uint64) string {
	if age < 1 {
		return "Age can not be less than 1"
	}
	return ""
}
func ValidateGender(gender string) string {
	if gender == "male" || gender == "female" {
		return ""
	}
	return "Gender can be 'male' or 'female'"
}