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
 connStr := "user=postgres password=Lexso246357090 dbname=mydb sslmode=disable"
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
 http.HandleFunc("/", healthCheck)                // Главный маршрут для проверки
 http.HandleFunc("/create", CreateHandler)        // Роут для создания записи

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

// Обработчик для формы
func CreateHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method == "POST" {
  // Парсим форму
  err := r.ParseForm()
  if err != nil {
   log.Println(err)
   http.Error(w, "Error parsing form", http.StatusBadRequest)
   return
  }

  // Получаем значения из формы
  email := r.FormValue("email")
  password := r.FormValue("password")

  // Вставляем в базу данных
  _, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)
  if err != nil {
   log.Println(err)
   http.Error(w, "Error inserting into database", http.StatusInternalServerError)
   return
  }

  // После успешного выполнения перенаправляем на главную страницу
  http.Redirect(w, r, "/", http.StatusMovedPermanently)
 } else {
  // Если это не POST запрос, отправляем HTML-страницу с формой
  http.ServeFile(w, r, "index.html")
 }
}
