package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defined structure of product API
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `josn:"_"`
	UpdatedAt   string  `josn:"_"`
	DeletedAt   string  `josn:"_"`
}

type Products []*Product

var ProductList = Products{
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

// Validate incoming data
func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// Returns all products
func GetProducts() Products {
	return ProductList
}

// Marshels struct data to JSON
func (p *Products) ToJSON(rw io.Writer) error {
	e := json.NewEncoder(rw)
	return e.Encode(p)
}

// Returns id for new Product ::: Can be done by math/rand stl lib
func genNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.ID + 1
}

// Marshels JSON to struct
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Adds a new product to Product slice
func AddProduct(p Product) {
	p.ID = genNextID()
	ProductList = append(ProductList, &p)
}

// Get product by id
func GetProduct(id int) (*Product, error) {
	for _, item := range ProductList {
		if item.ID == id {
			return item, nil
		}
	}

	return nil, errors.New("Can not find product")
}

// Delete Product by id
func DeleteProduct(id int) error {
	for idx, prod := range ProductList {
		if prod.ID == id {
			ProductList = append(ProductList[:idx], ProductList[idx+1:]...)
			return nil
		}
	}

	return errors.New("Cannot find product")
}
