package oppgave3b

import (
  "io/ioutil"
  "bytes"
  "strconv"
  "log"

)

func Sumfromfile() {
  content, err := ioutil.ReadFile("oppgave3b/file.txt")
  tall := bytes.Split([]byte(content), []byte(" "))
  string1 := string(tall[0])
  tall1, _ := strconv.Atoi(string1)
  string2 := string(tall[1])
  tall2, _ := strconv.Atoi(string2)
  sum := tall1 + tall2
  stringsum := strconv.Itoa(sum)
  sumasbyte := []byte(stringsum)
  err = ioutil.WriteFile("oppgave3b/file.txt", sumasbyte, 0644)
  if err != nil {
    log.Fatal(err)
  }
}
