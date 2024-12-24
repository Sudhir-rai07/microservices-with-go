package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sudhir-rai07/microservices-with-go/data"
)

type Product struct {
	l *log.Logger
}

func Products(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Handle an update

	// Else
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetProducts()
	err := json.NewEncoder(rw).Encode(pl)
	if err != nil {
		http.Error(rw, "Error encoding products list", http.StatusInternalServerError)

	}
}
