package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Creates tcp Server to listen for http requests
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()

	//  Accept  Request
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleRequests(conn)
	}

}

func HandleRequests(con net.Conn) {
	defer con.Close()
	// Read from the request
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("Read from connection::", ln)
		// writing to the connection
		fmt.Fprintf(con, "writing to connection::%s\n", ln)
	}

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code never got here")
}
