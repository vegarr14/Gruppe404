package main

import "net"
import "fmt"
import "bufio"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:17")
	// read in input from stdin
	//reader := bufio.NewReader(os.Stdin)
  message, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Print("Message from server: "+message)
	//for {
		//fmt.Print("Text to send: ")
		//text, _ := reader.ReadString('\n')
		// send to socket
		//fmt.Fprintf(conn, text + "\n")
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message from server: "+message)
//	}
}
