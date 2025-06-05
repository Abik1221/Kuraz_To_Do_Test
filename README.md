# 📝 Task Management REST API

A simple RESTful API built using the **Gin Web Framework** and **MySQL**, designed to manage tasks (Create, Read, Update, Delete).

> 🚀 Project developed for **Kuraz Internship Program**

---

## 📌 Features

- ✅ Get all tasks
- ➕ Add new task
- ✅ Mark task as completed
- ❌ Delete task

---

## 📁 API Endpoints

| Method | Endpoint             | Description             |
|--------|----------------------|-------------------------|
| GET    | `/api/tasks/`        | Get all tasks           |
| POST   | `/api/tasks/`        | Create a new task       |
| PUT    | `/api/tasks/:id`     | Mark task as completed  |
| DELETE | `/api/tasks/:id`     | Delete a task           |

---

## 🛠️ Tech Stack

- [Golang](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [MySQL](https://www.mysql.com/)
- [GORM](https://gorm.io/)
- [.env](https://github.com/joho/godotenv) for config

---

## ⚙️ Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/task-api.git
   cd task-api
2.**Create the .env file and set**
     ```bash
        DB_USER=root
        DB_PASS=yourpassword
        DB_HOST=127.0.0.1
        DB_PORT=3306
        DB_NAME=taskdb
3. **Run MySQL server and create the database:**
  ```bash
  CREATE DATABASE taskdb;

4. **Install Go dependencies**

  ```bash
  go mod tidy

5. **Run the API**

 ```bash 
 go run main.go