package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"users/server/database"
	"users/server/users/models"
)

var db database.Database

func InitDatabase(database database.Database) {
	db = database
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserHandler(w, r)
	case http.MethodPost:
		postUserHandler(w, r)
	case http.MethodPut:
		putUserHandler(w, r)
	case http.MethodDelete:
		deleteUserHandler(w, r)
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/user/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var wg sync.WaitGroup

	wg.Add(1)

	userChan := make(chan models.User)
	errorChan := make(chan error)

	go func() {
		defer wg.Done()
		user, err := db.GetUserById(id)
		if err != nil {
			errorChan <- err
			return
		}
		userChan <- user
	}()

	go func() {
		wg.Wait()
		close(userChan)
		close(errorChan)
	}()

	var user models.User

	select {
	case u := <-userChan:
		user = u
	case err := <-errorChan:
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		log.Println("Error fetching ", err)
		return
	}

	duration := time.Since(startTime)
	log.Printf("Request took %s\n", duration)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	userChan := make(chan models.User)
	errorChan := make(chan error)

	go func() {
		defer wg.Done()

		userId, err := db.PostUser(user)
		if err != nil {
			errorChan <- err
			return
		}
		user.ID = userId
		userChan <- user
	}()

	go func() {
		wg.Wait()
		close(userChan)
		close(errorChan)
	}()

	select {
	case u := <-userChan:
		user = u
	case err := <-errorChan:
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		log.Println("Error inserting user ", err)
		return
	}

	duration := time.Since(startTime)
	log.Printf("Request took %s\n", duration)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func putUserHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/user/")
	id, err := strconv.Atoi(idStr)

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	user.ID = id
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	userChan := make(chan bool)
	errorChan := make(chan error)

	go func() {
		defer wg.Done()

		err := db.PutUser(user)
		if err != nil {
			errorChan <- err
			return
		} else {
			userChan <- true
		}
	}()

	go func() {
		wg.Wait()
		close(userChan)
		close(errorChan)
	}()

	var puted bool
	select {
	case u := <-userChan:
		puted = u
	case err := <-errorChan:
		http.Error(w, "Error inserting user", http.StatusNotFound)
		log.Println("Error inserting user ", err)
		return
	}

	duration := time.Since(startTime)
	log.Println("Update completed - ", puted)
	log.Printf("Request took %s\n", duration)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/user/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	userChan := make(chan bool)
	errorChan := make(chan error)

	go func() {
		defer wg.Done()
		err := db.DeleteUser(id)
		if err != nil {
			errorChan <- err
			return
		} else {
			userChan <- true
		}
	}()

	go func() {
		wg.Wait()
		close(userChan)
		close(errorChan)
	}()
	var deleted bool
	select {
	case u := <-userChan:
		deleted = u
	case err := <-errorChan:
		http.Error(w, "Error deleting user", http.StatusNotFound)
		log.Println("Error deleting ", err)
		return
	}
	duration := time.Since(startTime)
	log.Printf("Request took %s\n", duration)

	w.WriteHeader(http.StatusNoContent)
	log.Println("Remove completed - ", deleted)
}
