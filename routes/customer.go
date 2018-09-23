package routes

import (
	"net/http"
	"encoding/json"
	"log"
	"simple-go-rest/customer"
)

func HandleCustomerRoutes() {
	// Get all customers
	http.HandleFunc("/customer/all", getAll)

	// Store new customer
	http.HandleFunc("/customer/store", store)
}

type response struct {
	HasError bool        `json:"hasError"`
	Message  interface{} `json:"message"`
	Data     interface{} `json:"data"`
}

func getAll(w http.ResponseWriter, r *http.Request) {
	response := response{}
	records, messageSlice := customer.GetAll()

	//if len(messageSlice) > 0 {
	//	response.HasError = true
	//}

	response.Message = messageSlice
	response.Data = records
	sendResponse(w, response)
}

func store(w http.ResponseWriter, r *http.Request) {
	response := response{}
	data, messageSlice := customer.Store(r)

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
