package db

import (
	"log"

	"github.com/Sudhir-rai07/microservices-with-go/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dns := "host=localhost user=postgres dbname=mygo sslmode=disable password=password"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database %v", err)
	}

	// migrate the user schema
	db.AutoMigrate(&model.User{})

	return db
}
