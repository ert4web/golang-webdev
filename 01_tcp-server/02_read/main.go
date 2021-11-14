package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Creates Server
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()
	// Accepts incoming connections
	conn, err := li.Accept()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	// Read From Connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code never got here")
}
