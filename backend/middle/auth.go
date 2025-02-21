package middle

import (
	"backend/repository"
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string `json:"message"`
}

var DB *sql.DB

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var jwtKey = []byte("sectet_key")

type Claims struct {
	UserID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
		http.Error(w, "error create user", http.StatusInternalServerError)
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
	var userID int
	err := DB.QueryRow(repository.SQLGetPassword, authReq.Email).Scan(&userID, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email", http.StatusUnauthorized)
		} else {
			http.Error(w, "Invalid email", http.StatusUnauthorized)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(authReq.Password)); err != nil {
		log.Println(hashedPassword)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	token, err := GenerateToken(userID)
	if err != nil {
		http.Error(w, "Failed generate Token", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   token,
	})

}
