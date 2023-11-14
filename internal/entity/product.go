package entity

import (
	"errors"
	"time"

	"github.com/victornascimento22/modgo/internal/entity"
)

var (
	ErrIDIsRequired      = errors.New("id is required")
	ErrNameIsRequired    = errors.New("name is required")
	ErrPriceIsRequired   = errors.New("price is required")
	ErrInvalidID         = errors.New("invalid id")
	ErrInvalidPrice      = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt string    `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().Format(time.RFC3339), // Convert time to string
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(string(p.ID)); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
