package router

import (
  "net/http"

  "meeting_recorders/controller/web"
  "meeting_recorders/service"
  "meeting_recorders/tool"
)

func RegisterMeetingRoutes(mux *http.ServeMux, svc *service.MeetingService) {
  meeting := web.NewMeetingController(svc)

  mux.HandleFunc("/meeting/test", tool.Method(http.MethodGet, meeting.Test))
  mux.HandleFunc("/meeting/findAll", tool.Method(http.MethodGet, meeting.FindAll))
  mux.HandleFunc("/meeting/findByMeetingID", tool.Method(http.MethodGet, meeting.FindByMeetingID))
}
