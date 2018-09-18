package main

import (
	"net/http"
	"fmt"
	"log"
	"simple-go-rest/routes"
	"strconv"
)

const serverPort = 8282

func main() {
	fmt.Printf("WebServer start on: %d port\n", serverPort)

	routes.HandleCustomerRoutes()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(serverPort), nil))
}
