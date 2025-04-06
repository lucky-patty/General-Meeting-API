package types

import "time"

type User struct {
  ID          int   `json:"user_id"`
  Email       string `json:"email"`
  Password    string `json:"password"`
  CreateDate  time.Time `json:"create_date"`
}

type UserRequest struct {
  Email     string   `json:"email"`
  Password  string  `json:"password"`
}

