package main

import (
  "net/http"
  "os"
  "log"
  "io"
  "strings"
)

func setContentType(writer http.ResponseWriter, contentType string) {
  writer.Header().Add("Content-Type", contentType)
}

func main() {
  http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
    path := request.URL.Path
    file, err := os.Open("src/file-server" + path)
    if err != nil {
      response.WriteHeader(http.StatusInternalServerError)
      switch {
      case strings.HasSuffix(path, ".html"):
        setContentType(response, "text/html")
      case strings.HasSuffix(path, ".css"):
        setContentType(response, "text/css")
      case strings.HasSuffix(path, ".js"):
        setContentType(response, "application/javascript")
      case strings.HasSuffix(path, ".png"):
        setContentType(response, "image/png")
      case strings.HasSuffix(path, ".ico"):
        setContentType(response, "image/x-icon")
      default:
        setContentType(response, "text/plain")
      }
      log.Println(err)
    }
    defer file.Close()
    io.Copy(response, file)
  })
  http.ListenAndServe(":8080", nil)
}
