// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS entities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()

	r := mux.NewRouter()

	r.HandleFunc("/entities", createEntity).Methods("POST")
	r.HandleFunc("/entities/{id}", readEntity).Methods("GET")
	r.HandleFunc("/entities/{id}", updateEntity).Methods("PUT")
	r.HandleFunc("/entities/{id}", deleteEntity).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
