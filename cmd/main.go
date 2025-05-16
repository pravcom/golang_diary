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

	storage.DB.AutoMigrate(&models.Diary{})

	http.HandleFunc("/diary/create", handlers.SaveDiary(storage))
	http.HandleFunc("/diary/get", handlers.GetDiary(storage))
	http.HandleFunc("/diary/delete", handlers.DeleteDiary(storage))

	fmt.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
