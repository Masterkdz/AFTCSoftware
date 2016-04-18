package main

import (
  "log"
  "net/http"
  "./router"
)

func main() {
  router := router.NewRouter()
  log.Println("Listening...")
  log.Fatal(http.ListenAndServe(":8080", router))
}

