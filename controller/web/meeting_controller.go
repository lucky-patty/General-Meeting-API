package web

import (
  "fmt"
  "net/http"
  "encoding/json"
  "time"
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

  // Create context if no issue 
  ctx := r.Context()
  // Transcript the audio
  transcribe, errTranscribe := c.Service.Whisper.Transcribe(ctx, targetPath)
  if errTranscribe != nil {
    fmt.Printf("Whisper error: %s \n", errTranscribe)
    http.Error(w, fmt.Sprintf("Whisper failed: ", errTranscribe), http.StatusInternalServerError)
    return
  } 

  fmt.Println("Whisper transcribe: ", transcribe)
  
  var transcript types.WhisperTranscript
  errTranscript := json.Unmarshal([]byte(transcribe), &transcript)
  if errTranscript != nil {
    fmt.Printf("Convert error: %s \n", errTranscript)
    http.Error(w, fmt.Sprintf("Convert from whisper failed: ", errTranscript), http.StatusInternalServerError)
    return
  }

  // Summarize
  summary, err := c.Service.Gpt.Summarize(ctx, transcript.Text)
  if err != nil {
  fmt.Println("GPT Error: ", err)
  http.Error(w, fmt.Sprintf("GPT Error: ", err), http.StatusInternalServerError)
  return
  }
  fmt.Println("Summary: \n", summary)
 
  note := &types.MeetingNote{
    Title:     insertReq.MeetingTitle,
    Content:   transcript.Text,
    Summary:   summary,
    UserID:    insertReq.UserID,
    MeetingID: insertReq.MeetingID,
    CreateDate: time.Now().String(),
  }
  errInsert := c.Service.InsertMeetingNote(ctx, note)
  if errInsert != nil {
    fmt.Println("ES Errror: ", errInsert)
    http.Error(w, fmt.Sprintf("ES Insert error: ", errInsert), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
  fmt.Printf("Uploaded %s for meeting %s \n", handler.Filename, insertReq.MeetingID)
  
}
