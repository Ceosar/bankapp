package handlers

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

// Структура для данных пользователя
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/users", GetUsers).Methods("GET")
    r.HandleFunc("/users", CreateUser).Methods("POST")
    r.HandleFunc("/message", GetMessage).Methods("GET")
}

// GET /api/users - получить список пользователей
func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    users := []User{
        {ID: 1, Name: "Иван Иванов", Email: "ivan@example.com", CreatedAt: time.Now()},
        {ID: 2, Name: "Мария Петрова", Email: "maria@example.com", CreatedAt: time.Now()},
    }
    
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "data":    users,
        "count":   len(users),
    })
}

// POST /api/users - создать пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var newUser User
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "success": "false",
            "error":   "Invalid request body",
        })
        return
    }
    
    // Симулируем создание пользователя
    newUser.ID = 3
    newUser.CreatedAt = time.Now()
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "data":    newUser,
        "message": "User created successfully",
    })
}

// GET /api/message - простой пример получения сообщения
func GetMessage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "message": "Привет с сервера!",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}