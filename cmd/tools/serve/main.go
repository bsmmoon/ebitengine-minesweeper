package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("web"))
  http.Handle("/", fs)
  log.Println("Serving ./web at http://0.0.0.0:8080")
  log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
