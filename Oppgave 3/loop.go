package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
  i := 0
  c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
  for a := 1 ;a < 2; {


    // Block until a signal is received.
    select {
    case s := <-c:
      fmt.Println("Got signal:", s)
      a = 2
    default:
      fmt.Println(i)
      i++
    }
  }
}
