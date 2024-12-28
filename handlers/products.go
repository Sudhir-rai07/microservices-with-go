package handlers

import (
	"log"
	"net/http"

	"github.com/Sudhir-rai07/microservices-with-go/data"
)

type Products struct {
	l *log.Logger
}

func HandleProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Print("Get Products")
	prod := data.GetProducts()

	err := prod.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Failed to marsher json", http.StatusBadRequest)
	}
}
