package customer

import (
	"simple-go-rest/messages"
	"strconv"
	"net/http"
	"simple-go-rest/persistence"
)

func storeNewCustomer(r *http.Request, c *Customer) (*Customer, messages.Messages) {
	messageCollection := &messages.Messages{}

	err := r.ParseForm()
	if err != nil {
		messageCollection.AddError("error When try to parse request", "system")
		return nil, messageCollection.Get()
	}

	c.Name = r.FormValue("name")
	err = validateName(c.Name)
	if err != nil {
		messageCollection.AddError(err.Error(), "name")
	}

	c.Email = r.FormValue("email")
	err = validateEmail(c.Email)
	if err != nil {
		messageCollection.AddError(err.Error(), "email")
	}

	age, _ := strconv.ParseUint(r.FormValue("age"), 10, 64)
	if err != nil {
		messageCollection.AddError("invalid age", "age")
	}
	c.Age = age
	err = validateAge(c.Age)
	if err != nil {
		messageCollection.AddError(err.Error(), "age")
	}

	c.Gender = r.FormValue("gender")
	err = validateGender(c.Gender)
	if err != nil {
		messageCollection.AddError(err.Error(), "gender")
	}

	if messageCollection.GetCount() > 0 {
		return nil, messageCollection.Get()
	}

	storage := persistence.Connect()
	defer storage.Close()

	query := "INSERT INTO customers (Name, Email, Age, Gender) VALUES (?, ?, ?, ?)"
	stmt, _ := storage.Prepare(query)
	res, err := stmt.Exec(c.Name, c.Email, c.Age, c.Gender)
	if err != nil {
		messageCollection.AddError("error when try to insert... Sorry", "system")
		return nil, messageCollection.Get()
	}

	lastId, _ := res.LastInsertId()
	c.Id = lastId

	return c, messageCollection.Get()
}
