package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "bankapp/backend/internal/handlers"
    "bankapp/backend/internal/middleware"

    "github.com/gorilla/mux"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r := mux.NewRouter()

    // Применяем CORS middleware
    r.Use(middleware.CORS)

    // API routes
    api := r.PathPrefix("/api").Subrouter()
    handlers.RegisterRoutes(api)

    // Health check
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })

    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}