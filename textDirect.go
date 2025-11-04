package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var sizeFlag = flag.Int("s", 1024, "Size of message receipt buffer. Defaults to 1024.")
	progMode := len(os.Args)

	fmt.Println(progMode)
	fmt.Println(*sizeFlag)
	/*
		go listeningForMsg()
		time.Sleep(2 * time.Second)
		sendMsg()
	*/
}

func sendMsg(addr string, msg string) {
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

func listeningForMsg(bufferSize int) {
	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on ", ln.Addr().String())
	fmt.Println("Press Ctrl+C to Quit.")
	if err != nil {
		fmt.Println("Dial error: ", err)
		return //empty return seems standard but unclear on result in goroutine
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		if conn != nil {
			tmp := make([]byte, bufferSize)
			size, err := conn.Read(tmp)
			if err != nil && size > 0 {
				fmt.Println(err)
				return //if empty return is valid, then ok here?
			}
			fmt.Println("Received: ", string(tmp[:size]))

			//testing
			conn.Close()
			ln.Close()
			return
		}
	}
}
