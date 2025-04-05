package router 

import (
  "net/http"

  "meeting_recorders/middleware"
)

func NewRouter() http.Handler {
  mux := http.NewServeMux()
  // RegisterTransactionRoutes(mux)
  // RegisterTranscripRoutes(mux)

  RegisterUserRoutes(mux)

  return middleware.Chain(
      mux,
      middleware.RequestID,
      middleware.CORS,
  )
}

