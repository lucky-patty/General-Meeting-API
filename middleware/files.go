package middleware

import (
  "net/http"
)

func FileServer(prefix, dir string) http.Handler {
  return http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
}

