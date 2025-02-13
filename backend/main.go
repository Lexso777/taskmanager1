package main

import (
 "database/sql"
 "fmt"
 "log"
 "net/http"
 "github.com/rs/cors"
 _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
 // Строка подключения к базе данных
//  connStr := "user=myuser password=password dbname=taskmanager2 sslmode=disable"
  connStr := "user=postgres password=Lexso246357090 dbname=my_database sslmode=disable"
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