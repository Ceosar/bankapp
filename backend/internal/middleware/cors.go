package middleware

import (
    "net/http"

    "github.com/rs/cors"
)

var CORS = cors.New(cors.Options{
    AllowedOrigins: []string{
        "http://localhost:8081",  // Expo dev server
        "http://localhost:19006", // Expo web
        "exp://localhost:8081",   // Expo mobile
    },
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    AllowCredentials: true,
}).Handler