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
