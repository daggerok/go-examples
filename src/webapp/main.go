package main

import "net/http"

//import (
//  _ "log"
//  _ "os"
//  _ "text/template"
//  "log"
//)
//
//type Usr struct {
//  id uint64
//  name string
//}
//
//func main() {
//  usr1 := Usr {1, "Max"}
//  usr2 := Usr {2, "Bax"}
//  log.Fatal(usr1, usr2)
//}

func main() {
  http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("hey!"))
  })
  http.ListenAndServe(":8080", nil)
}
