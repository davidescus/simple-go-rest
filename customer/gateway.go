package customer

import (
	"errors"
	"simple-go-rest/persistence"
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
func (c *Customer) Store() (error) {
	storage := persistence.Connect()
	defer storage.Close()

	query := "INSERT INTO customers (Name, Email, Age, Gender) VALUES (?, ?, ?, ?)"
	stmt, err := storage.Prepare(query)
	if err != nil {
		return errors.New("error when try to insert... Sorry")
	}

	res, err := stmt.Exec(c.Name, c.Email, c.Age, c.Gender)
	if err != nil {
		return errors.New("error when try to insert... Sorry")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return errors.New("error when try to insert... Sorry")
	}

	c.Id = lastId

	return nil
}

// Customers ...
type Customers []Customer

// GetAllCustomers ...
func (c *Customers) All() (*Customers, error) {

	storage := persistence.Connect()
	defer storage.Close()

	query := "SELECT * FROM customers"
	rows, err := storage.Query(query)
	if err != nil {
		return c, errors.New("something went wrong")
	}

	for rows.Next() {
		cst := Customer{}
		err := rows.Scan(&cst.Id, &cst.Name, &cst.Email, &cst.Age, &cst.Gender)
		if err != nil {
			return c, errors.New("something went wrong")
		}
		*c = append(*c, cst)
	}

	return c, nil
}