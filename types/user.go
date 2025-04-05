package types

import "time"

type User struct {
  ID          int   `json:"user_id"`
  Username    string `json:"username"`
  CreateDate  time.Time `json:"create_date"`
}
