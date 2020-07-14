package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
        fmt.Println(err.Error())
        panic("Failed to connect to database")
    }
    defer db.Close()

	var users[]User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
	fmt.Println("All users endpoint hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
    db, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("Failed to connect to database")
    }
    defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
	fmt.Println("New user endpoint hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    db, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("Failed to connect to database")
    }
    defer db.Close()

	vars := mux.Vars(r)
    name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

    fmt.Fprintf(w, "User Successfully Deleted")
    fmt.Println("Delete user endpoint hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
   db, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("Failed to connect to database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    name := vars["name"]
    email := vars["email"]

    var user User
    db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)
    fmt.Fprintf(w, "Successfully Updated User")
    fmt.Println("Update user endpoint hit")
}

