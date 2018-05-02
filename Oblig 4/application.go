package main

import (
    "html/template"
    "log"
    "net/http"
    "strings"
    "./weather"
    "fmt"
)

func start(w http.ResponseWriter, r *http.Request) {
    type htmlStruct struct {
      PageTitle string

    }
    tmpl := template.Must(template.ParseFiles("weather.html"))


    data := htmlStruct{
      PageTitle: "Søk her",
    }

    tmpl.Execute(w, data)
}

func getSite (w http.ResponseWriter, r *http.Request) {

  tmpl := template.Must(template.ParseFiles("result.html"))

  r.ParseForm()
  _, data := weather.Weather(strings.Join(r.Form["search"], ""))
  fmt.Println(data)
  tmpl.Execute(w, data)
}

func css(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w, req, "style.css")
}

func main() {
    http.HandleFunc("/", start)
    http.HandleFunc("/result", getSite) // setting router rule
    http.HandleFunc("/style.css", css)
    err := http.ListenAndServe("localhost:8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
