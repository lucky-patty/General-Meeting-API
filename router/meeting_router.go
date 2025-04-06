package router

import (
  "net/http"

  "meeting_recorders/controller/web"
  "meeting_recorders/service"
)

func RegisterMeetingRoutes(mux *http.ServeMux, svc *service.MeetingService) {
  meeting := web.NewMeetingController(svc)

  mux.HandleFunc("/meeting/test", meeting.Test)
  mux.HandleFunc("/meeting/findAll", meeting.FindAll)

}
