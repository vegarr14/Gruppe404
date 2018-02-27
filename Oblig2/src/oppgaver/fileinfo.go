package main

import (
  "os"
  "log"
  "fmt"
)

func main () {

  fi, err := os.Lstat("text.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Information about file", fi.Name())
  fmt.Println("Size:", fi.Size(), "bytes." , fi.Size()/1024, "KB.", fi.Size()/1024/1024, "MB.", fi.Size()/1024/1024/1024, "GB.")
}
