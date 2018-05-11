package main

import (
    "html/template"
    "log"
    "net/http"
    "strings"
    "weather"
    "fmt"
)
//Returnerer startsiden til klient
func start(w http.ResponseWriter, r *http.Request) {
    type htmlStruct struct {
      PageTitle string

    }
    tmpl := template.Must(template.ParseFiles("weather.html"))


    data := htmlStruct{
      PageTitle: "Search for a city here",
    }

    tmpl.Execute(w, data)
}
/*
* Sender søkeord fra klient til weather.Weather
* og legger return i template før den sender resultatet til klient.
* om Feilmelding oppstår sender klient til 404 siden.
*/
func getSite(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("result.html"))
  r.ParseForm()
  data := weather.Weather(strings.Join(r.Form["search"], ""))
  if data.Feilmelding.Message != "" {
    get404(w,r)
    fmt.Println(data.Feilmelding)
  }else{
    fmt.Println(data)
    tmpl.Execute(w, data)
  }
}
//Sender .CSS fil til klient
func css(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w, req, "style.css")
}
//Sender Værikon til klient
func getWeatherIcon(w http.ResponseWriter, req *http.Request){
  imgPath := weather.ImgPath
  http.ServeFile(w, req, imgPath)
}
//Sender 404 side til klient
func get404(w http.ResponseWriter, req *http.Request){

  type htmlStruct struct{
    PageTitle string
    ErrorMessage string
  }

  tmpl := template.Must(template.ParseFiles("404.html"))

  data := htmlStruct{
    PageTitle: "404 Page not found",
    ErrorMessage: "We were unable to locate the place you were asking for.\n Check your spelling and try again.",
  }

  tmpl.Execute(w, data)
}
//Starter server på port 8080 og kjører funkjsoner basert på http.HandleFunc 
func main() {
    http.HandleFunc("/", start)
    http.HandleFunc("/result", getSite) // setting router rule
    http.HandleFunc("/style.css", css)
    http.HandleFunc("/weatherIcon.png", getWeatherIcon)
    http.HandleFunc("/404", get404)
    err := http.ListenAndServe("localhost:8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
