package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "hei" {
			fmt.Println("hallo")
		}
		if scanner.Text() == "exit" {
			fmt.Println("bye")
			os.Exit(1)
		}
	}
}
