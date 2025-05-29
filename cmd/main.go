package main

import (
	"fmt"
	"log"
	"net/http"

	"app-diary/internal/handlers"
	"app-diary/internal/models"
	"app-diary/internal/storage/postgresql"
)

const (
	dsn  = "host=localhost user=postgres dbname=postgres password=postgres sslmode=disable port=8090"
	port = "8081"
)

func main() {
	storage := postgresql.New(dsn)

	defer storage.Close()

	storage.DB.AutoMigrate(&models.Diary{}, &models.Users{})

	http.HandleFunc("/diary/create", handlers.JWTMiddleware(handlers.SaveDiary(storage)))
	http.HandleFunc("/diary/get", handlers.GetDiary(storage))
	http.HandleFunc("/diary/delete", handlers.DeleteDiary(storage))
	http.HandleFunc("/diary/createUser", handlers.SaveUser(storage))
	http.HandleFunc("/diary/login", handlers.Login(storage))
	http.HandleFunc("/diary/logout", handlers.Logout())

	fmt.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
