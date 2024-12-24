package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sudhir-rai07/microservices-with-go/handlers"
)

func main() {
	fmt.Print("Happy Coding")
	l := log.New(os.Stdout, "Product-api", log.LstdFlags)

	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)

	gp := handlers.Products(l)

	sm := http.NewServeMux()
	sm.Handle("/", gp)
	// sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":8808",
		Handler:      sm,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal("Error : ", err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Printf("Recieved terminate, Graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.Shutdown(tc)
}
