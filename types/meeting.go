package types 

import "time"

type Meeting struct {
  MeetingID       int             `json:"meeting_id`
  MeetingTitle    string          `json:"meeting_title"`
  UserID          int             `json:"user_id"`
  CreateDate      time.Time  `json:"create_date"`
}
