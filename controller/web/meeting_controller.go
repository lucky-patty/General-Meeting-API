package web

import (
  "fmt"
  "net/http"
  "encoding/json"

  "os"
  "io"


  "meeting_recorders/types"
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

func (c *MeetingController) FindByMeetingID(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")
  
  if id == "" {
    http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
    return
  }
  
  ctx := r.Context()
  res, err := c.Service.FindByMeetingID(ctx, id)
  if err != nil {
    http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(res)
}

func (c *MeetingController) UploadMeeting (w http.ResponseWriter, r *http.Request) {
//  ctx := r.Context()
  err := r.ParseMultipartForm(50 << 20)
  if err != nil {
    http.Error(w, fmt.Sprintf("Failed to parse form: %v", err), http.StatusBadRequest)
    return
  }

  insertReq := types.MeetingRequest{
    UserID: r.FormValue("user_id"),
    MeetingID: r.FormValue("meeting_id"),
    MeetingTitle: r.FormValue("meeting_title"),
  }
 
  if insertReq.MeetingID == "" || insertReq.UserID == "" || insertReq.MeetingTitle == "" {
    http.Error(w, "Missing meeting_id, user_id or meeting_title", http.StatusBadRequest)
    return
  }

  // Extract file 
  file, handler, err := r.FormFile("file")
  if err != nil {
    http.Error(w, "Failed to get file", http.StatusBadRequest)
    return
  }
  defer file.Close()

  targetDir := fmt.Sprintf("record/%s", insertReq.MeetingID)
  if err := os.MkdirAll(targetDir, 0755); err != nil {
    http.Error(w, "Failed to create directory", http.StatusInternalServerError)
    return
  }

  targetPath := fmt.Sprintf("%s/original.mp3", targetDir)
  dst, err := os.Create(targetPath)
  if err != nil {
    http.Error(w, "Failed to create file", http.StatusInternalServerError)
    return
  }
  defer dst.Close()

  if _, err := io.Copy(dst, file); err != nil {
    http.Error(w, "Failed to save file", http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
  fmt.Printf("Uploaded %s for meeting %s \n", handler.Filename, insertReq.MeetingID)
  //  ctx := r.Context()
  // err := c.Service.InsertTranscript(ctx, transcript)
}
