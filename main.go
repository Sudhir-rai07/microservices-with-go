package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleHomeRoute)

	err := http.ListenAndServe(":8808", nil)
	if err != nil {
		log.Fatal("Error : ", err)
	}
}

func handleHomeRoute(rw http.ResponseWriter, r *http.Request) {
	res := []byte("Happy Coding")
	rw.Write(res)
}
