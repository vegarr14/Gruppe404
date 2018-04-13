package main

import (
  //"encoding/json"
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
    dataTL := strings.TrimLeft(string(miljoGet),'{"entries":[')
    dataTR := strings.TrimRight(DataTR),`],"page":1,"pages":1,"posts":14}`))
    data := sdataTR)
    fmt.Printf("%s", data)

  /*  type Miljostasjon struct {
      navn string
      plast string
    }
    dec := json.NewDecoder(strings.NewReader(data))
    for {
        var m Miljostasjon
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil{
            log.Fatal(err)
        }
        fmt.Printf("%+v", m.navn)
    }
*/
}

func main(){

  http.HandleFunc("/", webserver)
  http.HandleFunc("/1", json1)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}
