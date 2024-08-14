package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	cln := http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:7373/ping", nil)
	if err != nil {
		log.Fatalf("client: new request error: %v", err)
	}

	resp, err := cln.Do(req)
	if err != nil {
		log.Fatalf("do request error: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read response bytes error: %v", err)
	}
	log.Printf("response message: %v", string(respBytes))
}
