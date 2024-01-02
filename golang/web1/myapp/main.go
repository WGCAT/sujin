package main

import (
	"fmt"
	"net/http"
)

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
			GetFooHandler(&w)
		}
	case "POST":
		{
			resp := "POST디지몬"
			w.Write([]byte(resp))
		}
	case "PUT":
		{
			resp := "PUT디지몬"
			w.Write([]byte(resp))
		}
	case "DELETE":
		{
			resp := "DELETE디지몬"
			w.Write([]byte(resp))
		}
	}
}

func GetFooHandler(w *http.ResponseWriter) {
	resp := "안녕디지몬"
	(*w).Write([]byte(resp))
}
