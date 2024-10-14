package main

import (
	"fmt"
	"net/http"
	"users/server/database"
	"users/server/router"
	"users/server/users"
)

func main() {
	router := router.NewRouter()
	db := database.InitDatabase("host=localhost user=gorm password=gorm dbname=users port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	users.InitDatabase(db)
	router.Handle("/user", users.UserHandler)
	router.Handle("/user/{id}", users.UserHandler)
	fmt.Println("Server run on port :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
