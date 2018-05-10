package main

import (
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

func udp() {
	fmt.Println("Launching udp server...")
	ServerAddr,err := net.ResolveUDPAddr("udp",":17")
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		_,addr,err := ServerConn.ReadFromUDP(buf)
		CheckError(err)
		qoute := "Shoutouts to Simpleflips"
		ServerConn.WriteToUDP([]byte(qoute + "\n"),addr)

	}
}

func tcp() {
	fmt.Println("Launching tcp server...")

	ln, err := net.Listen("tcp", ":17")
	CheckError(err)
	qoute := "Hei"

	for {
		conn, _ := ln.Accept()
		conn.Write([]byte(qoute + "\n"))
		conn.Close()
	}
}

func main() {
	go tcp()
	go udp()
	for{}
}
