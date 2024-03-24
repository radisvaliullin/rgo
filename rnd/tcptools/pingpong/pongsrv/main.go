package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

// simple tcp server
// accept connect and read from conn
// respond to request with copy message
func main() {
	fmt.Println("start pong server.")

	// config
	addr := flag.String("addr", ":4000", "server address")
	flag.Parse()

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen fail: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept fail: %v", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	log.Printf("pong: conn accepted: local - %v, remote - %v", conn.LocalAddr(), conn.RemoteAddr())
	defer func() {
		log.Printf("pong: conn disconnected: local - %v, remote - %v", conn.LocalAddr(), conn.RemoteAddr())
	}()

	bufSize := 2048
	readBuf := make([]byte, bufSize)

	for {

		// read
		n, err := conn.Read(readBuf)
		if err != nil {
			log.Printf("pong: read err: %v", err)
			return
		}
		fmt.Println(string(readBuf[:n]))

		// write copy message
		_, err = conn.Write(readBuf[:n])
		if err != nil {
			log.Printf("pong: write err: %v", err)
			return
		}
	}
}
