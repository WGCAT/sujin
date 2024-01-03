package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var users = map[string]*User{}

type User struct {
	ID        string    `json:"ID"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {

	server := http.NewServeMux()
	server.HandleFunc("/foo", fooHandler)

	err := http.ListenAndServe(":8081", server)
	if err != nil {
		fmt.Println(err)
	}

}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			json.NewEncoder(w).Encode(users)
		}
	case "POST":
		{
			var user User

			json.NewDecoder(r.Body).Decode(&user)
			users[user.ID] = &user

		}
	case "PUT":
		{
			resp := "PUT디지몬"
			w.Write([]byte(resp))
		}
	case "DELETE":
		{
			delete(users, "ID")
		}
	}
}

// func GetFooHandler(w *http.ResponseWriter) {

// }
