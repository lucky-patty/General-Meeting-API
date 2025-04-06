package router 

import (
  "net/http"

  "meeting_recorders/middleware"
  "meeting_recorders/service"
)

func NewRouter(svc *service.Service) http.Handler {
  mux := http.NewServeMux()
  RegisterMeetingRoutes(mux, svc.Meeting)
  //RegisterUserRoutes(mux)

  return middleware.Chain(
      mux,
      middleware.RequestID,
      middleware.CORS,
  )
}

