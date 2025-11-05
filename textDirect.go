package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defBufSize := 1024
	var sizeFlag = flag.Int("s", defBufSize, fmt.Sprintf("Size of message receipt buffer. Defaults to %d.", defBufSize))
	argLen := len(os.Args)
	flag.Parse()

	if argLen == 1 || *sizeFlag != defBufSize {
		go cleanExit()
		listeningForMsg(*sizeFlag)
	} else if argLen == 3 {
		//todo: check type?
		sendMsg(os.Args[1], os.Args[2])
	} else {
		fmt.Printf("--Usage--\nListen:\t\t>program\nListenConfig:\t>program -s [listeningBufferSize]\nSend:\t\t>program targetAddress messageToSend")
	}
}

func cleanExit() {
	sigs := make(chan os.Signal, 1) //size 1 standard?
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sigReceived := <-sigs
	if sigReceived != nil {
		os.Exit(0)
	}
}

func sendMsg(addr string, msg string) {
	fmt.Println("Transmitting...")
	conn, err := net.Dial("tcp", addr) //"127.0.0.1:8080"
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(conn, msg)
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
			conn.Close()
		}
	}
}
