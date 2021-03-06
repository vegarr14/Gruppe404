package main

import (
  "log"
  "io"
  "net/http"
  "html/template"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "bytes"
)

func hello(w http.ResponseWriter, req *http.Request) {
 io.WriteString(w, "Hello, client!\n")
}

func json1(w http.ResponseWriter, req *http.Request) {
  type Variabler struct {
    Navn, Plast, Glass_metall, Tekstil_sko string
  }

  type htmlStruct struct {
    PageTitle string
    Linje2 string
    Array []Variabler
  }

  var m []Variabler
  tmpl := template.Must(template.ParseFiles("layout.html"))
  link := ("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
  err := json.Unmarshal(getJson(link), &m)
  CheckError(err)

  data := htmlStruct{
    PageTitle: "Miljøstasjoner i Stavanger",
    Linje2: "Sted / Plast / glass&metall / tekstil&sko",
    Array: m,
  }

  tmpl.Execute(w, data)
}

func json2(w http.ResponseWriter, req *http.Request) {
  type Variabler struct {
    Klokkeslett, Sted, Antall_ledige_plasser string
  }

  type htmlStruct struct {
    PageTitle string
    Linje2 string
    Array []Variabler
  }

  var m []Variabler
  tmpl := template.Must(template.ParseFiles("layout2.html"))
  link := ("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
  err := json.Unmarshal(getJson(link), &m)
  CheckError(err)

  data := htmlStruct{
    PageTitle: "Parkering i Stavanger",
    Linje2: "Klokkeslett / Lokalisering / Antall ledige plasser",
    Array: m,
  }

  tmpl.Execute(w, data)
}
func json3(w http.ResponseWriter, req *http.Request) {
  type Variabler struct {
    Virksomhetsomraade, Art_tekst, Inntekt_utgift string
  }

  type htmlStruct struct {
    PageTitle string
    Linje2 string
    Array []Variabler
  }

  var m []Variabler
  tmpl := template.Must(template.ParseFiles("layout3.html"))
  link := ("https://hotell.difi.no/api/json/stavanger/budsjettforslag2016")
  err := json.Unmarshal(getJson(link), &m)
  CheckError(err)

  data := htmlStruct{
    PageTitle: "Budsjettforslag for Stavanger, 2016",
    Linje2: "Virksomhetsområde / Beskrivelse / Inntekt eller utgift",
    Array: m,
  }

  tmpl.Execute(w, data)
}

func json4(w http.ResponseWriter, req *http.Request) {
  type Variabler struct {
    Barnehagens_navn, Adresse string
}

type htmlStruct struct {
  PageTitle string
  Linje2 string
  Array []Variabler
}
  var m []Variabler
  tmpl := template.Must(template.ParseFiles("layout4.html"))
  link := ("https://hotell.difi.no/api/json/stavanger/barnehager")
  err := json.Unmarshal(getJson(link), &m)
  CheckError(err)

  data := htmlStruct{
    PageTitle: "Barnehager i Stavanger",
    Linje2: "Navn / Addresse",
    Array: m,
  }

  tmpl.Execute(w, data)
}

func json5(w http.ResponseWriter, req *http.Request) {
  type Variabler struct {
    Name, Adressenavn string
  }

    type htmlStruct struct {
    PageTitle string
    Linje2 string
    Array []Variabler
  }

  var m []Variabler
  tmpl := template.Must(template.ParseFiles("layout5.html"))
  link := ("https://hotell.difi.no/api/json/stavanger/utsiktspunkt")
  err := json.Unmarshal(getJson(link), &m)
  CheckError(err)

  data := htmlStruct{
    PageTitle: "Utsiktspunkter i Stavanger",
    Linje2: "Navn / Addresse",
    Array: m,
  }

  tmpl.Execute(w, data)
}

  func css(w http.ResponseWriter, req *http.Request){
    http.ServeFile(w, req, "style.css")
  }

func main(){
  http.HandleFunc("/", hello)
  http.HandleFunc("/1", json1)
  http.HandleFunc("/2", json2)
  http.HandleFunc("/3", json3)
  http.HandleFunc("/4", json4)
  http.HandleFunc("/5", json5)
  //Henter stylesheet
  http.HandleFunc("/style.css", css)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

//Får link til json, henter innholdet og fjerner unødvendige tegn fra begge ender.
func getJson(s string) []byte {
  resp, err := http.Get(s)
  CheckError(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  body2 := bytes.Trim(body, `{entries},"pages:1ot234567890}`)
  return body2
}

func CheckError(err error) {
  if err != nil {
    fmt.Println("Error: ", err)
  }
}
