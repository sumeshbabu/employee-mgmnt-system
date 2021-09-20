package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	http.ListenAndServe(":9000", router)
}
func main() {
	InitializeMigration()
	initializeRouter()
	fmt.Println("Hello, playground")
}
