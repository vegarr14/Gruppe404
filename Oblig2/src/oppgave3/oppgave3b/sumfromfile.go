package oppgave3b

import (
  "io/ioutil"
  "bytes"
  "strconv"
  "log"
  "fmt"
  "os"

)
func errortest(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func Sumfromfile() {
  content, err := ioutil.ReadFile("file.txt") //Henter verdiene fra fil
  tall := bytes.Split([]byte(content), []byte(" ")) //Splitter ved mellomrom
  //Konverterer begge tallene til int
  string1 := string(tall[0])
  tall1, err := strconv.Atoi(string1)
  errortest(err)
  string2 := string(tall[1])
  tall2, err := strconv.Atoi(string2)
  errortest(err)
  //Legger sammen og konverterer tilbake til bytes
  sum := tall1 + tall2
  stringsum := strconv.Itoa(sum)
  sumasbyte := []byte(stringsum)
  if (tall1 > 0 && tall2 > 0 && sum < 0) {
    fmt.Println("Too high numbers")
    os.Exit(2)
  } else if (tall1 < 0 && tall2 < 0 && sum > 0) {
    fmt.Println("Too high numbers")
    os.Exit(2)
  }
  err = ioutil.WriteFile("file.txt", sumasbyte, 0644) //Skriver svaret til fil
  errortest(err)
}
