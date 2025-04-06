package types 

import "time"

type Meeting struct {
  MeetingID       int             `json:"meeting_id`
  MeetingTitle    string          `json:"meeting_title"`
  UserID          int             `json:"user_id"`
  CreateDate      time.Time  `json:"create_date"`
}

type MeetingNote struct {
  Title   string `json:"title"`
  Content string `json:"content"`
  Summary string `json:"summary"`
  UserID    string `json:"user_id"`
  MeetingID string  `json:"meeting_id"`
  Tags      []string `json:"tags"`
  CreateDate string  `json:"create_date"`
}

type MeetingRequest struct{
  UserID    string 
  MeetingID string
  MeetingTitle string
}
