package middle

import (
	"backend/repository"
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string `json:"message"`
}

var DB *sql.DB

type AuthRequest struct {
	Email    string `json:"'email'"`
	Password string `json:"password"`
}

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

	body, _ := io.ReadAll(r.Body)
	decodedBody := string(body)
	log.Printf("RAW JSON BODY (decoded): %s", decodedBody)

	r.Body = io.NopCloser(bytes.NewBuffer([]byte(decodedBody)))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	_, err = DB.Exec(repository.SQLCreateUsers, data.Email, string(hashedPassword))
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}
	log.Printf("Email перед вставкой: %s", data.Email)
	log.Printf("Пароль перед вставкой (hash): %s", string(hashedPassword))

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
	err := DB.QueryRow(repository.SQLGetPassword, authReq.Email).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "Invalid email", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(authReq.Password)); err != nil {
		log.Println(hashedPassword)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"email":   authReq.Email,
	})

}
