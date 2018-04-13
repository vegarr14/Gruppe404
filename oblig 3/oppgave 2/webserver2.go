package main

import (
  "encoding/json"
  "net/http"
  "log"
  "io/ioutil"
  "io"
  "fmt"
  "strings"
)

func webserver(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, client.\n")
}

func json1(w http.ResponseWriter, req *http.Request)  {
    res, err := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
    if err != nil {
      log.Fatal(err)
    }
    miljoGet, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
      log.Fatal(err)
    }
    dataTL := strings.Replace(string(miljoGet),`{"entries":[`,"", 1)
    dataTR := strings.Replace(dataTL,`],"page":1,"pages":1,"posts":56}`,"", 1)
    data := strings.Replace(dataTR,"},{", "}{", -1)
    fmt.Printf("%s", data)

    type Miljostasjon struct {
      Navn, Plast string
    }
    dec := json.NewDecoder(strings.NewReader(data))
    for {
        var m Miljostasjon
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil{
            log.Fatal(err)
        }
        fmt.Printf("%s: %s\n", m.Navn, m.Plast)
    }

}

func main(){

  http.HandleFunc("/", webserver)
  http.HandleFunc("/1", json1)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}
