package main

import (
  "fmt"
  "os"
  "strconv"
  "os/signal"
  "log"
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

func add(test chan int) {
  // Legger sammen to verdier fra channel test og sender summen tilbake på samme channel
  i := <-test + <-test
  test <- i
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
  //Gjør om argumentene til int.
  x, err := strconv.Atoi(arg[0])
  errortest(err)
  y, err := strconv.Atoi(arg[1])
  errortest(err)
  //Setter opp channel, kjører funskjonen add og sender x og y
  test := make(chan int)
  go add(test)
  test <- x
  test <- y
  //Henter sum fra channelen.
  out := <-test
  if (x > 0 && y > 0 && out < 0) {
    fmt.Println("Too high numbers")
    os.Exit(2)
  } else if (x < 0 && y < 0 && out > 0) {
    fmt.Println("Too high numbers")
    os.Exit(2)
  }
  fmt.Println(out)
}
