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
  leakedMemory := [][1048576]byte{}
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    var memory [1048576]byte
    leakedMemory = append(leakedMemory, memory)
    fmt.Fprintf(w, "ok")
  })
  fmt.Println("Listening on " + port)
  http.ListenAndServe(":" + port, nil)
}


