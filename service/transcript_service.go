package service



func (s *TranscriptService) ProcessMeeting( ctx context.Context, audio string) {
  transcript := whisper.Transcribe()
  summary := gpt.Summarize(transcript, apiKey)
}
