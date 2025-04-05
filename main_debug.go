package main 

import (
  "context"
  "fmt"
  "log"
  "os"
  "time"
  "encoding/json"
  
  //"meeting_recorders/thirdparty/whisper"
  "meeting_recorders/thirdparty/gpt"
  "meeting_recorders/tool"
  mType "meeting_recorders/types"
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

//  audioPath := "record/eng.mp3"
  openAIKey := os.Getenv("OPENAI_API_KEY")

  fmt.Println("Open API Key: ", openAIKey)
  if openAIKey == "" {
    log.Fatal("There is no OPENAI KEY")
  }

  // Init clients 
 // w := &whisper.WhisperClient{APIKey: openAIKey}
  g := &gpt.GPTClient{APIKey: openAIKey, Model: "gpt-3.5-turbo"}

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
