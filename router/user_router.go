package router

import (
  "net/http"

  "meeting_recorders/controller/web"
  "meeting_recorders/service"
  "meeting_recorders/tool"
)

func RegisterUserRoutes(mux *http.ServeMux, svc *service.UserService) {
  user := web.NewUserController(svc)
  //mux.HandleFunc("/meeting/test", tool.Method(http.MethodGet, meeting.Test))
  //mux.HandleFunc("/meeting/findAll", tool.Method(http.MethodGet, meeting.FindAll))
  //mux.HandleFunc("/meeting/findByMeetingID", tool.Method(http.MethodGet, meeting.FindByMeetingID))
 // mux.HandleFunc("/meeting/findByUserID", tool.Method(http.MethodGet, meeting.FindByUserID))
  mux.HandleFunc("/user/login", tool.Method(http.MethodPost, user.Login))
  mux.HandleFunc("/user/register", tool.Method(http.MethodPost, user.Register))
}
