package service

import (
  "fmt"
  "context"
  //"meeting_recorders/types"
  "github.com/elastic/go-elasticsearch/v8"
  "meeting_recorders/db"
)

type TranscriptService struct {
  Psql *db.Psql
  Est  *elasticsearch.Client
}

func (s *TranscriptService) ProcessMeeting( ctx context.Context, audio string) {
   fmt.Println("Process Meeting") 
}


func (s *TranscriptService) Test() {
   fmt.Println("Process Meeting") 
}
