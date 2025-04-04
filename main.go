package main

import (
  "log"
  "net/http"
  "context"
  "os"
  "os/signal"
  "syscall"
  "time"

  "meeting_recorders/router"
)

func main () {
  // Cancelable context shared across app 
  ctx, cancel := context.WithCancel(context.Background())

  // OS signal trap 
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

  // Graceful shutdown goroutine
  go func() {
    <-sigs
    log.Println("Shutdown signal received")
    cancel()
  }()



  go monitorWeb(ctx)
  //  go monitoBot(ctx)
  
  <-ctx.Done()
  log.Println("All services shutdown gracefully")
}


func monitorWeb(ctx context.Context) {
  for {
    err := startWebServer(ctx)
    if err != nil {
      log.Printf("Web server crashed: %v \n", err)
    }

    select {
    case <-ctx.Done():
      log.Println("Stopping Web Server Monitor")    
      return
    default:
      log.Println("Restarting web server ....")
      time.Sleep(2 * time.Second)
    }
  }
}

func startWebServer(ctx context.Context) error {
  srv := &http.Server{
    Addr: ":8080",
    Handler: router.NewRouter(),
  }

  go func() {
    <-ctx.Done()
    log.Println("Web server shutdown initiated")
    ctxShutdown, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
    defer cancel()
    if err := srv.Shutdown(ctxShutdown); err != nil {
      log.Printf("Web server forced shutdown: %v", err)
    }
  }()

  log.Println("Web Server is running on :8080")
  return srv.ListenAndServe() // blocks
}
