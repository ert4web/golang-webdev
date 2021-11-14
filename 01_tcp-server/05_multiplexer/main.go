package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(con net.Conn) {
	defer con.Close()
	method, path := read(con)
	write(con, method, path)
}

func read(con net.Conn) (string, string) {
	var (
		i      int
		method string
		path   string
	)
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			sl := strings.Fields(ln)
			method = sl[0]
			path = sl[1]
		}

		if ln == "" {
			break
		}
		i++
	}
	return method, path
}

func write(con net.Conn, method string, path string) {

	// blank line; CRLF; carriage-return line-feed
	if method == "GET" && path == "/about" {
		body := "Request Method::" + method + "\nRequest Path::" + path + "\nWelocme to about Page"
		io.WriteString(con, "HTTP/1.1 200 OK\r\n")            // status line
		fmt.Fprintf(con, "Content-Length: %d\r\n", len(body)) // header
		fmt.Fprint(con, "Content-Type: text/plain\r\n")       // header
		io.WriteString(con, "\r\n")
		io.WriteString(con, body) // body, aka, payload
	} else if method == "GET" && path == "/contactus" {
		body := "Request Method::" + method + "\nRequest Path::" + path + "\nContact us for more information"
		io.WriteString(con, "HTTP/1.1 200 OK\r\n")            // status line
		fmt.Fprintf(con, "Content-Length: %d\r\n", len(body)) // header
		fmt.Fprint(con, "Content-Type: text/plain\r\n")       // header
		io.WriteString(con, "\r\n")
		io.WriteString(con, body) // body, aka, payload
	} else if method == "GET" && path == "/products" {
		body := "Request Method::" + method + "\nRequest Path::" + path + "\nHere is the list of products"
		io.WriteString(con, "HTTP/1.1 200 OK\r\n")            // status line
		fmt.Fprintf(con, "Content-Length: %d\r\n", len(body)) // header
		fmt.Fprint(con, "Content-Type: text/plain\r\n")       // header
		io.WriteString(con, "\r\n")
		io.WriteString(con, body) // body, aka, payload
	} else {
		body := "Request Method::" + method + "\nRequest Path::" + path + "\nWelocme to to Homepage"
		io.WriteString(con, "HTTP/1.1 200 OK\r\n")            // status line
		fmt.Fprintf(con, "Content-Length: %d\r\n", len(body)) // header
		fmt.Fprint(con, "Content-Type: text/plain\r\n")       // header
		io.WriteString(con, "\r\n")
		io.WriteString(con, body) // body, aka, payload
	}

}
