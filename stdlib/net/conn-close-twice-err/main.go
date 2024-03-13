package main

import (
	"log"
	"net"
)

func main() {

	addr := "localhost:7373"
	srvReady := make(chan struct{})
	srvDone := make(chan struct{})

	go func() {

		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("tcp listend: %v\n", err)
		}
		srvReady <- struct{}{}

		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("accept conn: %v\n", err)
		}

		buff := make([]byte, 10)
		n, err := conn.Read(buff)
		if err != nil {
			log.Fatalf("read first bytes: %v\n", err)
		}
		log.Printf("read n - %v bytes: %+v\n", n, buff[:n])

		if err := conn.Close(); err != nil {
			log.Printf("first close attempt: %v\n", err)
		}
		if err := conn.Close(); err != nil {
			log.Printf("second close attempt: %v\n", err)

			if ope, ok := err.(*net.OpError); ok {
				log.Printf("close error is OpError")
				log.Printf("OpError Err is: %+v", ope.Err)
				if ope.Err.Error() == "use of closed network connection" {
					log.Println("err message is \"use of closed network connection\"")
				}
			}
		}
		srvDone <- struct{}{}
	}()

	<-srvReady

	// client
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("client dial: %v\n", err)
	}

	msg := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	_, err = conn.Write(msg)
	if err != nil {
		log.Fatalf("client send message: %v\n", err)
	}

	<-srvDone
}
