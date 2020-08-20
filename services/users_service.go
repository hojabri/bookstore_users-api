package services

import (
	"github.com/hojabri/bookstore_users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
