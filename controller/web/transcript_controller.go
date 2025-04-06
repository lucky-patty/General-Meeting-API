package web

import (
  "net/http"
  "meeting_recorders/service"
)

type TranscriptController struct {
  Service *service.TranscriptService
}

func NewTranscriptController(s *service.TranscriptService) *TranscriptController {
  return &TranscriptController{
    Service: s,
  }
}

func (c *TranscriptController) Insert(w http.ResponseWriter, r *http.Request) {
 // ctx := r.Context()
  // err := c.Service.InsertTranscript(ctx, transcript)
}
