package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Information about text.txt")
	var a [256]int
	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(content); i++ {
		a[content[i]]++
	}
	fmt.Println("Number of lines in file:", a[10]+1)
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
