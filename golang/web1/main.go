package main

import (
	myapp "command-line-argumentsC:\\GO_LAB\\sujin\\golang\\web1\\myapp\\app.go"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
