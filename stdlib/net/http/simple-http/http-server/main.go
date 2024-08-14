package main

import (
	"log"
	"net/http"
)

var srvAddr = ":7373"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	srv := http.Server{
		Addr:    srvAddr,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server listen and server error: %v", err)
	}
}
