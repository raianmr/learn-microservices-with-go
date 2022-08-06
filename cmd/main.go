package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/raianmr/learn-microservices-with-go/msa/handlers"
)

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	sm := http.NewServeMux()

	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sichan := make(chan os.Signal)
	signal.Notify(sichan, os.Interrupt)
	signal.Notify(sichan, os.Kill)

	sig := <-sichan
	l.Println("gracefully shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
