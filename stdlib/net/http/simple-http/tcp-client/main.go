package main

import (
	"log"
	"net"
)

var srvAddr = ":7373"

func main() {

	cln, err := net.Dial("tcp", srvAddr)
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}

	// get
	httpRequestGet := `GET /ping HTTP/1.1
Host: localhost:7373
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip

`
	_ = httpRequestGet
	// post
	httpRequestPost := `POST /ping HTTP/1.1
Host: localhost:7373
User-Agent: curl/8.7.1
Accept: */*
Content-Type: application/json
Content-Length: 12

{"abc":1234}`
	_ = httpRequestPost

	// make requet / response

	// _, err = cln.Write([]byte(httpRequestGet))
	_, err = cln.Write([]byte(httpRequestPost))
	if err != nil {
		log.Fatalf("write request error: %v", err)
	}

	respBuf := make([]byte, 1024)
	n, err := cln.Read(respBuf)
	if err != nil {
		log.Fatalf("read response error: %v", err)
	}
	log.Printf("response:\n====\n%v\n====\n", string(respBuf[:n]))

	// save response example
	// get
	httpGetResp := `HTTP/1.1 200 OK
Date: Wed, 14 Aug 2024 04:02:20 GMT
Content-Length: 4
Content-Type: text/plain; charset=utf-8

pong`
	_ = httpGetResp
	// post
	httpPostResp := `HTTP/1.1 200 OK
Date: Wed, 14 Aug 2024 04:10:42 GMT
Content-Length: 4
Content-Type: text/plain; charset=utf-8

pong`
	_ = httpPostResp
}
