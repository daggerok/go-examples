package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
    path := request.URL.Path
    http.ServeFile(response, request, "src/04-file-server"+path)
  })
  http.ListenAndServe(":8080", nil)
}
