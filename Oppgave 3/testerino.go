package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall")

func main() {
	i := 0
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGINT)

	exit_chan := make(chan int)
	go func() {
		for {
			fmt.Println(i)
			i++
			s := <-signal_chan
			switch s {

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Warikomi")
				exit_chan <- 0

			default:
				fmt.Println("Unknown signal.")
				exit_chan <- 1


			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)
}
