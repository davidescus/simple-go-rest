package customer

import (
	"simple-go-rest/messages"
	"strconv"
	"net/http"
)

// Store ...
func Store(r *http.Request) (*Customer, messages.Messages) {
	msgr := &messages.Messages{}
	c := Customer{}

	err := r.ParseForm()
	if err != nil {
		msgr.AddError("error When try to parse request", "system")
		return nil, msgr.Get()
	}

	c.Name = r.FormValue("name")
	msgr.AddError("name", validateName(c.Name))

	c.Email = r.FormValue("email")
	msgr.AddError("email", validateEmail(c.Email))

	age, _ := strconv.ParseUint(r.FormValue("age"), 10, 64)
	if err != nil {
		msgr.AddError("age", "invalid age")
	}
	c.Age = age
	msgr.AddError("age", validateAge(c.Age))

	c.Gender = r.FormValue("gender")
	msgr.AddError("gender", validateGender(c.Gender))

	if msgr.GetCount() > 0 {
		return nil, msgr.Get()
	}

	err = c.Store()
	if err != nil {
		msgr.AddError("system", err.Error())
	}

	if msgr.GetCount() > 0 {
		return nil, msgr.Get()
	}

	return &c, nil
}

// GetAll ...
func GetAll() (*Customers, messages.Messages) {
	msgr := &messages.Messages{}
	c := Customers{}

	customers, err := c.All()
	if err != nil {
		msgr.AddError("system", err.Error())
		return customers, msgr.Get()
	}

	return customers, nil
}
