package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const serverPort = 8282

type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []User

	user1 := User{"John doe", "john.doe@Email.com", 33, "male"}
	user2 := User{"Hilary clinton", "hilary.clinton@Email.com", 41, "female"}
	user3 := User{"Jessica Jameson", "jessica.jameson@Email.com", 19, "female"}
	user4 := User{"Derek Bananas", "derek.bananas@Email.com", 52, "male"}

	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}

func main() {
	fmt.Printf("WebServer start on: %d port\n", serverPort)

	http.HandleFunc("/user/all", getAllUsers)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(serverPort), nil))
}
