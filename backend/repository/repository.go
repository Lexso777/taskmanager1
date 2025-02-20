package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Task struct {
	Email      string `json:"email"`
	TitleTask  string `json:"title_task"`
	TextTask   string `json:"text_task"`
	StatusTask int    `json:"status_task"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTasks вызван")

	var request struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		http.Error(w, "Некорректное тело запроса", http.StatusBadRequest)
		return
	}

	log.Println("Email из запроса:", request.Email)

	rows, err := DB.Query(SQLGetTasks, request.Email)
	if err != nil {
		log.Println("Ошибка выполнения запроса к БД:", err)
		http.Error(w, "Ошибка запроса к базе данных", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Email, &task.TitleTask, &task.TextTask, &task.StatusTask); err != nil {
			log.Println("Ошибка чтения строки:", err)
			http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка итерации строк:", err)
		http.Error(w, "Ошибка обработки результата", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println("Ошибка кодирования ответа:", err)
		http.Error(w, "Ошибка формирования ответа", http.StatusInternalServerError)
		return
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	log.Println("AddTask вызван")

	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		http.Error(w, "Некорректное тело запроса", http.StatusBadRequest)
		return
	}
	log.Println("Данные задачи из запроса:", task)

	_, err := DB.Exec(SQLCreateTask, task.Email, task.TitleTask, task.TextTask, task.StatusTask)
	if err != nil {
		log.Println("Ошибка при добавлении задачи:", err)
		http.Error(w, "Ошибка при добавлении задачи", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Задача успешно добавлена"))
}
