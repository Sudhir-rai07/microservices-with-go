package data

import (
	"time"
)

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

var productList = []*Product{
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

func GetProducts() []*Product {
	return productList
}
