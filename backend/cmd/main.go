package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"backend/middle"
	"backend/repository"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

type Response struct {
	Message string `json:"message"`
}

func main() {
	var err error
	connStr := "user=postgres password=lfybkf2412 dbname=mydb sslmode=disable"

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()
	log.Println("БД подключена и работает! 🚀")

	middle.DB = db
	repository.DB = db

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/create", middle.CreateHandler)
	http.HandleFunc("/login", middle.LoginHandler)
	http.HandleFunc("/tasks", repository.GetTasks)
	http.HandleFunc("/tasks/addtasks", repository.AddTask)
	http.HandleFunc("/task/updtask", repository.UpdateTask)
	http.HandleFunc("/task/statusupd", repository.UpdateStatus)

	handlerWithCors := c.Handler(http.DefaultServeMux)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Server is running and connected to the database!"})
}
