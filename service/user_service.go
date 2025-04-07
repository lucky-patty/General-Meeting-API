package service

import (
  "fmt"
  "context"

  "meeting_recorders/types"
  "meeting_recorders/db"


	"github.com/jackc/pgx/v5"
)

type UserService struct {
  Psql *db.Psql
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*types.User, error) {
  query := `select id, password from users where email = $1`
  var user types.User 
  err := s.Psql.Pool.QueryRow(ctx,query,email).Scan(&user.ID, &user.Password)

  if err != nil {
    if err == pgx.ErrNoRows {
      return nil, nil
    }

    fmt.Println("Psql error: ", err)
    return nil, err 
  }

  return &user, nil
}

func (s *UserService) Register(ctx context.Context, u *types.UserRequest) (bool, error) {
  query := `insert into users(email,password) values ($1, $2)`
  _, err := s.Psql.Pool.Exec(ctx, query, u.Email, u.Password)
  if err != nil {
    fmt.Println("Psql error: ", err)
    return false, err 
  }

  return true, nil
} 


