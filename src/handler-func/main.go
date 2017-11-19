package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
    response.Write([]byte("hi there!"))
  })
  http.ListenAndServe(":8080", nil)
}
