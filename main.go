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

  // Load env 
  err := tool.LoadEnvFile(".env")
  if err != nil {
    log.Fatal("Error Loading .env: ", err)
    os.Exit(1)
  }

  
  es, errElastic := db.ElasticNewClient(elasticAddr)
  if errElastic != nil {
    log.Fatal("Error connect elastic db: ", errElastic)
    os.Exit(1)
  }
  
  

  // OS signal trap 
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

  // Graceful shutdown goroutine
  go func() {
    <-sigs
    log.Println("Shutdown signal received")
    cancel()
  }()

  // Connect to database
  dbClient, err := db.NewDatabase(ctx, os.Getenv("DB_URL"), "meeting_bot")
  if err != nil {
    log.Fatalf("DB Failed: %v", err)
  }
  defer dbClient.Close(ctx)

  // Run the app
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
