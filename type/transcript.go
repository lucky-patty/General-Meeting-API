package type

import "time"

type Transcript struct {
  MeetingID  int            `json:"meeting_id"` 
  UserID     int            `json:"user_id"`
  Transcript string         `json:"transcript"`
  CreateDate time.Timestamp `json:"create_date"`
}
