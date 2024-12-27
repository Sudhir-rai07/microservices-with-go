package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defined structure of product API
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `josn:"createdAt"`
	UpdatedAt   string  `josn:"updatedAt"`
	DeletedAt   string  `josn:"deletedAt"`
}

type Products []*Product

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy Milky coffee",
		Price:       3.45,
		SKU:         "ab21b",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Expresso",
		Description: "Short and strong coffee without milk",
		Price:       5.55,
		SKU:         "es34",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().String(),
	},
}

// Gets all products
func GetProducts() Products {
	return productList
}

// Adds a new Product to productList
func AddProduct(p *Product) {
	p.ID = genID()
	productList = append(productList, p)
}

// genID generates id for the next product
func genID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// Converts the ProductList to JSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Converts to struct from JSON
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
