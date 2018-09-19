package customer

import (
		"net/http"
	"simple-go-rest/persistence"
	"log"
	"simple-go-rest/messages"
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
func (c *Customer) Store(r *http.Request) (*Customer, messages.Messages) {
    return storeNewCustomer(r, c)
}

// GetAll ...
func GetAll() ([]Customer, messages.Messages) {
	storage := persistence.Connect()
	defer storage.Close()

	messageCollection := &messages.Messages{}

	query := "SELECT * FROM customers"
	rows, err := storage.Query(query)
	if err != nil {
		messageCollection.AddError("error When try to parse request", "system")
		return nil, messageCollection.Get()
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		c := &Customer{}
		err :=rows.Scan(&c.Id, &c.Name, &c.Email, &c.Age, &c.Gender)
		if err != nil {
			log.Fatal(err)
		}
		customers = append(customers, *c)
	}

	return customers, messageCollection.Get()
}
