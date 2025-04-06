package service

import (
  "fmt"
  "bytes"
  "encoding/json"
  "context"
  //"meeting_recorders/types"
  "github.com/elastic/go-elasticsearch/v8"

  "meeting_recorders/db"
  "meeting_recorders/types"
  "meeting_recorders/thirdparty/whisper"
  "meeting_recorders/thirdparty/gpt"
)

type MeetingService struct {
  Psql *db.Psql
  Es  *elasticsearch.Client
  Whisper *whisper.WhisperClient
  Gpt  *gpt.GPTClient
}

func (s *MeetingService) ProcessMeeting( ctx context.Context, audio string) {
   fmt.Println("Process Meeting") 
}


func (s *MeetingService) Test() {
   fmt.Println("Process Meeting") 
}

// Query all
func (s *MeetingService) FindAll(ctx context.Context) ([]types.MeetingNote, error) {
  var buf bytes.Buffer 

  // Build query 
  query := map[string]interface{} {
    "query": map[string]interface{}{
      "match_all": map[string]interface{}{},
    },
  }

  if err := json.NewEncoder(&buf).Encode(query); err != nil {
    return nil, err
  }

  res, errFind := s.Es.Search(
    s.Es.Search.WithContext(ctx),
    s.Es.Search.WithIndex("meeting_notes"),
    s.Es.Search.WithBody(&buf),
    s.Es.Search.WithTrackTotalHits(true),
    s.Es.Search.WithPretty(),
  )

  if errFind != nil {
    return nil, errFind
  }
  defer res.Body.Close() 

  var esResp struct {
    Hits struct {
      Hits []struct {
        Source types.MeetingNote `json:"_source"`
      } `json:"hits"`
    } `json:"hits"`
  }

  if err := json.NewDecoder(res.Body).Decode(&esResp); err != nil {
    return nil, err 
  }

  notes := make([]types.MeetingNote, 0)
  for _, hit := range esResp.Hits.Hits {
    notes = append(notes, hit.Source)
  }

  fmt.Println("Notes: ", notes)
  return notes, nil
}

func (s *MeetingService) FindByMeetingID(ctx context.Context, id string) ([]types.MeetingNote, error) {
  fmt.Println("Meeting ID: ", id)
  var buf bytes.Buffer 
  // Build query 
  query := map[string]interface{} {
    "query": map[string]interface{}{
      "term": map[string]interface{}{
        "meeting_id": map[string]interface{}{
          "value": id,
        },
      },
      
    },
  }

  if err := json.NewEncoder(&buf).Encode(query); err != nil {
    return nil, err
  }

  res, err := s.Es.Search(
    s.Es.Search.WithContext(ctx),
    s.Es.Search.WithIndex("meeting_notes"),
    s.Es.Search.WithBody(&buf),
    s.Es.Search.WithTrackTotalHits(true),
    s.Es.Search.WithPretty(),
  )

  if err != nil {
    return nil, err 
  }
  defer res.Body.Close() 

  var esResp struct {
    Hits struct {
      Hits []struct {
        Source types.MeetingNote `json:"_source"`
      } `json:"hits"`
    } `json:"hits"`
  }

  if err := json.NewDecoder(res.Body).Decode(&esResp); err != nil {
    return nil, err 
  }

  notes := make([]types.MeetingNote, 0)
  for _, hit := range esResp.Hits.Hits {
    notes = append(notes, hit.Source)
  }

  fmt.Println("Notes: ", notes)
  return notes, nil
}

// Store and translate audio to script
func (s *MeetingService) StoreAndTranslate(path string) (string,error) {
  return "", nil
}
// Summarise everything
func (s *MeetingService) SummariseAudio(transcript string) (string, error) {
  return "", nil
}

// Insert to elastic search 
func (s *MeetingService) InsertMeetingNote(ctx context.Context, note types.MeetingNote) error {
  body, err := json.Marshal(note)
  if err != nil {
    return fmt.Errorf("Marshal error: %w", err)
  }

  res, err := s.Es.Index(
    "meeting_notes",
    bytes.NewReader(body),
    s.Es.Index.WithContext(ctx),
    s.Es.Index.WithRefresh("true"),
  )
  if err != nil {
    return fmt.Errorf("ES insert error: %w", err)
  }
  defer res.Body.Close()

  if res.IsError() {
    fmt.Println("ES insert failed: %s", res.String())
    return fmt.Errorf("ES Error: %s", res.String())
  }

  fmt.Println("Meeting note inserted")
  return nil
}
// List All
// List one
