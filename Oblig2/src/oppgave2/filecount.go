package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
func main() {
	fmt.Println("Information about text.txt")
	var a [256]int //Setter opp array med 256 int. Alle lik 0
	content, err := ioutil.ReadFile(os.Args[1]) //Leser fil til content
	if err != nil {
		log.Fatal(err)
	}
	//Kjører over hele content og øker tilsvarende verdi i arrayen a for hver verdi content
	for i := 0; i < len(content); i++ {
		a[content[i]]++
	}
	// 0A og 0D tilsvarer linjeskift i tekstfilen. Legger til 1 for å få antall linjer
	fmt.Println("Number of lines in file:", a[10]+1)
	fmt.Println("Most common runes")
	//Loop som kjører 5 ganger(for å printe de 5 meste brukte runene)
	for i := 0; i < 5; i++ {
		antall := 0
		tegn := 0
		//Finner den høyeste verdien i arrayen a
		for i := 0; i < len(a); i++ {
			if antall < a[i] {
				antall = a[i]
				tegn = i
			}
		}
		//Printer tegnet og antall ganger det forekommer i teksten
		fmt.Printf("%d", i+1)
		fmt.Printf(". Rune: ",)
		fmt.Printf("%q", tegn)
		fmt.Println(", Counts: ",antall)
		a[tegn] = 0 //Setter verdien til 0 i arrayen slik at ikke den samme verdien printes om igjen
	}
}
