package main

import (
  "log"
  "io"
  "net/http"
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
  var m []Variabler
  link := ("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
  err := json.Unmarshal(getJson(link), &m)
  if err != nil {
    fmt.Println("error:", err)
  }
  io.WriteString(w, "Miljøstasjoner i Stavanger\n")
  io.WriteString(w, "Sted / Tar plast / tar glass&metall / tar tekstil&sko\n")
  for i := 0 ; i < len(m) ; i++ {
    io.WriteString(w, m[i].Navn + " / " + m[i].Plast + " / " + m[i].Glass_metall + " / " + m[i].Tekstil_sko + "\n")
  }
 }

func main(){
  http.HandleFunc("/", hello)
  http.HandleFunc("/1", json1)
  http.HandleFunc("/2", json2)
  http.HandleFunc("/3", json3)
  http.HandleFunc("/4", json4)
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
   body2 := bytes.TrimRight(bytes.TrimLeft(body, `{"entries":`), `,"pages:1ot234567890}`)
   return body2
 }
