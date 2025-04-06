package web

import (
  "fmt"
  "net/http"
  "encoding/json"
  //"meeting_recorders/types"
  "meeting_recorders/service"
)

type MeetingController struct {
  Service *service.MeetingService
}

func NewMeetingController(s *service.MeetingService) *MeetingController{
  return &MeetingController{
    Service: s,
  }
}

func (c *MeetingController) Test(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Hello World")
}

func (c *MeetingController) FindAll(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
  res, err := c.Service.FindAll(ctx)
  if err != nil {
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}

func (c *MeetingController) FindByMeetingID(w http.ResponseWriter, r *http.Request) {}

func (c *MeetingController) FindByUserID(w http.ResponseWriter, r *http.Request) {}

func (c *MeetingController) Insert(w http.ResponseWriter, r *http.Request) {
//  ctx := r.Context()
  // err := c.Service.InsertTranscript(ctx, transcript)
}
