package main

import (
	"log"
	"net/http"
)

func callBack(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/callback" {
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	log.Println("code: ", code)
	log.Println("state: ", state)
	log.Println("Method type: ", r.Method)
	return
}

func main() {
	http.HandleFunc("/callback", callBack)

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	return
}
