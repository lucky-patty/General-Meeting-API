package web

import (
  "fmt"
  "net/http"
  "encoding/json"
  "time"

  "meeting_recorders/types"
  "meeting_recorders/service"
  "meeting_recorders/middleware"
)

type UserController struct {
  Service *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
  return &UserController{
    Service: s,
  }
}

func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
  
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
  var req types.UserRequest
  // Decode
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    http.Error(w, fmt.Sprintf("Invalid Request: %w", err), http.StatusBadRequest)
    return
  }

  ctx := r.Context()
  user, errUser := c.Service.FindByEmail(ctx, req.Email)
  if errUser != nil {
    http.Error(w, fmt.Sprintf("Psql error: %w", err), http.StatusUnauthorized)
    return
  }

  checkHash := middleware.CheckPasswordHash(req.Password, user.Password)
  if checkHash == false {
    http.Error(w, "Invalid email or password", http.StatusUnauthorized)
    return
  }

  // Create JWT 
  token := middleware.GenerateJWT(user.ID) 

  http.SetCookie(w, &http.Cookie{
    Name: "session_token",
    Value: token,
    HttpOnly: true,
    Secure: false,
    Path: "/",
    Expires: time.Now().Add(24 * time.Hour),
  })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Logged in",
    })
  
}
