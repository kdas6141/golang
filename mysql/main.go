package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial");

	//Open database connection
	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
    fmt.Println("Successfully connected to the database");

    testdb, err := db.Query("USE testdb")

    if err != nil {
        panic(err.Error())
    }
    defer testdb.Close()

    fmt.Println("Successfully switch to the database");

	insert, err := db.Query("INSERT INTO users VALUES('Soumik Das')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted into user tables");
}

