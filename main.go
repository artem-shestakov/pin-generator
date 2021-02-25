package main

import (
	"log"
	"net/http"
	"os"
	"pin-salt-hash/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "Pin_Hash_App", log.LstdFlags)
	hashHandler := handlers.NewHashMux(l)

	serverMux := http.NewServeMux()
	serverMux.Handle("/api/v1/pin", hashHandler)

	s := http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		l.Fatal(err)
	}
}
