package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sudhir-rai07/microservices-with-go/model"
	"gorm.io/gorm"
)

type DB struct {
	mydb *gorm.DB
}

func NewDB(db *gorm.DB) *DB {
	return &DB{mydb: db}
}

func (db *DB) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to unmarshal data from json", http.StatusBadRequest)
		return
	}

	db.mydb.Create(&user)
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to marshal data to json", http.StatusBadRequest)
		return
	}
}
