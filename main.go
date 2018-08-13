package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"simple-go-rest/customer"
)

const serverPort = 8282

type response struct {
	HasError bool        `json:"hasError"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	records := customer.GetAll()
	data, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}

func storeUser(w http.ResponseWriter, r *http.Request) {
	response := response{}
    c := customer.Customer{}
    data, err := c.Store(r)
	if err != nil {
		response.HasError = true
		response.Message = err.Error()
	}

    response.Data = data

	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
}

func main() {
	fmt.Printf("WebServer start on: %d port\n", serverPort)

	// Get all customers
	http.HandleFunc("/customer/all", getAllCustomers)

	// Store new customer
	http.HandleFunc("/customer/store", storeUser)

	// Start server and log fatal err
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(serverPort), nil))
}
