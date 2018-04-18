package main

import (
	"net"
	"fmt"
//	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":17")
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}

	// accept connection on port
	conn, _ := ln.Accept()
	// sample process for string received
	newmessage := "MEME"
	// send new string back to client
	conn.Write([]byte(newmessage + "\n"))
	// run loop forever (or until ctrl-c)
	//for {
		// will listen for message to process ending in newline (\n)
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		//fmt.Print("Message Received:", string(message))
		// sample process for string received
		//newmessage := strings.ToUpper(message)
		// send new string back to client
		//conn.Write([]byte(newmessage + "\n"))
	}
