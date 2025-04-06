package router 

import (
  "net/http"

  "meeting_recorders/controller"
)

func RegisterTranscriptRoutes(mux *http.ServeMux) {
  mux.HandleFunc("/user/test", controller.HelloTest)
}
