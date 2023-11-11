package entity

import (
	"github.com/victornascimento22/modgo/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

// value object = ddd
type User struct {
	ID       entity.ID `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       "id",       // You need to provide a valid ID here, not just the string "id"
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
