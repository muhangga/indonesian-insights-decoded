package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	nameParam := r.URL.Query().Get("name")

	if nameParam == "" {
		nameParam = "World"
	}

	greeting := Greeting{
		Message: fmt.Sprintf("Hello, %s!", nameParam),
	}

	response, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
