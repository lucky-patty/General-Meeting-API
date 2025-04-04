package router 

import (
  "net/http"

  "meeting_recorders/controller"
)

func RegisterUserRoutes(mux *http.ServeMux) {
  mux.HandleFunc("/user/test", controller.HelloTest)
}
