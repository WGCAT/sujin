package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

/*
[구현해볼것]
1. GET함수로 모든 User정보 Response
2. GET함수로 ID입력해서 해당ID값을 가진 User정보 Response
3. POST함수로 User정보 users map에 저장 후 Response
4. PUT함수로 User정보 전체 수정 후 Response
5. PUT함수로 User정보 일부 수정 후 Response
6. DELETE함수로 ID입력해서 해당ID값을 가진 User정보 삭제 후 Response
7. DELETE함수로 다수의 ID입력해서 해당ID값을 가진 User정보들 삭제 후 Response
*/

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
	server.HandleFunc("/foo/{ID}", GetFooHandler)

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
		{	pathParams := strings.Split(r.URL.Path, "/")
			userId := pathParams[len(pathParams)-1]
		 
			delete(users, "ID")
		}
	}
}

func GetFooHandler(w *http.ResponseWriter, r *http.Request) {
	vars := server.Vars(r)
	ID, _ := strconv.Atoi(vars{"ID"})
	for _, user := range users {
		if user.ID == ID {
			json.NewEncoder(&w).Encode(user)
			return
		}
	}

}
