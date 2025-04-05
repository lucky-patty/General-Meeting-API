package gpt

import (
  "net/http"
  "encoding/json"
  "io"
  "fmt"
  "context"
  "bytes"
)

type GPTClient struct {
  APIKey string 
  Model  string 
}

type ChatMessage struct {
  Role string `json:"role"`
  Content string `json:"content"`
}

type ChatRequest struct {
  Model string `json:"model"`
  Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
  Choices [] struct {
    Message struct {
      Content string `json:"content"`
    } `json:"message"`
  } `json:"choices"`
}

func (g *GPTClient) Summarize(ctx context.Context, transcript string) (string, error) {

  fmt.Println("API Key: ", g.APIKey)
 // payload := map[string]interface{}{
  //  "model": g.Model,
  //  "messages": []map[string]string{
  //    {"role": "system", "content": "You are a summartization assistant"},
  //    {"role": "user", "content": "Summarize this meeting: " + transcript},
  //  },
 // }
  payload := ChatRequest{
    Model: g.Model, 
    Messages: []ChatMessage{
      { Role: "system", Content: "You are a summartization assistant"},
      { Role: "user", Content: "Summarize this meeting: " + transcript},
    },
  }

  body, _ := json.Marshal(payload)

  req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewReader(body))
  req.Header.Set("Authorization", "Bearer "+ g.APIKey)
  req.Header.Set("Content-Type", "application/json")

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return "", err 
  }
  defer resp.Body.Close() 

  if resp.StatusCode != 200 {
    body, _ := io.ReadAll(resp.Body)
    return  "", fmt.Errorf("GPT Error (%d): %s", resp.StatusCode, string(body))
  }

  var result ChatResponse
  errDecode  := json.NewDecoder(resp.Body).Decode(&result)
  if errDecode != nil {
    fmt.Println("Failed to decode response: ", errDecode)
  }

  var res string = ""

  if len(result.Choices) > 0 {
    fmt.Println("Summary: ", result.Choices[0].Message.Content)
    res = result.Choices[0].Message.Content 
  }
  //choices := result["choices"].([]interface{})
  //message := choices[0].(map[string]interface{})["mesasage"].(map[string]interface{})["content"].(string)


  return res, nil 
}
