package type

import "time"

type User struct {
  ID          int   `json:"user_id"`
  Username    string `json:"username"`
  CreateDate  time.Timestamp `json:"create_date"`
}
