package main

import "fmt"

func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 2a
  fmt.Printf("%#X ", sl)
  fmt.Printf("%c", sl)
  fmt.Printf(" %b\n", sl)

}

// Kode for Oppgave 2b
func ExtendedASCIIText() {}




func main() {
  for tekst := 0x80 ; tekst <= 0xFF ; tekst++ {
    fmt.Println(tekst)
  }
}
