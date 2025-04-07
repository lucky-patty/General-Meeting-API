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

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
  var req types.UserRequest 

  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    http.Error(w, fmt.Sprintf("Invalid Request: %w", err), http.StatusBadRequest)
    return
  }

  ctx := r.Context()
  findUser, errFind  := c.Service.FindByEmail(ctx, req.Email)
  if errFind != nil {
    http.Error(w, fmt.Sprintf("Psql error: %w", errFind), http.StatusInternalServerError)
    return
  }
  
  if findUser != nil {
    http.Error(w, "Duplicate email", http.StatusUnprocessableEntity)
    return
  }

  pass, errPass := middleware.HashPassword(req.Password)
  if errPass != nil {
    http.Error(w, fmt.Sprintf("Auth hash error: %w", errPass), http.StatusInternalServerError)
    return
  }

  user := &types.UserRequest{
    Email: req.Email,
    Password: pass,
  }
  res, errRes := c.Service.Register(ctx, user)
  if errRes != nil {
    http.Error(w, fmt.Sprintf("Register error: %w", errRes), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(res)
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
