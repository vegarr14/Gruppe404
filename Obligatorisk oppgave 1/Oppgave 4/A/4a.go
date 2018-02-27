package main

import "fmt"

func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 2a

//Skriver ut annenhver byte i forskjellige format
  for i := 0; i <len(sl); i++{

    if(i%2 != 0){

      fmt.Printf("%#x ", sl[i])
      fmt.Printf("%c", sl[i])
      fmt.Printf(" %b\n", sl[i])
    }
  }
}

func main() {

  var nrstring string
// generer en string med alle verdiene
  for tekst := 0x80 ; tekst <= 0xFF ; tekst++ {

    nr := (tekst)
    nrstring = (nrstring + string(nr))
  }

  resultat := nrstring

  //Sender Resultat til funksjon
  IterateOverASCIIStringLiteral(resultat)
}
