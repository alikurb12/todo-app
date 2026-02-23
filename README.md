# 📝 Todo App — REST API in Go

A simple and clean REST API for task management, built with Go and PostgreSQL.

## 🛠 Tech Stack

- **Go** — programming language
- **Chi** — lightweight HTTP router
- **pgx / pgxpool** — PostgreSQL driver with connection pooling
- **godotenv** — loading configuration from `.env` file

---

## 📁 Project Structure

```
todo-app-go/
├── cmd/
│   └── main.go                  # Entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration loading
│   ├── handler/
│   │   └── task_handler.go      # HTTP handlers
│   ├── service/
│   │   └── task_service.go      # Business logic
│   ├── repository/
│   │   └── task_repository.go   # Database layer
│   └── model/
│       └── task.go              # Data models
├── .env                         # Environment variables
├── go.mod
└── go.sum
```

---

## ⚙️ Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/alikurb12/todo-app-go.git
cd todo-app-go
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Create a `.env` file in the project root

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=your_password
DB_NAME=todo
SSL_MODE=disable
```

### 4. Create the database table

```sql
CREATE TABLE tasks (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);
```

### 5. Run the application

```bash
go run cmd/main.go
```

Server will start at `http://localhost:8080`

---

## 🌐 API Reference

### Get all tasks

```
GET /tasks
```

**Response:**
```json
[
  {
    "id": 1,
    "title": "Buy milk",
    "description": "From the store around the corner",
    "completed": false,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
]
```

---

### Get task by ID

```
GET /tasks/{id}
```

**Response:**
```json
{
  "id": 1,
  "title": "Buy milk",
  "description": "From the store around the corner",
  "completed": false,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

---

### Create a task

```
POST /tasks
```

**Request body:**
```json
{
  "title": "New task",
  "description": "Task description",
  "completed": false
}
```

**Response:** `201 Created`

---

### Update a task

```
PUT /tasks/{id}
```

**Request body:**
```json
{
  "title": "Updated title",
  "description": "Updated description",
  "completed": true
}
```

**Response:** `200 OK`

---

### Delete a task

```
DELETE /tasks/{id}
```

**Response:** `204 No Content`

---

## 🏗 Architecture

The project follows a clean three-layer architecture:

```
HTTP Request
    ↓
Handler     — receives request, sends response
    ↓
Service     — business logic & validation
    ↓
Repository  — database operations
    ↓
PostgreSQL
```
