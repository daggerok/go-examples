package main

import (
  "net/http"
  "fmt"
)

// custom type
type MyHandler struct {
  greeting string
}

// implement a handler interface which is will call ServeHTTP
func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
  // create instance of custom request handler
  handler := &MyHandler{greeting: "ololo"}
  http.Handle("/", handler)
  http.ListenAndServe(":8080", handler)
}
