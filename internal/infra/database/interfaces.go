package database

import "os/user"

type UserInterface interface {
	Create(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}
