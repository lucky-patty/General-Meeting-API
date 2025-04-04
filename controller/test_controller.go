package controller 

import (
  "fmt"
  "net/http"
)

func HelloTest(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "Hello World")
}
