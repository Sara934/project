// handlers.go
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Entity struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func createEntity(w http.ResponseWriter, r *http.Request) {
	var entity Entity
	_ = json.NewDecoder(r.Body).Decode(&entity)

	result, err := db.Exec("INSERT INTO entities (name, description) VALUES (?, ?)", entity.Name, entity.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity.ID, _ = result.LastInsertId()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity)
}

func readEntity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var entity Entity
	err := db.QueryRow("SELECT id, name, description FROM entities WHERE id = ?", id).Scan(&entity.ID, &entity.Name, &entity.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity)
}

func updateEntity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var entity Entity
	_ = json.NewDecoder(r.Body).Decode(&entity)

	_, err := db.Exec("UPDATE entities SET name = ?, description = ? WHERE id = ?", entity.Name, entity.Description, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteEntity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := db.Exec("DELETE FROM entities WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
