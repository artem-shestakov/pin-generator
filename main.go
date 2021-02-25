package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pin-salt-hash/env"
	"pin-salt-hash/handlers"
	"time"
)

var (
	listenAddres string = env.GetEnv("LISTEN_ADDRESS", "0.0.0.0")
	listenPort   string = env.GetEnv("LISTEN_PORT", "8080")
)

func main() {
	l := log.New(os.Stdout, "Pin_Hash_App", log.LstdFlags)
	hashHandler := handlers.NewHashMux(l)

	serverMux := http.NewServeMux()
	serverMux.Handle("/api/v1/pin", hashHandler)

	s := http.Server{
		Addr:         fmt.Sprintf("%s:%s", listenAddres, listenPort),
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	go func() {
		l.Printf("Server starting: %s", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate signal. Server is shuting down.\n", sig)

	s.Shutdown(ctx)
}
