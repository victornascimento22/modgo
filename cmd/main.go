package main

import (
	"encoding/json"
	"net/http"

	"github.com/victornascimento22/modgo/configs"
	"github.com/victornascimento22/modgo/internal/dto"
	"github.com/victornascimento22/modgo/internal/entity"
	"github.com/victornascimento22/modgo/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")

	if err != nil {

		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	http.ListenAndServe(":8080", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {

	return &ProductHandler{
		ProductDB: db,
	}

}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
