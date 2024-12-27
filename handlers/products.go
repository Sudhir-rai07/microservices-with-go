package handlers

import (
	"log"
	"net/http"

	"github.com/Sudhir-rai07/microservices-with-go/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProduct created a products handlers with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry for the handler and satisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Handle the request to add a product
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	// fetch the products from data store
	pl := data.GetProducts()
	// Serialize the list to JSON
	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error encoding products list", http.StatusInternalServerError)

	}
}

// addProduct, adds a new product to product list
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Failed to parse to Product from JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}
