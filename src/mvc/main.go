package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./src/mvc/templates/*.gohtml"))
}

func about(writer http.ResponseWriter, request *http.Request) {
	err := tpl.ExecuteTemplate(writer, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func setContentType(writer http.ResponseWriter, contentType string) {
	writer.Header().Add("Content-Type", contentType)
}

func assets(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	file, err := os.Open("static/" + path)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	switch {
	case strings.HasSuffix(path, ".html"):
		setContentType(writer, "text/html")
	case strings.HasSuffix(path, ".css"):
		setContentType(writer, "text/css")
	case strings.HasSuffix(path, ".js"):
		setContentType(writer, "application/javascript")
	case strings.HasSuffix(path, ".png"):
		setContentType(writer, "image/png")
	case strings.HasSuffix(path, ".ico"):
		setContentType(writer, "image/x-icon")
	default:
		setContentType(writer, "text/plain")
	}
	defer file.Close()
	io.Copy(writer, file)
}

type user struct {
	Name string
}

func apply(writer http.ResponseWriter, request *http.Request) {
	model := user{
		Name: "guest",
	}
	if request.Method == http.MethodPost {
		model.Name = request.FormValue("name")
	}
	err := tpl.ExecuteTemplate(writer, "apply.gohtml", model)
	if err != nil {
		log.Println(err)
		return
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	err := tpl.ExecuteTemplate(writer, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
		return
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/assets/", assets)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
