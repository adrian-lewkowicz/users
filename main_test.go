package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"users/server/database"
	"users/server/router"
	"users/server/users"
	"users/server/users/models"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	db := database.InitMockDatabase()
	router := router.NewRouter()
	users.InitDatabase(db)
	router.Handle("/user", users.UserHandler)
	router.Handle("/user/{id}", users.UserHandler)

	t.Run("test post ok", func(t *testing.T) {
		w := httptest.NewRecorder()
		user := models.User{
			Name: "John",
			Age:  33,
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			t.Fail()
			return
		}

		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("test get by id", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/user/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("test put ok", func(t *testing.T) {

		w := httptest.NewRecorder()
		user := models.User{
			Name: "John",
			Age:  33,
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			t.Fail()
			return
		}
		req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(jsonData))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("test put not found", func(t *testing.T) {

		w := httptest.NewRecorder()
		user := models.User{
			Name: "John",
			Age:  33,
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			t.Fail()
			return
		}
		req, _ := http.NewRequest("PUT", "/user/0", bytes.NewBuffer(jsonData))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("test DELETE", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/user/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

}
