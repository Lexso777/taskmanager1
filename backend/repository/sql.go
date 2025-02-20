package repository

const (
	SQLCreateScheduler = `
	CREATE TABLE users (
	id		SERIAL PRIMARY KEY,
	email	VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	);
	`
	SQLCreateTableTasks = `	
	CREATE TABLE tasks (
	id			SERIAL,
	email 		VARCHAR(255) NOT NULL PRIMARY KEY,
	title_task	VARCHAR(255) NOT NULL,
	text_task	TEXT NOT NULL,
	status_task INT NOT NULL,
	) 
	`
	SQLCreateUsers = `
	INSERT INTO users(email, password) VALUES ($1,$2)
	`
	SQLGetPassword = `
	SELECT password FROM users WHERE email=$1
	`

	SQLCreateTask = `
	INSERT INTO tasks (email, title_task, text_task, status_task) VALUES ($1, $2, $3, $4)
	`
	SQLGetTasks = ` 
	SELECT
		email,
		title_task,
		text_task,
		status_task
		FROM tasks
		WHERE email =$1`
)
