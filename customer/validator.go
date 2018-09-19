package customer

import (
	"simple-go-rest/messages"
)

var messenger messages.Messenger

// SetMessenger ...
func SetMessenger(m messages.Messenger) {
	messenger = m
}

func validateName(name string) {
	if name == "" {
		messenger.AddError("name can not be empty", "name")
	}
}

func validateEmail(email string) {
	if email == "" {
		messenger.AddError("email can not be empty", "email")
	}
}

func validateAge(age uint64) {
	if age < 1 {
		messenger.AddError("age can not be less than 1", "age")
	}
}
func validateGender(gender string) {
	if gender == "male" || gender == "female" {
		return
	}

	messenger.AddError("gender can be 'male' or 'female'", "gender")
}