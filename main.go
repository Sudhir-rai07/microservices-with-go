package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sudhir-rai07/microservices-with-go/db"
	"github.com/Sudhir-rai07/microservices-with-go/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Init db
	db := db.InitDB()

	h := handlers.NewDB(db)
	r := mux.NewRouter()

	r.HandleFunc("/createuser", h.CreateUser).Methods(http.MethodPost)

	// Server setup
	s := &http.Server{
		Addr:         "localhost:5000",
		Handler:      r,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Second,
	}

	// go routine
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal("Error : ", err)
		}
	}()
	fmt.Printf("Server is runninn on port : %v", s.Addr)

	// Make signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill) // Ctrl + C

	sig := <-sigChan
	fmt.Println("Received terminate, Graceful shutdown ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.Shutdown(ctx) // ShutDown server gracefully
}
