package main

  import ("fmt"
          "bufio"
          "os")


  func ExtendedASCIIText(s string){

    //legger input i en slice
    sl := []byte(s)

    for i := 0; i <len(sl); i++{

        //sjekker hver byte fra input etter extended ASCII
        //Avslutter funksjonen om ved første extended ASCII
        if (s[i] >= 0x80) {
          fmt.Println("Stringen inneholder extended ASCII.")
          return
        }

      }
      fmt.Println("Ingen extended ASCCI funnet!")
  }

  func main(){

    scanner := bufio.NewScanner(os.Stdin)
    //skriver Velkomstmelding
    fmt.Println("Please write an input or type exit to close:")

    for scanner.Scan() {

      //lukker programmet og skriver en avsluttningsmelding
      if scanner.Text() == "exit" {
  			fmt.Println("Closing")
  			os.Exit(1)
      }

      //Sender input til funcksjon
      ExtendedASCIIText(scanner.Text())
      fmt.Println("Please write an input or type exit to close:")
    }
  }
