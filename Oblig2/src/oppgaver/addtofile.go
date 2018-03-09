package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "strconv"
  "./oppgave3b"
)

func main() {
  arg := os.Args[1:]
  a := []byte(arg[0])
  b := []byte(arg[1])
  slice := append([]byte(a), " "...)
  slice2 := append([]byte(slice), b...)
  err := ioutil.WriteFile("oppgave3b/file.txt", slice2, 0644)
  if err != nil {
    log.Fatal(err)
  }
  oppgave3b.Sumfromfile()
  content, _ := ioutil.ReadFile("oppgave3b/file.txt")
  string := string(content)
  sum, _ := strconv.Atoi(string)
  fmt.Println(sum)
}
