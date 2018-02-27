package main

  import "fmt"
  import "strings"


  func ExtendedASCIIText(){

      //Slice med Hexverdier
      s := strings.Split("\x22,\x20,\xE2\x82\xac,\x20,\xc3\xb7,\x20,\xc2\xbe,\x20,\x64,\x6f,\x6c,\x6c,\x61,\x72,\x20,\x22", ",")

      //Printer ut hver verdi i slicen s
      for i := 0; i <len(s); i++{

          fmt.Printf(s[i])
        }
  }

  func main(){

    ExtendedASCIIText();
  }
