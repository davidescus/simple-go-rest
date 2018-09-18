package routes

import (
	"net/http"
	"simple-go-rest/customer"
	"encoding/json"
	"log"
	)

func HandleCustomerRoutes() {
	// Get all customers
	http.HandleFunc("/customer/all", getAllCustomers)

	// Store new customer
	http.HandleFunc("/customer/store", storeCustomer)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	records := customer.GetAll()
	sendResponse(w, records)
}

type response struct {
	HasError bool        `json:"hasError"`
	Message  interface{} `json:"message"`
	Data     interface{} `json:"data"`
}

func storeCustomer(w http.ResponseWriter, r *http.Request) {
	response := response{}
	c := customer.Customer{}
	data, messageSlice := c.Store(r)
	if len(messageSlice) > 0 {
		response.HasError = true
	}

	response.Message = messageSlice
	response.Data = data
    sendResponse(w, response)
}

func sendResponse(w http.ResponseWriter, data interface{}) {
	jsonData, encodeError := json.Marshal(data)
	if encodeError != nil {
		log.Fatal(encodeError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
}
