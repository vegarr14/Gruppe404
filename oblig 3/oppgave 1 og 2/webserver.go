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
  if err != nil {
    fmt.Println("error:", err)
  }
  data := htmlStruct{
    PageTitle: "Miljøstasjoner i Stavanger",
    Linje2: "Sted / Plast / glass&metall / tekstil&sko",
    Array: m,
  }
  tmpl.Execute(w, data)
  //io.WriteString(w, "Miljøstasjoner i Stavanger\n")
  //io.WriteString(w, "Sted / Tar plast / tar glass&metall / tar tekstil&sko\n")
  //for i := 0 ; i < len(m) ; i++ {
  //  io.WriteString(w, m[i].Navn + " / " + m[i].Plast + " / " + m[i].Glass_metall + " / " + m[i].Tekstil_sko + "\n")
  //}
 }

 func json2(w http.ResponseWriter, req *http.Request) {
   type Variabler struct {
     Klokkeslett, Sted, Antall_ledige_plasser string
   }
   var m []Variabler
   link := ("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
   err := json.Unmarshal(getJson(link), &m)
   if err != nil {
     fmt.Println("error:", err)
   }
   io.WriteString(w, "viser navn, lokalisering og antall ledige plasser i Stavanger Parkerings 9 parkeringshus\n")
   for i := 0 ; i < len(m) ; i++ {
     io.WriteString(w, m[i].Klokkeslett + " / " + m[i].Sted + " / " + m[i].Antall_ledige_plasser + "\n")
   }
  }
  func json3(w http.ResponseWriter, req *http.Request) {
    type Variabler struct {
      Art, Tjeneste_tekst, Virksomhetsomraade, Art_tekst, Tjeneste_nr, Inntekt_utgift string
   }
    var m []Variabler
    link := ("https://hotell.difi.no/api/json/stavanger/budsjettforslag2016")
    err := json.Unmarshal(getJson(link), &m)
    if err != nil {
      fmt.Println("error:", err)
   }
    io.WriteString(w, "Viser budsjettforslag for 2016 i Stavanger\n")
    for i := 0 ; i < len(m) ; i++ {
      io.WriteString(w, m[i].Art + " / " + m[i].Tjeneste_tekst + " / " + m[i].Virksomhetsomraade + " / " + m[i].Virksomhetsomraade + " / " + m[i].Art_tekst + " / " + m[i].Tjeneste_nr + " / " + m[i].Inntekt_utgift + "\n")
   }
  }
<<<<<<< HEAD

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
    if err != nil {
      fmt.Println("error:", err)
    }
    data := htmlStruct{
      PageTitle: "Utsiktspunkter i Stavanger",
      Linje2: "Navn / Addresse",
      Array: m,
    }
    tmpl.Execute(w, data)

   }

=======
  }
   }
  func json4(w http.ResponseWriter, req *http.Request) {
    type Variabler struct {
      Nord, Barnehagens_navn, Lengdegrad, Oest, Adresse, Breddegrad string
   }
    var m []Variabler
    link := ("https://hotell.difi.no/api/json/stavanger/barnehager")
    err := json.Unmarshal(getJson(link), &m)
    if err != nil {
      fmt.Println("error:", err)
   }
    io.WriteString(w, "Viser budsjettforslag for 2016 i Stavanger\n")
    for i := 0 ; i < len(m) ; i++ {
      io.WriteString(w, m[i].Nord + " / " + m[i].Barnehagens_navn + " / " + m[i].Lengdegrad + " / " + m[i].Oest + " / " + m[i].Adresse + " / " + m[i].Breddegrad + "\n")
   }
  }
>>>>>>> 09fdec80f8adf6837c64c1ffa65159403612243a

func main(){
  http.HandleFunc("/", hello)
  http.HandleFunc("/1", json1)
  http.HandleFunc("/2", json2)
  http.HandleFunc("/3", json3)
  //  http.HandleFunc("/4", json4)
  http.HandleFunc("/5", json5)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
 }

 func getJson(s string) []byte {
   resp, err := http.Get(s)
   if err != nil {
     fmt.Println("error:", err)
   }
   defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   body2 := bytes.Trim(body, `{entries},"pages:1ot234567890}`)
   return body2
 }
