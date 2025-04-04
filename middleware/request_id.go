package middleware

import (
  "context"
  "net/http"

  "github.com/google/uuid"
)

type ctxKey string 
const RequestIDKey ctxKey = "requestID"


func RequestID(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    reqID := uuid.New().String() 

    ctx := context.WithValue(r.Context(), RequestIDKey, reqID)

    w.Header().Set("X-Request-ID", reqID)

    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

