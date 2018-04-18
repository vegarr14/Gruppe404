package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "strconv"
  "./oppgave3b"
  "os/signal"
)
func a() {
	i := 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for a := 1 ;a < 2; {
    // Case for når et signal kommer på channel c som stopper programmet. Default kjører ellers
		select {
		case s := <-c:
			fmt.Println("Got signal:", s)
			a = 2
			os.Exit(2)
		default:
			i++
		}
	}
}
func errortest(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  go a()
  // Henter argumenter og sjekker om det er nøyaktig 2 argumenter
  arg := os.Args[1:]
  if len(arg) != 2 {
    fmt.Println("Two arguments required")
    os.Exit(2)
  }
  //Gjør om til bytes
  a := []byte(arg[0])
  b := []byte(arg[1])
  //appender for å få en slice som inneholder begge verdiene separert med mellomrom
  slice := append([]byte(a), " "...)
  slice2 := append([]byte(slice), b...)
  err := ioutil.WriteFile("file.txt", slice2, 0644) //Skriver til fil
  errortest(err)
  oppgave3b.Sumfromfile()
  content, err := ioutil.ReadFile("file.txt") //henter svaret fra fil
  errortest(err)
  //Konverterer fra bytes til int
  string := string(content)
  sum, err := strconv.Atoi(string)
  errortest(err)
  fmt.Println(sum)
}
