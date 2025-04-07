package middleware

import ( 
  "time"
  "os"

  "golang.org/x/crypto/bcrypt"
  "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int) string {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
  })
  tokenStr, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
  return tokenStr
}

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err 
}

func CheckPasswordHash(password, hash string) bool {
  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
