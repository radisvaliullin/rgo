package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	// cancle signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		tk := time.NewTimer(time.Second * 15)
		defer tk.Stop()

		select {
		case <-tk.C:
			log.Printf("request: ok")
			fmt.Fprintf(w, "requested path - %q", html.EscapeString(r.URL.Path))
		case <-r.Context().Done():
			log.Printf("request: shutdown")
		}
	})

	// server
	bctx, bcancel := context.WithCancel(context.Background())
	srv := &http.Server{
		BaseContext: func(_ net.Listener) context.Context {
			return bctx
		},
		Addr:         ":7373",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// server listen
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("server: listen fail, %v", err)
		}
	}()

	// shutdonw
	<-sigs
	// notify handlers about shutdonw via cancel context
	bcancel()
	// _ = bcancel
	log.Println("shutdown begin")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("server: shutdonw err, %v", err)
	}
	log.Printf("wait")
	wg.Wait()
	log.Println("done")
}
