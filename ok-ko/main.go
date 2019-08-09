package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  port, ok := os.LookupEnv("PORT")
  if !ok {
    port = "8080"
  }
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "/ok - returns 200 responses\n/ko - returns 500 responses")
  })
  http.HandleFunc("/ok", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ok")
  })
  http.HandleFunc("/ko", func (w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "ko")
  })
  fmt.Println("Listening on " + port)
  http.ListenAndServe(":" + port, nil)
}


