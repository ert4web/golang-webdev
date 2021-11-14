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
		HandleRequests(conn)
	}
}

func HandleRequests(con net.Conn) {
	method, path := read(con)
	write(con, method, path)
}

func read(con net.Conn) (string, string) {
	i := 0
	var (
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

func write(con net.Conn, m string, p string) {
	body := "Request Method::" + m + "\nRequest Path::" + p
	io.WriteString(con, "HTTP/1.1 200 OK\r\n")            // status line
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body)) // header
	fmt.Fprint(con, "Content-Type: text/plain\r\n")       // header
	io.WriteString(con, "\r\n")                           // blank line; CRLF; carriage-return line-feed
	io.WriteString(con, body)                             // body, aka, payload
}
