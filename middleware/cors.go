package middleware

import (
  "net/http"
)

var allowedOrigins = map[string]bool{
  "http://localhost:4321": true,
  "http://192.168.2.41:4321": true,
}

func CORS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    origin := r.Header.Get("Origin")    
    if allowedOrigins[origin] {
            w.Header().Set("Access-Control-Allow-Origin", origin)
            w.Header().Set("Access-Control-Allow-Credentials", "true")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    }
    if r.Method == http.MethodOptions {
          w.WriteHeader(http.StatusNoContent)
          return
    }

    next.ServeHTTP(w, r)
  })
}
