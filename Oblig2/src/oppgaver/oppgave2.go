package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Information about text.txt")
	file, _ := os.Open("text.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
	    lineCount++
	}
	fmt.Println("Number of lines in file:", lineCount)

	var a [256]int
	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(content); i++ {
		a[content[i]]++
	}
	fmt.Println("Most common runes")
	for i := 0; i < 5; i++ {
		antall := 0
		tegn := 0
		for i := 0; i < len(a); i++ {
			if antall < a[i] {
				antall = a[i]
				tegn = i
			}
		}
		fmt.Println(i+1)
		fmt.Printf("Rune: ",)
		fmt.Printf("%q", tegn)
		fmt.Println(", Counts: ",antall)
		a[tegn] = 0
	}
}
