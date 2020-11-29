package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gmgale/go_micro/2_intro_to_microservices/handlers"
	"golang.org/x/net/context"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, gracefully shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
