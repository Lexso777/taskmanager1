package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type Response struct {
	Message string `json:"message"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	connStr := "user=postgres password=lfybkf2412 dbname=mydb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/login", LoginHandler)

	handlerWithCors := c.Handler(http.DefaultServeMux)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Server is running and connected to the database!"})
}

// Обработчик для формы
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Хеширование пароля перед сохранением в базу данных
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", data.Email, string(hashedPassword))
	if err != nil {
		http.Error(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Record created successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var authReq AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email=$1", authReq.Email).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "Invalid email ", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(authReq.Password)); err != nil {
		log.Println(hashedPassword)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Login successful"})
}
