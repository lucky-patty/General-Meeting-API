package main 

import (
  "context"
  "fmt"
  "log"
  "os"
  "time"
  "encoding/json"
  
  "meeting_recorders/thirdparty/whisper"
  "meeting_recorders/thirdparty/gpt"
  "meeting_recorders/tool"
  "meeting_recorders/service"
  mType "meeting_recorders/types"
  "meeting_recorders/db"
)

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 60 *time.Second)
  defer cancel()

  // Load env 
  err := tool.LoadEnvFile(".env")
  if err != nil {
    log.Fatal("Error Loading .env: ", err)
    os.Exit(1)
  }

  openAIKey := os.Getenv("OPENAI_API_KEY")
  fmt.Println("Open API Key: ", openAIKey)
  if openAIKey == "" {
    log.Fatal("There is no OPENAI KEY")
  }


  elasticAddr := os.Getenv("ELASTICS_ADDR")
 
  fmt.Println("Elastic Address: ", elasticAddr)

  es, errElastic := db.ElasticNewClient(elasticAddr)
  if errElastic != nil {
    log.Fatal("Error connect elastic db: ", errElastic)
    os.Exit(1)
  }

  // Init clients 
  w := &whisper.WhisperClient{APIKey: openAIKey}
  g := &gpt.GPTClient{APIKey: openAIKey, Model: "gpt-3.5-turbo"}
  transcriptService := &service.TranscriptService{
    Psql: nil,
    Est: es,
  }

  meetingService := &service.MeetingService{
    Psql: nil,
    Es: es,
    Gpt: g,
    Whisper: w,
  }
 
  service := &service.Service{
    Transcript: transcriptService,
    Meeting:  meetingService,
  }

  fmt.Println("Run Transcript Check")
  service.Transcript.Test()


  service.Meeting.FindAll(ctx)
  //  audioPath := "record/eng.mp3"

  //test := &type.WhisperTranscript{}
  testResp := `{
  "text": "I hate verbs in English. I dance, you dance, he dances. Why? Is he dancing more than me? I don't think so. 645 people dance and he dances. How much is this mother****** dancing?"
  }`

  var transcript mType.WhisperTranscript
  errTranscript := json.Unmarshal([]byte(testResp), &transcript)
  if errTranscript != nil {
    log.Fatal("Failed to parse transcript: %v ", errTranscript)
  }
  // It work 
  // Transcript: 
// {"text":"I hate verbs in English. I dance, you dance, he dances. Why? Is he dancing more than me? I don't think so. 645 people dance and he dances. How much is this mother****** dancing?"}
  //transcript, err := w.Transcribe(ctx, audioPath)
  //if err != nil {
  //  log.Fatalf("Whisper failed: %v", err)
  //}
  
  fmt.Println("Transcript: \n", transcript.Text)

  summary, err := g.Summarize(ctx, transcript.Text)
  if err != nil {
    log.Fatalf("GPT Summarization failed: %v", err)
  }
  fmt.Println("Summary: \n", summary)
}
