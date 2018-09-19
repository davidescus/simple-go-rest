package customer

import (
	"simple-go-rest/messages"
	"strconv"
	"net/http"
	"simple-go-rest/persistence"
	)

func storeNewCustomer(r *http.Request, c *Customer) (*Customer, messages.Messages) {
	mess := &messages.Messages{}
	SetMessenger(mess)

	err := r.ParseForm()
	if err != nil {
		mess.AddError("error When try to parse request", "system")
		return nil, mess.Get()
	}

	c.Name = r.FormValue("name")
	validateName(c.Name)

	c.Email = r.FormValue("email")
	validateEmail(c.Email)

	age, _ := strconv.ParseUint(r.FormValue("age"), 10, 64)
	if err != nil {
		mess.AddError("invalid age", "age")
	}
	c.Age = age
	validateAge(c.Age)

	c.Gender = r.FormValue("gender")
	validateGender(c.Gender)

	if mess.GetCount() > 0 {
		return nil, mess.Get()
	}

	storage := persistence.Connect()
	defer storage.Close()

	query := "INSERT INTO customers (Name, Email, Age, Gender) VALUES (?, ?, ?, ?)"
	stmt, _ := storage.Prepare(query)
	res, err := stmt.Exec(c.Name, c.Email, c.Age, c.Gender)
	if err != nil {
		mess.AddError("error when try to insert... Sorry", "system")
		return nil, mess.Get()
	}

	lastId, _ := res.LastInsertId()
	c.Id = lastId

	return c, mess.Get()
}
