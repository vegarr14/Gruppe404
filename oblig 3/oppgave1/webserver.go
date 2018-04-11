package main

import (
  "net/http"
  "log"
  "io"
)

// hello world, the web server
func webserver(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, client!\n")
}

func main(){

  http.HandleFunc("/", webserver)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}
