package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// simple tcp client
// connect to server and send message
// expect to get copy message
// message read from stdin
func main() {
	fmt.Println("ping client start.")

	// config
	addr := flag.String("addr", ":4000", "server address")
	flag.Parse()

	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("ping: dial conn err: %v", err)
	}

	if err := session(conn); err != nil {
		log.Fatalf("ping: session err: %v", err)
	}
}

func session(conn net.Conn) error {
	bufSize := 2048
	readBuf := make([]byte, bufSize)

	input := bufio.NewReaderSize(os.Stdin, bufSize)

	for {
		in, err := input.ReadString('\n')
		if err != nil {
			log.Printf("session: read input err: %v", err)
			return err
		}

		respMsg := strings.TrimSpace(in)
		if len(respMsg) == 0 {
			respMsg = "nope"
		}
		_, err = conn.Write([]byte(respMsg))
		if err != nil {
			log.Printf("session: write to conn err: %v", err)
			return err
		}

		n, err := conn.Read(readBuf)
		if err != nil {
			log.Printf("session: read from conn err: %v", err)
			return err
		}
		fmt.Println(string(readBuf[:n]))
	}
}
