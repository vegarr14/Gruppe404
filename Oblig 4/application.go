package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
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

    type htmlStruct struct {
    PageTitle string
  }


  tmpl := template.Must(template.ParseFiles("result.html"))

  data := htmlStruct{
    PageTitle: "Utsiktspunkter i Stavanger",
  }

  fmt.Println("method:", r.Method) //get request method
  if r.Method == "GET" {
      tmpl.Execute(w, data)
  } else {
      r.ParseForm()
      // logic part of log in
      fmt.Println("Søk:", r.Form["search"])
      data.PageTitle = strings.Join(r.Form["search"], "")
      tmpl.Execute(w, data)
  }

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
