package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Transmitting...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(conn, "Hello World")
	conn.Close()
	fmt.Println("...Complete.")
}

func listeningForMsg() {
	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on ", ln.Addr().String())
	if err != nil {
		fmt.Println("Dial error: ", err)
		return //empty return seems standard but unclear on result in goroutine
	}
}
