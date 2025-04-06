package tool

import (
  "net/http"
)

func Method(method string, h http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != method {
      http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
      return
    }
    h(w,r)
  }
}

