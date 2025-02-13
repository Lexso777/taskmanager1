package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	// Строка подключения к базе данных
	//  connStr := "user=myuser password=password dbname=taskmanager2 sslmode=disable"
	connStr := "user=postgres password=lfybkf2412 dbname=mydb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	// Настройка CORS (разрешаем запросы с localhost:3000)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Разрешаем только с этого адреса
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Роуты
	http.HandleFunc("/", healthCheck) // Главный маршрут для проверки

	// Обертываем сервер с CORS
	handlerWithCors := c.Handler(http.DefaultServeMux)

	// Запуск сервера
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))
}

// Обработчик для главной страницы
func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Отправляем ответ, что сервер работает
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running and connected to the database!"))
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		model := r.FormValue("model")
		company := r.FormValue("company")
		price := r.FormValue("price")

		_, err = db.Exec("insert into user (model, company, price) values (?, ?, ?)",
			model, company, price)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "http://localhost:3000/")
	}
}
