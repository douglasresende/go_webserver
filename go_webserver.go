package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", hello)
  http.HandleFunc("/test", test)
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }
  fmt.Printf("Listening on http://127.0.0.1:%s\n", port)
  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "hello")
}

func test(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "test")
}
