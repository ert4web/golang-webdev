package main

import (
	"io"
	"log"
	"net"
)

func main() {
	//Creates Server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//Accept Connection
	conn, err := listener.Accept()
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	//write to connection
	io.WriteString(conn, "connection accepted")
}
