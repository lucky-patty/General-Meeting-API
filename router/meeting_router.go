package router

import (
  "net/http"

  "meeting_recorders/controller/web"
)

func RegisterMeetingRoutes(mux *http.ServeMux) {
  meeting := web.NewMeetingController()

  mux.HandleFunc("/meeting/test", meeting.Test)
}
