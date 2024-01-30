package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    database, err := sql.Open("sqlite3", "./data/mydatabase.db")
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    // needs error handling added
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
    statement.Exec()

    // needs error handling added
    statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
    statement.Exec("John", "Doe")

    // needs error handling added
    rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
    var id int
    var firstname, lastname string
    for rows.Next() {
        rows.Scan(&id, &firstname, &lastname)
        log.Println(id, firstname, lastname)
    }
}