package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "hello world",
	}
	json, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("the json is valid!.")
	}
	w.Header().Set("content-type", "application/json")
	w.Write(json)
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("starting server at port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Printf("Error: %s", err)
		panic(err)
	}
}
