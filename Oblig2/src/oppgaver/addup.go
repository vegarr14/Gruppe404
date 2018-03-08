package main

import (
  "fmt"
  "os"
  "strconv"
)
func add(test chan int) {
  i := <-test + <-test
  test <- i
}

func main() {
  arg := os.Args[1:]
  x,_ := strconv.Atoi(arg[0])
  y,_ := strconv.Atoi(arg[1])
  test := make(chan int)
  go add(test)
  test <- x
  test <- y
  out := <-test
  fmt.Println(out)
}
