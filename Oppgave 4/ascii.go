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
  var slice [10] int = 0,1,2,3,4,5,6,7,8,9
  for i := 0; i < 256; i++ {
    fmt.Println(slice[i])
  }
}
