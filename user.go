package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const dsn = "host=localhost user=postgres password=postgres dbname=employee_mgmnt port=5432 sslmode=disable TimeZone=Asia/Shanghai"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func InitializeMigration() {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
	}
	db.AutoMigrate(&User{})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from user get")
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from user get")
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)

	json.NewEncoder(w).Encode(&user)
}
