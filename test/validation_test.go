package test

import (
	"log"
	"testing"

	"github.com/Sudhir-rai07/microservices-with-go/data"
)

func TestValidation(t *testing.T) {
	p := &data.Product{
		Name:        "Latte",
		Description: "Latte",
		Price:       0, // Price should be > 0. This test case will be fail
	}

	err := p.Validate()
	if err != nil {
		log.Fatal(err)
	}
}
