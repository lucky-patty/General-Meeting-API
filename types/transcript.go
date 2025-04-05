package types

import "time"

type Transcript struct {
  MeetingID  int            `json:"meeting_id"` 
  UserID     int            `json:"user_id"`
  Transcript string         `json:"transcript"`
  CreateDate time.Time `json:"create_date"`
}


//type TranscriptService struct {
//  Whisper     *WhisperClient 
//  Summarizer  *GPTClient
//  PDfWriter   *PDFGenerator
//  Storage     *MongoRepo
//}


type WhisperTranscript struct {
  Text string `json:"text"`
}
