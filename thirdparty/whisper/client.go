package whisper


import (
  "bytes"
  "context"
  "io"
  "mime/multipart"
  "net/http"
  "os"
)

type WhisperClient struct {
  APIKey string
}

func (w *WhisperClient) Transcribe(ctx context.Context, path string) (string, error) {
  file, err := os.Open(path)
  if err != nil {
    return "", err 
  }
  defer file.Close()

  var body bytes.Buffer
  writer := multipart.NewWriter(&body)
  part, _ := writer.CreateFormFile("file", file.Name())
  io.Copy(part, file)

  writer.WriteField("model", "whisper-1")
  writer.Close() 

  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.openai.com/v1/audio/transcriptions", &body)
  req.Header.Set("Authorization", "Bearer " +w.APIKey)
  req.Header.Set("Content-Type", writer.FormDataContentType())

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return "", err 
  }
  defer resp.Body.Close()

  data, _ := io.ReadAll(resp.Body)
  return string(data), nil
}
