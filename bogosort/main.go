package main

import (
  "fmt"
  "net/http"
  "os"
  "math/rand"
  "strings"
)

func isSorted(arr []string) bool {
  for i := 0; i < len(arr)-1; i++ {
    if arr[i] > arr[i+1] {
      return false
    }
  }
  return true
}

func shuffle(arr []string) []string {
  for i := len(arr) - 1; i > 0; i-- {
    if j := rand.Intn(i + 1); i != j {
      arr[i], arr[j] = arr[j], arr[i]
    }
  }
  return arr
}

func bogoSort(arr []string) []string {
  for isSorted(arr) == false {
    arr = shuffle(arr)
  }
  return arr
}

func main() {
  port, ok := os.LookupEnv("PORT")
  if !ok {
    port = "8080"
  }
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    thingsStr := query.Get("thing")
    things := []string{}
    if thingsStr != "" {
      things = strings.Split(thingsStr, ",")
    }
    sortedThings := bogoSort(things)
    fmt.Fprintf(w, "I sort things.\nGo to /?thing=$your,list,of,things and I'll sort them.\nThis time I sorted:\n\n")
    fmt.Fprintf(w, strings.Join(sortedThings, "\n"))
  })
  fmt.Println("Listening on " + port)
  http.ListenAndServe(":" + port, nil)
}


