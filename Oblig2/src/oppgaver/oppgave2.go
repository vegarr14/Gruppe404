package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
)

func main() {





	file, _ := os.Open("text.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
	    lineCount++
	}

	fmt.Println("number of lines:", lineCount)

	_, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}










	fmt.Printf("Ferdig")
}
