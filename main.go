package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")

	http.ListenAndServe(":9000", router)
}
func main() {
	initializeRouter()
	fmt.Println("Hello, playground")
}
