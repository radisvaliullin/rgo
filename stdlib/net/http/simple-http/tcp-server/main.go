package main

import (
	"log"
	"net"
)

var srvAddr = ":7373"

func main() {

	ln, err := net.Listen("tcp", srvAddr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("accept error: %v", err)
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	reqBuf := make([]byte, 1024)

	n, err := conn.Read(reqBuf)
	if err != nil {
		log.Printf("conn read request error: %v", err)
		return
	}
	request := reqBuf[:n]
	log.Printf("request:\n%v\n=====\n", string(request))

	respMsg := `HTTP/1.1 200 OK
Date: Wed, 14 Aug 2024 04:02:20 GMT
Content-Length: 4
Content-Type: text/plain; charset=utf-8

pong`

	_, err = conn.Write([]byte(respMsg))
	if err != nil {
		log.Printf("write response error: %v", err)
		return
	}
}
