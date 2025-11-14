package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/users", GetUsers).Methods("GET")
    r.HandleFunc("/users", CreateUser).Methods("POST")
    // Добавьте свои роуты здесь
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "users": []string{"user1", "user2"},
    })
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User created",
    })
}