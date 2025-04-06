package types

import (
  "meeting_recorders/db"
)

type UserService struct {
  Psql *db.Psql
}

type MeetingService struct {
  Psql *db.Psql
}
