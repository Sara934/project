// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEntity(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"TestEntity","description":"TestDescription"}`)

	req, _ := http.NewRequest("POST", "/entities", bytes.NewBuffer(payload))
	response := executeRequest(req)

	assert.Equal(t, http.StatusOK, response.Code)

	var entity Entity
	json.Unmarshal(response.Body.Bytes(), &entity)

	assert.NotEqual(t, 0, entity.ID)
	assert.Equal(t, "TestEntity", entity.Name)
	assert.Equal(t, "TestDescription", entity.Description)
}

// Add similar tests for other CRUD operations

func clearTable() {
	db.Exec("DELETE FROM entities")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mainHandler := http.HandlerFunc(main)
	mainHandler.ServeHTTP(rr, req)
	return rr
}
