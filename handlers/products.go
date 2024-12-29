package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Sudhir-rai07/microservices-with-go/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func HandleProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Returns all products form product list
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Print("Get Products")
	rw.Header().Set("content-type", "application/json")
	prod := data.GetProducts()

	// Marshle the data to json and returns and error
	err := json.NewEncoder(rw).Encode(prod)
	if err != nil {
		http.Error(rw, "Faild to marshal json", http.StatusBadRequest)
	}

	// err := prod.ToJSON(rw)
	// if err != nil {
	// 	http.Error(rw, "Failed to marsher json", http.StatusBadRequest)
	// }
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Faild to unmarshal ", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Failed to parse id to integer", http.StatusBadRequest)
		return
	}

	prod, err := data.GetProduct(id)
	if err != nil {
		http.Error(rw, "Failed to parse id to integer", http.StatusBadRequest)
		return
	}

	json.NewEncoder(rw).Encode(prod)
}

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Failed to convert id to integer", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write([]byte("Product Deleted"))
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Delete Prod
	err := data.DeleteProduct(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new Product
	var prod data.Product
	err = json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	prod.ID = id
	data.ProductList = append(data.ProductList, &prod)

	err = json.NewEncoder(rw).Encode(prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
