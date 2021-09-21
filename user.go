package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from user get")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Hello, from user get")
}
