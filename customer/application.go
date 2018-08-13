package customer

import (
	"strconv"
	"net/http"
	"simple-go-rest/persistence"
	"log"
	"errors"
)

// Customer ...
type Customer struct {
	Id     int64
	Name   string
	Email  string
	Age    uint64
	Gender string
}

// Store ...
func (c *Customer) Store(r *http.Request) (*Customer, error) {
	r.ParseForm()
	var message string

	c.Name = r.FormValue("name")
	message += ValidateName(c.Name)

	c.Email = r.FormValue("email")
	message += ValidateEmail(c.Email)

	age, err := strconv.ParseUint(r.FormValue("age"), 10, 64)
	if err != nil {
		message += "Please provide a valid age"
	}
	c.Age = age
	message += ValidateAge(c.Age)

	c.Gender = r.FormValue("gender")
	message += ValidateGender(c.Gender)

	if message != "" {
		return nil, errors.New(message)
	}

	storage := persistence.Connect()
	defer storage.Close()

	query := "INSERT INTO customers (Name, Email, Age, Gender) VALUES (?, ?, ?, ?)"
	stmt, _ := storage.Prepare(query)
	res, err := stmt.Exec(c.Name, c.Email, c.Age, c.Gender)
	if err != nil {
		message += "Error when try to insert... Sorry!"
		return nil, errors.New(message)
	}

	lastId, _ := res.LastInsertId()
	c.Id = lastId

	return c, nil
}

type Customers []Customer

// GetAll ...
func GetAll() []Customer {
	storage := persistence.Connect()
	defer storage.Close()

	query := "SELECT * FROM customers"
	rows, _ := storage.Query(query)

	customers := make([]Customer, 0)

	for rows.Next() {
		c := &Customer{}
		err :=rows.Scan(&c.Id, &c.Name, &c.Email, &c.Age, &c.Gender)
		if err != nil {
			log.Fatal(err)
		}
		customers = append(customers, *c)
	}

	return customers
}
